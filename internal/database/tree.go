package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Azpect3120/Web-Database-Viewer/internal/model"
	"github.com/Azpect3120/Web-Database-Viewer/internal/query"
	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/microsoft/go-mssqldb"
)

// Return an HTML string with the contents of the database tables
// in tree format
func TableTree(c *gin.Context) string {
	session := sessions.Default(c)
	connections_bytes, ok := session.Get("connections").([]byte)
	current, ok := session.Get("current").(string)
	if !ok {
		return templates.TableTreeError(errors.New("No connections found"))
	}

	var connections map[string][2]string
	if err := json.Unmarshal(connections_bytes, &connections); err != nil {
		fmt.Println(err)
		return templates.TableTreeError(err)
	}

	var (
		url    string = connections[current][0]
		driver string = connections[current][1]
	)

	tree, err := generateTableTree(url, driver)
	if err != nil {
		fmt.Println(err)
		return templates.TableTreeError(err)
	}

	return templates.TableTree(tree)
}

// Generate the tree of database table names. Column metadata is loaded on demand.
func generateTableTree(url, driver string) (map[string][]model.Column, error) {
	conn, err := sql.Open(sqlDriver(driver), url)
	if err != nil {
		return map[string][]model.Column{}, err
	}
	defer conn.Close()

	tree, err := tableList(conn, driver)
	if err != nil {
		return map[string][]model.Column{}, err
	}

	return tree, nil
}

// TableColumns returns the metadata for a single table in the current connection.
func TableColumns(c *gin.Context) string {
	table := c.Query("table")
	if table == "" {
		return templates.TableFieldsError(errors.New("No table selected"))
	}

	session := sessions.Default(c)
	connectionsBytes, ok := session.Get("connections").([]byte)
	current, ok := session.Get("current").(string)
	if !ok {
		return templates.TableFieldsError(errors.New("No connections found"))
	}

	var connections map[string][2]string
	if err := json.Unmarshal(connectionsBytes, &connections); err != nil {
		return templates.TableFieldsError(err)
	}

	conn, err := sql.Open(sqlDriver(connections[current][1]), connections[current][0])
	if err != nil {
		return templates.TableFieldsError(err)
	}
	defer conn.Close()

	columns, err := tableColumns(conn, connections[current][1], table)
	if err != nil {
		return templates.TableFieldsError(err)
	}

	return templates.TableFields(table, columns)
}

// Return a map with the keys being the table names and the values
// being blank which can be later used to store the columns.
func tableList(conn *sql.DB, driver string) (map[string][]model.Column, error) {
	var q string
	switch driver {
	case "postgres":
		q = query.GET_TABLE_LIST_PSQL
	case "mysql", "mariadb":
		q = query.GET_TABLE_LIST_MYSQL
	case "sqlite3":
		q = query.GET_TABLE_LIST_SQLITE
	case "sqlserver":
		q = query.GET_TABLE_LIST_MSSQL
	default:
		return map[string][]model.Column{}, errors.New("Table List: Unsupported driver")
	}

	rows, err := conn.Query(q)
	if err != nil {
		return map[string][]model.Column{}, err
	}
	defer rows.Close()

	tree := make(map[string][]model.Column)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return map[string][]model.Column{}, err
		}
		tree[table] = []model.Column{}
	}

	return tree, nil
}

// tableQueries returns the metadata queries for one table.
func tableQueries(driver string) ([4]string, error) {
	var qs [4]string
	switch driver {
	case "postgres":
		qs = [4]string{
			query.GET_TABLE_PK_PSQL,
			query.GET_TABLE_FKS_PSQL,
			query.GET_TABLE_RESTRAINS_PSQL,
			query.GET_TABLE_UNIQUE_COLS_PSQL,
		}
	case "mysql", "mariadb":
		qs = [4]string{
			query.GET_TABLE_PK_MYSQL,
			query.GET_TABLE_FKS_MYSQL,
			query.GET_TABLE_RESTRAINS_MYSQL,
			query.GET_TABLE_UNIQUE_COLS_MYSQL,
		}
	case "sqlite3":
		qs = [4]string{
			query.GET_TABLE_PK_SQLITE,
			query.GET_TABLE_FKS_SQLITE,
			query.GET_TABLE_RESTRAINS_SQLITE,
			query.GET_TABLE_UNIQUE_COLS_SQLITE,
		}
	case "sqlserver":
		qs = [4]string{
			query.GET_TABLE_PK_MSSQL,
			query.GET_TABLE_FKS_MSSQL,
			query.GET_TABLE_RESTRAINS_MSSQL,
			query.GET_TABLE_UNIQUE_COLS_MSSQL,
		}
	default:
		return qs, errors.New("Table Columns: Unsupported driver")
	}

	return qs, nil
}

// tableColumns returns all metadata required to display one table's columns.
func tableColumns(conn *sql.DB, driver, table string) ([]model.Column, error) {
	qs, err := tableQueries(driver)
	if err != nil {
		return nil, err
	}

	table = strings.ReplaceAll(table, "'", "''")

	unique := make(map[string]bool)
	rows, err := conn.Query(fmt.Sprintf(qs[3], table))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			rows.Close()
			return nil, err
		}
		unique[column] = true
	}
	rows.Close()

	primaryKeys := make(map[string]bool)
	rows, err = conn.Query(fmt.Sprintf(qs[0], table))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			rows.Close()
			return nil, err
		}
		primaryKeys[column] = true
	}
	rows.Close()

	foreignKeys := make(map[string]model.ForeignKey)
	rows, err = conn.Query(fmt.Sprintf(qs[1], table))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var foreignKey model.ForeignKey
		if err := rows.Scan(new(interface{}), new(interface{}), &foreignKey.Column, new(interface{}), &foreignKey.ForeignTable, &foreignKey.ForeignColumn); err != nil {
			rows.Close()
			return nil, err
		}
		foreignKeys[foreignKey.Column] = foreignKey
	}
	rows.Close()

	rows, err = conn.Query(fmt.Sprintf(qs[2], table))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := make([]model.Column, 0)
	for rows.Next() {
		var (
			column   model.Column
			enumType sql.NullString
		)
		if err := rows.Scan(&column.Name, &column.Nullable, &column.Type, &column.MaxLength, &enumType); err != nil {
			return nil, err
		}
		if column.Type == "USER-DEFINED" {
			column.Type = enumType.String
		}
		column.PrimaryKey = primaryKeys[column.Name]
		column.ForeignKey = foreignKeys[column.Name]
		column.Unique = unique[column.Name]
		columns = append(columns, column)
	}

	return columns, nil
}

// Generate the tree of the database enums and their values
func EnumTree(c *gin.Context) string {
	session := sessions.Default(c)
	connections_bytes, ok := session.Get("connections").([]byte)
	current, ok := session.Get("current").(string)
	if !ok {
		return templates.EnumTreeError(errors.New("No connections found"))
	}

	var connections map[string][2]string
	if err := json.Unmarshal(connections_bytes, &connections); err != nil {
		return templates.EnumTreeError(err)
	}

	var (
		url    string = connections[current][0]
		driver string = connections[current][1]
	)

	enums, err := genereteEnumTree(url, driver)
	if err != nil {
		return templates.EnumTreeError(err)
	}

	return templates.EnumTree(enums)
}

// Generate the tree of the database enums and their values from a
// provided connection URL.
func genereteEnumTree(url, driver string) (map[string][]string, error) {
	conn, err := sql.Open(sqlDriver(driver), url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	enums, err := enumList(conn, driver)
	if err != nil {
		return nil, err
	}

	return enums, nil
}

// Get a list/map of all the enums in the database.
// The key is the name of the enum and the value is a slice of the enum values.
func enumList(conn *sql.DB, driver string) (map[string][]string, error) {
	var q string
	switch driver {
	case "postgres":
		q = query.GET_ENUM_LIST_PSQL
	case "mysql", "mariadb", "sqlite3", "sqlserver":
		return map[string][]string{}, errors.New(fmt.Sprintf("%s does not support enum tree display.", driver))
	default:
		return map[string][]string{}, errors.New("Enum List: Unsupported driver")
	}
	rows, err := conn.Query(q)
	if err != nil {
		return map[string][]string{}, err
	}
	defer rows.Close()

	enums := make(map[string][]string)
	for rows.Next() {
		var enum, value string
		if err := rows.Scan(&enum, &value); err != nil {
			return map[string][]string{}, err
		}

		enums[enum] = append(enums[enum], value)
	}

	return enums, nil
}

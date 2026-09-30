package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

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

// Generate the tree of the database tables
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

	if err := fillColumns(conn, driver, tree); err != nil {
		return map[string][]model.Column{}, err
	}

	return tree, nil
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

// Fill the columns of the tables in the tree using the keys found
// in the tableList function.
func fillColumns(conn *sql.DB, driver string, tree map[string][]model.Column) error {
	// Pick the correct array of queries to use based on the driver
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
		return errors.New("Table Columns: Unsupported driver")
	}

	var pkey string
	var fkeys []model.ForeignKey
	for table := range tree {
		unique, err := getUniqueColumns(conn, table, qs[3])
		if err != nil {
			return err
		}

		// Get the primary key of the table
		pk, err := conn.Query(fmt.Sprintf(qs[0], table))
		if err != nil {
			return err
		}
		defer pk.Close()
		for pk.Next() {
			if err := pk.Scan(&pkey); err != nil {
				return err
			}
		}

		// Get the foreign keys of the table
		fk, err := conn.Query(fmt.Sprintf(qs[1], table))
		if err != nil {
			return err
		}
		defer fk.Close()
		for fk.Next() {
			var fkey model.ForeignKey
			if err := fk.Scan(new(interface{}), new(interface{}), &fkey.Column, new(interface{}), &fkey.ForeignTable, &fkey.ForeignColumn); err != nil {
				return err
			}
			fkeys = append(fkeys, fkey)
		}

		// Get the restraints of the table
		rows, err := conn.Query(fmt.Sprintf(qs[2], table))
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				column   model.Column
				enumType sql.NullString
			)
			if err := rows.Scan(&column.Name, &column.Nullable, &column.Type, &column.MaxLength, &enumType); err != nil {
				return err
			}
			if column.Type == "USER-DEFINED" {
				column.Type = enumType.String
			}
			if column.Name == pkey {
				column.PrimaryKey = true
			}
			for _, fkey := range fkeys {
				if column.Name == fkey.Column {
					column.ForeignKey = fkey
				} else {
					column.ForeignKey = model.ForeignKey{}
				}
			}

			for _, u := range unique {
				if column.Name == u {
					column.Unique = true
				}
			}

			tree[table] = append(tree[table], column)
		}
	}

	return nil
}

// Returns a list of the unique columns in a table
func getUniqueColumns(conn *sql.DB, table string, query string) ([]string, error) {
	var cols []string
	rows, err := conn.Query(fmt.Sprintf(query, table))
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var col string
		if err := rows.Scan(&col); err != nil {
			return []string{}, err
		}
		cols = append(cols, col)
	}

	return cols, nil
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

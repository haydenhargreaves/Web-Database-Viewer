package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	mssql "github.com/microsoft/go-mssqldb"
)

const maxQueryRows = 500

func QueryCurrent(c *gin.Context) {
	query := c.PostForm("sql")
	conn, driver := getConnection(c)

	if query == "" {
		c.String(200, templates.ErrorQueryResults(fmt.Errorf("No query provided")))
		return
	}

	queries := strings.Split(query, ";")
	var results []string
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		cols, data, truncated, err := queryConnection(query, conn, driver)
		if err != nil {
			c.String(200, templates.ErrorQueryResults(err))
			return
		}

		results = append(results, templates.QueryResult(cols, data, truncated))
	}

	c.String(200, templates.ConcatResults(results))
}

func queryConnection(query, url, driver string) ([]string, []map[string]interface{}, bool, error) {
	db, err := sql.Open(sqlDriver(driver), url)
	if err != nil {
		return []string{}, []map[string]interface{}{}, false, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return []string{}, []map[string]interface{}{}, false, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return []string{}, []map[string]interface{}{}, false, err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return []string{}, []map[string]interface{}{}, false, err
	}

	// Create values pointer and value pointer slices
	// We need to use pointers because the Scan function
	// requires a pointer to the value. However, there
	// is no simple way to create an array of pointers.
	values := make([]interface{}, len(cols))
	valuePtrs := make([]interface{}, len(cols))

	for i := range cols {
		if isUniqueIdentifier(driver, columnTypes[i].DatabaseTypeName()) {
			valuePtrs[i] = &mssql.NullUniqueIdentifier{}
		} else {
			valuePtrs[i] = &values[i]
		}
	}

	// Final data structure to store the results
	// An array of maps, where each map is a row
	// and the keys are the column names.
	var result []map[string]interface{}
	truncated := false

	for rows.Next() {
		if len(result) == maxQueryRows {
			truncated = true
			break
		}

		// Scan the result into the value pointers
		err := rows.Scan(valuePtrs...)
		if err != nil {
			fmt.Println(err)
		}

		// Create a map to store the column data
		row := make(map[string]interface{})
		for i, col := range cols {
			var v interface{}
			if guid, ok := valuePtrs[i].(*mssql.NullUniqueIdentifier); ok {
				v = uniqueIdentifierValue(guid)
			} else {
				v = queryValue(values[i])
			}

			row[col] = v
		}

		// Append the row to the result set
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return []string{}, []map[string]interface{}{}, false, err
	}

	return cols, result, truncated, nil
}

func isUniqueIdentifier(driver, databaseType string) bool {
	return driver == "sqlserver" && strings.EqualFold(databaseType, "UNIQUEIDENTIFIER")
}

func uniqueIdentifierValue(guid *mssql.NullUniqueIdentifier) interface{} {
	if !guid.Valid {
		return nil
	}

	return guid.UUID.String()
}

func queryValue(value interface{}) interface{} {
	if bytes, ok := value.([]byte); ok {
		return string(bytes)
	}

	return value
}

func getConnection(c *gin.Context) (url, driver string) {
	session := sessions.Default(c)
	conn_bytes, ok := session.Get("connections").([]byte)
	curr, ok := session.Get("current").(string)
	if !ok {
		fmt.Println("No current connection")
		return "", ""
	}

	var connections map[string][2]string
	if err := json.Unmarshal(conn_bytes, &connections); err != nil {
		fmt.Println(err)
		return "", ""
	}

	return connections[curr][0], connections[curr][1]
}

package database

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const CONNECTION_SUCCESS string = `
	<span id="connection-status" class="w-3 h-3 bg-green-500 rounded-full"></span>
	<span hx-swap-oob="outerHTML" id="connection-message" class="text-green-500"> Connection successful! </span>
	<span hx-swap-oob="outerHTML" id="db-url-invalid" class="text-xs text-red-500 hidden"></span>
`

const CONNECTION_FAILURE string = `
	<span id="connection-status" class="w-3 h-3 bg-red-500 rounded-full"></span>
	<span hx-swap-oob="outerHTML" id="connection-message" class="text-red-500"> Connection failed! </span>
	<span hx-swap-oob="outerHTML" id="db-url-invalid" class="text-xs text-red-500">%s</span>
`

// Test a connection to a database
func TestConnectionURL(c *gin.Context) {
	driver := c.PostForm("db-driver")
	switch driver {
	case "postgres":
	case "mysql", "mariadb":
	case "sqlite3":
	default:
		c.String(200, fmt.Sprintf(CONNECTION_FAILURE, "Unsupported driver"))
		return
	}

	// Open connection
	conn, err := sql.Open(sqlDriver(driver), c.PostForm("db-url"))
	if err != nil {
		fmt.Println(err)
		c.String(200, fmt.Sprintf(CONNECTION_FAILURE, err.Error()))
		return
	}

	// Ping/test connection
	if err := conn.Ping(); err != nil {
		c.String(200, fmt.Sprintf(CONNECTION_FAILURE, err.Error()))
		return
	} else {
		fmt.Printf("%+v\n", conn.Driver())
	}

	c.String(200, CONNECTION_SUCCESS)
}

package database

// sqlDriver returns the database/sql driver name for a selected connection type.
// MariaDB uses the MySQL-compatible go-sql-driver/mysql driver.
func sqlDriver(driver string) string {
	if driver == "mariadb" {
		return "mysql"
	}

	return driver
}

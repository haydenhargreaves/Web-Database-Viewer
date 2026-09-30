package database

import "testing"

func TestSQLDriver(t *testing.T) {
	tests := map[string]string{
		"postgres": "postgres",
		"mysql":    "mysql",
		"mariadb":  "mysql",
		"sqlite3":  "sqlite3",
	}

	for input, want := range tests {
		if got := sqlDriver(input); got != want {
			t.Errorf("sqlDriver(%q) = %q, want %q", input, got, want)
		}
	}
}

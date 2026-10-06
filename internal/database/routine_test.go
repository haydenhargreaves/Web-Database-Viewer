package database

import (
	"testing"

	"github.com/Azpect3120/Web-Database-Viewer/internal/query"
)

func TestRoutineListQuery(t *testing.T) {
	tests := map[string]string{
		"postgres":  query.GET_ROUTINE_LIST_PSQL,
		"mysql":     query.GET_ROUTINE_LIST_MYSQL,
		"mariadb":   query.GET_ROUTINE_LIST_MYSQL,
		"sqlserver": query.GET_ROUTINE_LIST_MSSQL,
		"sqlite3":   "",
	}

	for driver, want := range tests {
		got, err := routineListQuery(driver)
		if err != nil {
			t.Errorf("routineListQuery(%q) returned an error: %v", driver, err)
		}
		if got != want {
			t.Errorf("routineListQuery(%q) = %q, want %q", driver, got, want)
		}
	}

	if _, err := routineListQuery("oracle"); err == nil {
		t.Error("routineListQuery() accepted an unsupported driver")
	}
}

func TestRoutineListSQLite(t *testing.T) {
	routines, err := routineList(nil, "sqlite3")
	if err != nil {
		t.Fatalf("routineList() returned an error for SQLite: %v", err)
	}
	if len(routines) != 0 {
		t.Errorf("routineList() returned %d routines for SQLite, want 0", len(routines))
	}
}

func TestRoutineDefinitionRejectsInvalidIdentifiers(t *testing.T) {
	if _, err := routineDefinition(nil, "postgres", "not-an-oid", "routine", "function"); err == nil {
		t.Error("routineDefinition() accepted an invalid PostgreSQL routine ID")
	}
	if _, err := routineDefinition(nil, "sqlserver", "not-an-id", "routine", "procedure"); err == nil {
		t.Error("routineDefinition() accepted an invalid SQL Server routine ID")
	}
	if _, err := routineDefinition(nil, "mysql", "routine", "routine", "trigger"); err == nil {
		t.Error("routineDefinition() accepted an invalid MySQL routine type")
	}
}

func TestViewListQuery(t *testing.T) {
	tests := []struct {
		driver  string
		want    string
		wantErr bool
	}{
		{driver: "postgres", want: query.GET_VIEW_LIST_PSQL},
		{driver: "mysql", want: query.GET_VIEW_LIST_MYSQL},
		{driver: "mariadb", want: query.GET_VIEW_LIST_MYSQL},
		{driver: "sqlite3", want: query.GET_VIEW_LIST_SQLITE},
		{driver: "sqlserver", want: query.GET_VIEW_LIST_MSSQL},
		{driver: "oracle", wantErr: true},
	}

	for _, test := range tests {
		got, err := viewListQuery(test.driver)
		if test.wantErr {
			if err == nil {
				t.Errorf("viewListQuery(%q) returned nil error", test.driver)
			}
			continue
		}
		if err != nil {
			t.Errorf("viewListQuery(%q) returned error: %v", test.driver, err)
			continue
		}
		if got != test.want {
			t.Errorf("viewListQuery(%q) = %q, want %q", test.driver, got, test.want)
		}
	}
}

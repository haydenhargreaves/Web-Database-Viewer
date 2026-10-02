package database

import (
	"testing"

	mssql "github.com/microsoft/go-mssqldb"
)

func TestIsUniqueIdentifier(t *testing.T) {
	tests := []struct {
		driver       string
		databaseType string
		want         bool
	}{
		{driver: "sqlserver", databaseType: "UNIQUEIDENTIFIER", want: true},
		{driver: "sqlserver", databaseType: "uniqueidentifier", want: true},
		{driver: "sqlserver", databaseType: "VARBINARY", want: false},
		{driver: "postgres", databaseType: "UNIQUEIDENTIFIER", want: false},
	}

	for _, test := range tests {
		if got := isUniqueIdentifier(test.driver, test.databaseType); got != test.want {
			t.Errorf("isUniqueIdentifier(%q, %q) = %t, want %t", test.driver, test.databaseType, got, test.want)
		}
	}
}

func TestUniqueIdentifierValue(t *testing.T) {
	guid := &mssql.NullUniqueIdentifier{
		UUID:  mssql.UniqueIdentifier{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
		Valid: true,
	}
	if got, want := uniqueIdentifierValue(guid), "00112233-4455-6677-8899-AABBCCDDEEFF"; got != want {
		t.Errorf("uniqueIdentifierValue() = %v, want %v", got, want)
	}

	if got := uniqueIdentifierValue(&mssql.NullUniqueIdentifier{}); got != nil {
		t.Errorf("uniqueIdentifierValue() = %v, want nil", got)
	}
}

func TestQueryValue(t *testing.T) {
	if got, want := queryValue([]byte("text")), "text"; got != want {
		t.Errorf("queryValue() = %v, want %v", got, want)
	}
}

package templates

import (
	"strings"
	"testing"
)

func TestQueryResultTruncationWarning(t *testing.T) {
	rows := []map[string]interface{}{{"id": 1}}

	truncated := QueryResult([]string{"id"}, rows, true)
	if !strings.Contains(truncated, "Showing first 500 rows") {
		t.Error("QueryResult() did not include the truncation warning")
	}

	complete := QueryResult([]string{"id"}, rows, false)
	if strings.Contains(complete, "Showing first 500 rows") {
		t.Error("QueryResult() included the truncation warning for a complete result")
	}
}

package templates

import (
	"crypto/sha256"
	"fmt"
	"html"
	"strings"
	"testing"
)

func TestConnectionHistoryIdentity(t *testing.T) {
	names := []string{"", "A &amp; B", "A & B", "<script>", `"quote"`, " spaced ", " ", "日本語", "No connections"}
	connections := make(map[string][2]string)
	for _, name := range names {
		connections[name] = [2]string{"db", "sqlite3"}
	}
	for _, name := range names {
		digest := fmt.Sprintf("%x", sha256.Sum256([]byte(name)))
		if QueryHistoryID(name) != digest {
			t.Fatalf("wrong digest for %q", name)
		}
		got := ConnectionsList(connections, name)
		want := `data-query-history-id="` + digest + `">` + html.EscapeString(name) + `</option>`
		if !strings.Contains(got, want) {
			t.Errorf("missing exact option for %q", name)
		}
		if !strings.Contains(got, fmt.Sprintf(TREE_VIEW_NAME, html.EscapeString(name))) {
			t.Errorf("unsafe tree for %q", name)
		}
	}
	if strings.Contains(ConnectionsList(nil, ""), "data-query-history-id") {
		t.Fatal("placeholder has identity")
	}
	missing := ConnectionsList(map[string][2]string{"other": {"db", "sqlite3"}}, "missing")
	first := strings.Split(missing, "</option>")[0]
	if strings.Contains(first, "data-query-history-id") {
		t.Fatal("missing current has identity")
	}
}

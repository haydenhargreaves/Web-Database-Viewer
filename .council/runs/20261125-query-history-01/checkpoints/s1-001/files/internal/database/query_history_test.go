package database

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func TestQueryHistoryResponseIdentity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	names := []string{"", "A &amp; B", "A & B", "<script>", " spaced ", " ", "日本語", "No connections"}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			r := gin.New()
			r.Use(sessions.Sessions("test", cookie.NewStore([]byte("01234567890123456789012345678901"))))
			r.Use(func(c *gin.Context) {
				data, _ := json.Marshal(map[string][2]string{name: {":memory:", "sqlite3"}})
				session := sessions.Default(c)
				session.Set("connections", data)
				session.Set("current", name)
				c.Next()
			})
			r.POST("/", QueryCurrent)
			for _, sql := range []string{"", "select 1", "select * from missing"} {
				w := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(url.Values{"sql": {sql}}.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				r.ServeHTTP(w, req)
				if w.Code != 200 || w.Header().Get("X-Query-History-Connection") != templates.QueryHistoryID(name) {
					t.Fatalf("sql %q: status %d header %q", sql, w.Code, w.Header().Get("X-Query-History-Connection"))
				}
				if sql == "select 1" && !strings.Contains(w.Body.String(), "1") {
					t.Fatal("query result missing")
				}
			}
		})
	}
	r := gin.New()
	r.Use(sessions.Sessions("test", cookie.NewStore([]byte("01234567890123456789012345678901"))))
	r.Use(func(c *gin.Context) {
		s := sessions.Default(c)
		s.Set("connections", []byte(`{"other":[":memory:","sqlite3"]}`))
		s.Set("current", "missing")
		c.Next()
	})
	r.POST("/", QueryCurrent)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/", nil))
	if w.Header().Get("X-Query-History-Connection") != "" {
		t.Fatal("missing member has identity")
	}
}

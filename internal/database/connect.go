package database

import (
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Connections in the drop down are stored in a
// JSON format, which this object will be mapped to.
type ConnectListItem struct {
	Driver string `json:"driver"`
	URL    string `json:"url"`
}

// Change the current connection in the session
func ChangeConnection(c *gin.Context) {
	conn := c.PostForm("connected-database")

	var item ConnectListItem
	if err := json.Unmarshal([]byte(conn), &item); err != nil {
		fmt.Println(err)
		c.String(200, templates.ConnectionsList(nil, ""))
		return
	}

	session := sessions.Default(c)
	conn_bytes, ok := session.Get("connections").([]byte)
	if !ok {
		c.String(200, templates.ConnectionsList(nil, ""))
		return
	}

	var connections map[string][2]string
	if err := json.Unmarshal(conn_bytes, &connections); err != nil {
		c.String(200, templates.ConnectionsList(nil, ""))
		fmt.Println(err)
		return
	}

	var newName string
	for name, data := range connections {
		if data[0] == item.URL {
			newName = name
			break
		}
	}

	session.Set("current", newName)
	session.Save()

	c.String(200, templates.ConnectionsList(connections, newName)+TableTree(c)+ViewTree(c)+EnumTree(c)+RoutineTree(c))
}

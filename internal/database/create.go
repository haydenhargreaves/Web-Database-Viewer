package database

import (
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Create a new connection to a database and store the details
// in the session. The connections are stored as follows:
//
//	{
//	  "connections": {
//	    "connection-name": ["url", "driver"]
//	  },
//	  "current": "connection-name"
//	}
func CreateConnection(c *gin.Context) {
	var (
		url    string = c.PostForm("db-url")
		name   string = c.PostForm("db-conn-name")
		driver string = c.PostForm("db-driver")
	)

	session := sessions.Default(c)
	var connections map[string][2]string

	session_bytes, ok := session.Get("connections").([]byte)
	if !ok {
		connections = make(map[string][2]string)
	} else {
		if err := json.Unmarshal(session_bytes, &connections); err != nil {
			fmt.Println(err)
		}
	}

	// This is stupid as shit, but prevents a duplicate name from being
	// created.
	for true {
		var dupe bool = false
		for n := range connections {
			if n == name {
				name = fmt.Sprintf("%s (copy)", name)
				dupe = true
				break
			}
		}
		if !dupe {
			break
		}
	}

	connections[name] = [2]string{url, driver}

	conn_bytes, err := json.Marshal(connections)
	if err != nil {
		fmt.Println(err)
	}

	session.Set("connections", []byte(conn_bytes))
	session.Set("current", name)
	session.Save()

	html := templates.ConnectionsList(connections, name)
	html += TableTree(c)
	html += ViewTree(c)
	html += EnumTree(c)
	html += RoutineTree(c)

	c.String(200, html)
}

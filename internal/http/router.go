package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azpect3120/Web-Database-Viewer/internal/database"
	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Populate the server with routes
func populate(web, api *gin.RouterGroup) {
	web.GET("/view", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	api.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"time":   time.Now().String(),
		})
	})
	api.POST("/query", database.QueryCurrent)
	api.POST("/connections/test", database.TestConnectionURL)
	api.POST("/connections", database.CreateConnection)
	api.GET("/connections", func(c *gin.Context) {
		session := sessions.Default(c)
		connections_bytes, ok := session.Get("connections").([]byte)
		if !ok {
			c.JSON(200, gin.H{"connections": make(map[string][2]string), "current": "", "count": 0, "time": time.Now().String(), "status": 200})
			return
		}
		current := session.Get("current")

		var connections map[string][2]string
		if err := json.Unmarshal(connections_bytes, &connections); err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"connections": connections,
			"current":     current,
			"count":       len(connections),
			"time":        time.Now().String(),
			"status":      200,
		})
	})
	api.POST("/connections/delete", database.EditConnections)
	web.GET("/connections", func(c *gin.Context) {
		session := sessions.Default(c)
		connections_bytes, conn_ok := session.Get("connections").([]byte)
		current, curr_ok := session.Get("current").(string)

		var connections map[string][2]string
		if conn_ok || curr_ok {
			if err := json.Unmarshal(connections_bytes, &connections); err != nil {
				fmt.Println(err)
			}
		} else {
			connections = make(map[string][2]string)
		}

		html := templates.ConnectionsList(connections, current)
		c.String(200, html)
	})
	api.POST("/connections/connect", database.ChangeConnection)
	web.GET("/connections/tree/table", func(c *gin.Context) {
		c.String(200, database.TableTree(c))
	})
	web.GET("/connections/tree/table/columns", func(c *gin.Context) {
		c.String(200, database.TableColumns(c))
	})
	web.GET("/connections/tree/enum", func(c *gin.Context) {
		c.String(200, database.EnumTree(c))
	})
	web.GET("/connections/tree", func(c *gin.Context) {
		c.String(200, database.TableTree(c)+database.EnumTree(c))
	})
	web.GET("/query/auto", templates.ToggleQueryType)
	web.GET("/manager/open", templates.OpenManager)
	web.GET("/manager/hide", templates.HideManager)
}

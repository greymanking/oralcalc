package main

import (
	//"log"
	//"net/http"

	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var recio RecordIO = new(RecsSqlite)

func GetUser(c *gin.Context) string {
	return "jinyinuo"
}

func AddRecord(c *gin.Context) {
	var rec Record
	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(200, gin.H{"addrec": false})
		return
	}
	if err := recio.Add(GetUser(c), rec); err != nil {
		c.JSON(200, gin.H{"addrec": false})
		return
	}
	c.JSON(200, gin.H{"addrec": true, "data": rec})
}

func GetRecsAll(c *gin.Context) {
	recs := recio.All(GetUser(c))
	c.JSON(200, recs)
}

func GetRecsByKey(c *gin.Context) {
	recs := recio.Query(GetUser(c), c.Param("key"))
	if recs == nil {
		c.String(200, "[]")
	} else {
		c.JSON(200, recs)
	}
}

func GetDB(c *gin.Context) {
	c.File("./db.sqlite")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.Use(gin.Logger())

	//router.GET("/recs", GetRecsAll)

	router.GET("/greyman.db", GetDB)

	router.GET("/recs/:key", GetRecsByKey)

	router.POST("/rec", AddRecord)

	recio.Load()
	defer recio.Close()

	router.Run(":4000")
}

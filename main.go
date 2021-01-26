package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"ych.hub/common"
)

func main() {
	db := common.InitDb()
	defer db.Close()

	app := gin.Default()
	CollectRoute(app)

	panic(app.Run())
}

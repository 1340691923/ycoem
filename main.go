package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"ycoem/config"
	"ycoem/router"

	"github.com/kataras/iris"
)

func main() {
	config.Init()
	app := iris.New()
	router.Run(app)
	//task.Init()
	//defer task.Crontab.Stop()
	log.Fatal(app.Run(iris.Addr(config.Port)))
}

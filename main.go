package main

import (
	"fajar/clean/config"
	"fajar/clean/factory"
	"fajar/clean/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	e := echo.New()
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}

// lh3%*p6m
//dz9g86g@
//4fPd&*I3

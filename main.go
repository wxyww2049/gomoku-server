package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomoku-server/dao"
	"gomoku-server/router"
	"gomoku-server/service"
)

func main() {
	engine := gin.Default()
	dao.Setup()
	service.Setup()
	router.Setup(engine)
	err := engine.Run(fmt.Sprintf(":%v", "5521"))
	if err != nil {
		panic(err)
	}
	//user := po.User{1, "wxy", "2020"}
	//service.Test(&user)

}

package main

import (
	"Termbin/config"
	"Termbin/dao"
	"Termbin/router"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	_ = r.Run(":80")
}

func init() {
	rand.Seed(time.Now().UnixNano())
	config.InitConfig()
	dao.InitDB()
}

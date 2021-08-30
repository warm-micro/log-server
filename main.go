package main

import (
	"wm/logger/db"
	"wm/logger/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Connect()
}

func main() {
	r := gin.Default()
	routers.SetUpRouter(r)
	r.Run(":50051")
}

package main

import (
	"server/pkg/reader"
	"server/pkg/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	reader.InitVerses()

}

func main() {
	const PORT = ":8080"
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})

	r.GET("/getTodayVerse", routes.GetTodayVerse)

	r.Run(PORT)
}

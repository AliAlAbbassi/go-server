package main

import (
	"fmt"
	"log"

	"github.com/alialabbassi/go-server/handler"
	"github.com/alialabbassi/go-server/repository"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type key int

const (
	requestIDKey key = 0
)

var (
	Version      string = ""
	GitTag       string = ""
	GitCommit    string = ""
	GitTreeState string = ""
	listenAddr   string
	healthy      int32
)

func main() {
	fmt.Println("setup database connections")
	if err := repository.SetupRepo(); err != nil {
		log.Fatal(err)
	}
	defer repository.CloseRepo()

	fmt.Println("handling roures")

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/user", handler.UserGetHandler)
	router.POST("/user", handler.UserPostHandler)
	router.PUT("/user", handler.UserPutHandler)
	router.DELETE("/user/:id", handler.UserDeleteHandler)

	router.Run(":8080")

}

package main

import (
	"flag"

	"github.com/damianbalanz/api-hotel/api"
	"github.com/gin-gonic/gin"
)

func main() {

	listenAddr := flag.String("listenAddr", ":8080", "Server runing on port 8080")
	flag.Parse()

	app := gin.New()
	apiV1 := app.Group("/api/v1")

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola mundo",
		})
	})
	app.GET("/foo", handleFoo)
	apiV1.GET("/user", api.HandleGetUsers)

	app.Run(*listenAddr)
}

func handleFoo(c *gin.Context) {
	c.JSON(200, gin.H{
		"a": "a",
	})
}

// func handleUser(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"user": "dami",
// 	})
// }

// func runGinServer(config util.Config, store db.Store) {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal().Msgf("cannot create server: %s", err)
// 	}

// 	err = server.Start(fmt.Sprintf(":%d", config.Port))
// 	if err != nil {
// 		log.Fatal().Msgf("cannot start server: %s", err)
// 	}
// }

package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.New()

	app.GET("/", handleFoo)

	app.Run(":8080")
}
func handleFoo(c *gin.Context) {
	c.JSON(200, gin.H{
		"a": "a",
	})
}

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

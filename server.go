package main

import (
	"demo/graph"
	"demo/graph/common"
	"demo/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	customCtx := &common.CustomContext{
		Database: db,
	}

	return common.CreateContext(customCtx, h)
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()
	r.Static("/_nuxt", "./frontend/calendar/dist/_nuxt")
	r.Static("/images", "./frontend/calendar/dist/images")

	//r.Static("/home", "./frontend/calendar/dist/")

	r.LoadHTMLGlob("./frontend/calendar/dist/*.html")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/query", graphqlHandler())
	r.GET("/ide", playgroundHandler())
	r.Run()
}

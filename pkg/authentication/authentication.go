package authentication

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func DatabaseConnectionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_CONNECTION_STRING"))

		if err != nil {
			log.Print("Failed to connect to database.")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error"})
		}

		ctx.Set("persistence", conn)

		ctx.Next()
	}
}

// NewIdentityProvider designed to be used in a main package to unit the API module.
func NewIdentityProvider() *gin.Engine {
	router := gin.New()

	router.Use(DatabaseConnectionMiddleware())

	router.GET("/health_check", healthCheckHandler)
	router.GET("/users", getUsersHandler)

	return router
}

func healthCheckHandler(ctx *gin.Context) {
	log.Print("Health check.")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

type user struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func getUsersHandler(ctx *gin.Context) {
	conn, ok := ctx.MustGet("persistence").(*pgx.Conn)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error"})
	}

	users := []user{}
	rows, _ := conn.Query(context.Background(), "SELECT username, name FROM users")

	for rows.Next() {
		var username string
		var name string
		err := rows.Scan(&username, &name)
		if err != nil {
			log.Print(err)
		}

		user := user{Username: username, Name: name}
		users = append(users, user)
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

package main

import (
	"learn/internal/config"
	"learn/internal/controllers"
	"learn/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	config.LoadEnv()
}

func main() {
	var databaseConfig *config.Database
	var executionError error

	databaseConfig, executionError = config.LoadDatabase()
	if executionError != nil {
		log.Fatalf("Error loading database: %v", executionError)
		return
	}

	var databasePool *pgxpool.Pool
	databasePool, executionError = database.Connect(databaseConfig.Host, databaseConfig.Port, databaseConfig.Name, databaseConfig.User, databaseConfig.Password, databaseConfig.Timezone)
	if executionError != nil {
		log.Fatalf("Error connecting to database: %v", executionError)
		return
	}

	defer databasePool.Close()

	var router = gin.Default()

	router.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/user/index", controllers.UserIndex(databasePool))
	router.GET("/api/user/show/:id", controllers.UserShow(databasePool))
	router.POST("/api/user/store", controllers.UserStore(databasePool))
	router.PUT("/api/user/update/:id", controllers.UserUpdate(databasePool))
	router.DELETE("/api/user/delete/:id", controllers.UserDelete(databasePool))

	executionError = router.Run(":" + os.Getenv("PORT"))
	if executionError != nil {
		return
	}

	log.Println("Server started")
}

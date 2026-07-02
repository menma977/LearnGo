package controllers

import (
	"learn/internal/config"
	"learn/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserBodyStore struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserBodyUpdate struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func UserIndex(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(context *gin.Context) {
		users, err := repositories.UserAll(pool)
		if err != nil {
			config.InternalServerError(context, err.Error())
			return
		}

		config.Success(context, gin.H{
			"users": users,
		})
	}
}

func UserShow(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		user, err := repositories.UserFindByUuid(pool, id)
		if err != nil {
			config.InternalServerError(context, err.Error())
			return
		}

		config.Success(context, gin.H{
			"user": user,
		})
	}
}

func UserStore(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(context *gin.Context) {
		var userBody UserBodyStore
		if err := context.ShouldBindJSON(&userBody); err != nil {
			config.BadRequest(context, err.Error())
			return
		}

		user, err := repositories.UserCreate(pool, userBody.Name, userBody.Username, userBody.Email, userBody.Password)
		if err != nil {
			config.InternalServerError(context, err.Error())
			return
		}

		config.Success(context, gin.H{
			"user": user,
		})
	}
}

func UserUpdate(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		var userBody UserBodyUpdate
		if err := context.ShouldBindJSON(&userBody); err != nil {
			config.BadRequest(context, err.Error())
			return
		}

		existingUser, err := repositories.UserFindByUuid(pool, id)
		if err != nil {
			config.InternalServerError(context, err.Error())
			return
		}

		hasChanges := false
		if userBody.Name != "" && userBody.Name != existingUser.Name {
			hasChanges = true
		}
		if userBody.Username != "" && userBody.Username != existingUser.Username {
			hasChanges = true
		}

		if !hasChanges {
			config.Success(context, gin.H{
				"message": "Nothing to change",
			})
			return
		}

		name := existingUser.Name
		if userBody.Name != "" {
			name = userBody.Name
		}

		username := existingUser.Username
		if userBody.Username != "" {
			username = userBody.Username
		}

		user, err := repositories.UserEdit(pool, existingUser.ID, name, username, existingUser.Email, existingUser.Password)
		if err != nil {
			config.InternalServerError(context, err.Error())
			return
		}

		config.Success(context, gin.H{
			"user": user,
		})
	}
}

func UserDelete(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		existingUser, err := repositories.UserFindByUuid(pool, id)
		if err != nil {
			config.InternalServerError(context, err.Error())
			return
		}

		user, err := repositories.UserRemove(pool, existingUser.ID)
		if err != nil {
			config.InternalServerError(context, err.Error())
			return
		}

		config.Success(context, gin.H{
			"user": user,
		})
	}
}

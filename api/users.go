package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "user-handler/db/sqlc"
	"user-handler/helpers"

	"github.com/gin-gonic/gin"
)

type Users struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) CreateUsers(ctx *gin.Context) {
	var user Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.RespondWithJson(ctx, http.StatusBadRequest, "UserName and Password are required field", "error", nil)
		return
	}
	createdUser, err := server.Db.CreateUsers(ctx, db.CreateUsersParams{
		Username: sql.NullString{
			String: user.UserName,
			Valid:  true,
		},
		Password: sql.NullString{
			String: user.Password,
			Valid:  true,
		},
	})
	if err != nil {
		fmt.Println(err)
		helpers.RespondWithJson(ctx, http.StatusInternalServerError, "Something went wrong", "error", err.Error())
		return
	}
	strUserID, _ := createdUser.LastInsertId()
	helpers.RespondWithJson(ctx, 201, "The requested User created successfully", "success", strUserID)
}

func (server *Server) GetUsers(ctx *gin.Context) {
	userName := ctx.Query("username")
	if userName == "" {
		users, err := server.Db.GetAllUsers(ctx)
		if err != nil && err != sql.ErrNoRows {
			helpers.RespondWithJson(ctx, http.StatusInternalServerError, "Someting went wrong", "error", nil)
			return
		}
		helpers.RespondWithJson(ctx, http.StatusOK, "The requested users fetched successfully", "success", users)
		return
	}

	user, err := server.Db.GetUserByName(ctx, sql.NullString{
		String: userName,
		Valid:  userName != "",
	})
	if err != nil {
		if err == sql.ErrNoRows {
			helpers.RespondWithJson(ctx, http.StatusNotFound, "The requested user doesn't exists!!", "error", nil)
			return
		}
		helpers.RespondWithJson(ctx, http.StatusInternalServerError, "Something went wrong", "error", nil)
		return
	}
	helpers.RespondWithJson(ctx, http.StatusOK, "The requested users fetched successfully", "success", user)
}

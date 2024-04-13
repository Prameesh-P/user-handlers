package api

import (
	db "user-handler/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Db     db.Queries
}

func (s *Server) NewServer() {
	router := gin.Default()

	user := router.Group("/users")
	{
		user.POST("", s.CreateUsers)
		user.GET("", s.GetUsers)
	}
	router.Run(":8080")
}

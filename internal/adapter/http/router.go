package httpserver

import userhandler "sotoon/internal/adapter/http/handler/user"

func (r *Server) SetUpAPIRoutes(userH *userhandler.Handler) {
	router := r.engine

	{
		users := router.Group("users")

		users.POST("", userH.CreateUser)
	}

}

package routes

import (
	"ddd/infrastructure/eloquent"
	"ddd/interfaces"
	"ddd/interfaces/middlewares"
	"ddd/utils/helpers"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RegisterRoutes : Register Application Routes
func RegisterRoutes(port string) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return helpers.Respond(c, map[string]interface{}{
			"message": "Welcome to app",
		}, http.StatusOK, true)
	})

	repos, err := eloquent.NewRepositories()

	// Routes V1
	{
		api := e.Group("/api/v1")

		if err != nil {
			helpers.HandleErrors(err)
		}

		// Auth Routes
		{
			auth := interfaces.NewAuthHandler(repos.User)

			api.POST("/login", auth.Login)
			api.POST("/register", auth.Register)
			api.POST("/logout", auth.Logout, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.POST("/me", auth.Me, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.POST("/refresh", auth.Refresh, middlewares.CheckTokenIsNotInvalidated())
		}

		// Websockets
		{
			messages := interfaces.NewMessagesHandler(repos.Message)

			api.GET("/ws", messages.HandleConnections)
			api.GET("/messages-by-channel/:channel", messages.ListByChannelName)

			go interfaces.HandleWSMessages()
		}

		{
			users := interfaces.NewUserHandler(repos.User)

			api.GET("/users", users.Index, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.POST("/users", users.Create, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.GET("/users/:id", users.Show, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.PATCH("/users/:id", users.Update, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.DELETE("/users/:id", users.Delete, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
		}

		{
			groups := interfaces.NewGroupHandler(repos.Group)

			api.GET("/groups", groups.Index, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.POST("/groups", groups.Create, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.GET("/groups/:id", groups.Show, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.PATCH("/groups/:id", groups.Update, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
			api.DELETE("/groups/:id", groups.Delete, middlewares.UserMustBeAuthenticated(), middlewares.CheckTokenIsNotInvalidated())
		}
	}

	fmt.Println(
		fmt.Sprintf("Server started as http://127.0.0.1:%s", port),
	)

	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", port),
		),
	)
}

package routes

import (
	"e-complaint-api/controllers/admin"
	"e-complaint-api/controllers/user"
	"e-complaint-api/middlewares"
	"os"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AdminController *admin.AdminController
	UserController  *user.UserController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	var jwt = echojwt.JWT([]byte(os.Getenv("JWT_SECRET")))

	// Route For Super Admin
	superAdmin := e.Group("/api/v1")
	superAdmin.Use(jwt, middlewares.IsSuperAdmin)
	superAdmin.POST("/admins", r.AdminController.CreateAccount)

	// Route For Admin
	admin := e.Group("/api/v1")
	admin.POST("/admins/login", r.AdminController.Login)

	// Route For User
	user := e.Group("/api/v1")
	user.POST("/users/login", r.UserController.Login)
	user.POST("/users/register", r.UserController.Register)

	// Route For Admin and User

	// Route For Public
}

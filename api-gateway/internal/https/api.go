package https

import (
	"api-gateway/internal/https/handler"
	rlimit "api-gateway/internal/https/middleware/rate-limiting"
	"api-gateway/internal/service"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Youth Union API Gateway
// @version 1.0
// @description API Gateway for the Youth Union microservice platform.
// @host localhost:9000
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(service *service.ServiceRepositoryClient, port int) *http.Server {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(rlimit.RateLimitWithIP())

	apiHandler := handler.NewApiHandler(service)

	// Main API group
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", apiHandler.CreateUser)             // Register new user
			auth.POST("/login", apiHandler.GetUserByIdentifier)       // Login user
			auth.PATCH("/change-password", apiHandler.ChangePassword) // Change password
		}

		users := api.Group("/users")
		{
			users.GET("/", apiHandler.ListUsers)        // Get all users
			users.GET("/:id", apiHandler.GetUserById)   // Get user by ID
			users.PUT("/:id", apiHandler.UpdateUser)    // Update user
			users.DELETE("/:id", apiHandler.DeleteUser) // Delete user

			users.POST("/:id/roles", apiHandler.AssignRoleToUser)              // Assign role
			users.GET("/:id/roles", apiHandler.ListUserRoles)                  // List roles
			users.DELETE("/:id/roles/:role_id", apiHandler.RemoveRoleFromUser) // Remove role
		}

		roles := api.Group("/roles")
		{
			roles.POST("/", apiHandler.CreateRole)      // Create role type
			roles.GET("/", apiHandler.ListRoles)        // List all roles
			roles.GET("/:id", apiHandler.GetRoleById)   // Get role by ID
			roles.PUT("/:id", apiHandler.UpdateRole)    // Update role
			roles.DELETE("/:id", apiHandler.DeleteRole) // Delete role
		}
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	address := fmt.Sprintf(":%d", port)
	srv := &http.Server{
		Addr:      address,
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv
}

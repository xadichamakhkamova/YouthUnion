package https

import (
	"api-gateway/internal/https/handler"
	"api-gateway/internal/https/middleware/authorization"
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

	authMiddleware := &authorization.AuthMiddleware{}

	r.Use(authMiddleware.MiddleWare())
	r.Use(rlimit.RateLimitWithIP())

	apiHandler := handler.NewApiHandler(service)

	// Main API group
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", apiHandler.GetUserByIdentifier)       // Login user
		}

		users := api.Group("/users")
		{
			users.POST("/", apiHandler.CreateUser)
			users.GET("/", apiHandler.ListUsers)
			users.GET("/:id", apiHandler.GetUserById)
			users.PUT("/:id", apiHandler.UpdateUser)
			users.DELETE("/:id", apiHandler.DeleteUser)

			users.PATCH("/:id/password", apiHandler.ChangePassword)   


			users.POST("/:id/roles", apiHandler.AssignRoleToUser)
			users.GET("/:id/roles", apiHandler.ListUserRoles)
			users.DELETE("/:id/roles/:role_id", apiHandler.RemoveRoleFromUser)
		}

		roles := api.Group("/roles")
		{
			roles.POST("/", apiHandler.CreateRole)      // Create role type
			roles.GET("/", apiHandler.ListRoles)        // List all roles
			roles.GET("/:id", apiHandler.GetRoleById)   // Get role by ID
			roles.PUT("/:id", apiHandler.UpdateRole)    // Update role
			roles.DELETE("/:id", apiHandler.DeleteRole) // Delete role
		}

		event := api.Group("/events")
		{
			event.POST("/", apiHandler.CreateEvent)
			event.PUT("/:id", apiHandler.UpdateEvent)
			event.GET("/:id", apiHandler.GetEvent)
			event.GET("/", apiHandler.ListEvents)
			event.DELETE("/:id", apiHandler.DeleteEvent)

			//event.GET("/:id/participants", apiHandler.ListParticipants)
		}

		teams := api.Group("/teams")
		{
			teams.POST("/", apiHandler.CreateTeam)
			teams.PUT("/:id", apiHandler.UpdateTeam)
			teams.GET("/event/:event_id", apiHandler.GetTeamsByEvent)

			teams.DELETE("/:team_id/members/:user_id", apiHandler.RemoveTeamMember)
			teams.GET("/:team_id/members", apiHandler.GetTeamMembers)

			// teams.POST("/:team_id/invite", apiHandler.InviteMember)
			// teams.POST("/:team_id/respond", apiHandler.RespondInvite)
		}

		scoring := api.Group("/scoring")
		{
			scoring.POST("/give-score", apiHandler.GiveScore)

			scoring.GET("/event/:event_id", apiHandler.GetScoresByEvent)
			scoring.GET("/user/:user_id", apiHandler.GetScoresByUser)
			scoring.GET("/team/:team_id", apiHandler.GetScoresByTeam)

			scoring.GET("/ranking", apiHandler.GetGlobalRanking)
		}
		notifications := api.Group("/notifications")
		{
			notifications.POST("/send", apiHandler.SendNotification)
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

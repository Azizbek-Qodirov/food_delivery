package api

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth-service/api/docs"
	"auth-service/api/handlers"
	"auth-service/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(h *handlers.HTTPHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", h.Register)
	router.POST("/confirm-registration", h.ConfirmRegistration)
	router.POST("/login", h.Login)
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/recover-password", h.RecoverPassword)

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile)

	router.GET("/user/:id", h.GetByID)

	router.GET("")

	return router
}

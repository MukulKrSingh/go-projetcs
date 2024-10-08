package controllers

import (
	"gin-demo/internal/utils"
	"gin-demo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	as *services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		as: authService,
	}
}

func (a *AuthController) InitRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.GET("/login", a.Login())
	routes.POST("/register", a.Register())

}
func (a *AuthController) Register() gin.HandlerFunc {

	type RegisterBody struct {
		Email    string `json:"email" bindings:"required"`
		Password string `json:"password" bindings:"required"`
	}
	return func(c *gin.Context) {
		var registerBody RegisterBody

		if err := c.BindJSON(&registerBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := a.as.RegisterService(&registerBody.Email, &registerBody.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Registered",
			"email":   user,
		})

	}
}

func (a *AuthController) Login() gin.HandlerFunc {
	type LoginBody struct {
		Email    string `json:"email" bindings:"required"`
		Password string `json:"password" bindings:"required"`
	}

	return func(c *gin.Context) {
		var loginBody LoginBody

		if err := c.BindJSON(&loginBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := a.as.LoginService(&loginBody.Email, &loginBody.Password)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		token, tokenErr := utils.GenerateToken(user.Email, user.Id)

		if tokenErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": tokenErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": user,
			"token":   token,
		})

	}
}

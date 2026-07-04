package endpoints

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/repository"
)

type AuthHandler struct {
	repo repository.UserRepository
}

func NewAuthHandler(repo repository.UserRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body entities.RegisterRequest true "Registration"
// @Success 201 {object} entities.AuthResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req entities.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.repo.Register(req.Email, req.Password)
	if err != nil {
		slog.Error("failed to register", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, entities.AuthResponse{User: user, Token: ""})
}

// Login godoc
// @Summary Log in a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body entities.LoginRequest true "Login"
// @Success 200 {object} entities.AuthResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req entities.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.repo.Login(req.Email, req.Password)
	if err != nil {
		slog.Error("failed to login", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, entities.AuthResponse{User: user, Token: ""})
}

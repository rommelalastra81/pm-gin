package controllers

import (
	"net/http"

	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register handles new user registration.
// POST /api/auth/register
func RegisterUser(c *gin.Context) {
	var req dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var existingUser models.Users
	if result := config.DB.Where("email = ?", req.Email).First(&existingUser); result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := models.Users{
		FullName: req.FullName,
		Email:    req.Email,
		Password: string(hashedPassword),
		JobRole:  req.JobRole,
	}

	if result := config.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, dto.RegisterSuccessful{
		FullName: user.FullName,
		Email:    user.Email,
		JobRole:  user.JobRole,
		Message:  "Registered successfully",
	})
}

// GetUserProfile returns the authenticated user's profile.
// GET /api/auth/profile
func GetUserProfile(c *gin.Context) {
	// Retrieve user claims set by the AuthMiddleware
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtClaims := claims.(*config.JWTClaims)

	var user models.Users
	if result := config.DB.First(&user, jwtClaims.UserID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
		JobRole:  user.JobRole,
	})
}

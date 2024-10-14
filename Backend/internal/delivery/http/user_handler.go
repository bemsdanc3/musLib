package http

import (
	"Backend/internal/entities"
	"Backend/internal/usecases"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateUSer(&user); err != nil {
		c.JSON(500, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User created successfully!"})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.usecase.GetUserByID(userID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(200, user)
}

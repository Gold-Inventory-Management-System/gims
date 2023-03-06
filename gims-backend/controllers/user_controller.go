package controllers

import (
	"net/http"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gin-gonic/gin"
)

type registerAdminBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type UserController struct {
	userUseCase domains.UserUseCase
}

func NewUserController(uu domains.UserUseCase) *UserController {
	return &UserController{
		userUseCase: uu,
	}
}

func (uc *UserController) RegisterAdmin(c *gin.Context) {
	body := new(registerAdminBody)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	newAdmin := &models.User{
		Username: body.Username,
		Password: body.Password,
	}
	if err := uc.userUseCase.RegisterAdmin(newAdmin, body.Secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (uc *UserController) SignIn(c *gin.Context) {
	body := new(models.User)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, token, err := uc.userUseCase.SignIn(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username":    user.Username,
		"role":        user.Role,
		"accesstoken": token,
	})
}

func (uc *UserController) Register(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Only admin can register account",
		})
		return
	}
	body := new(models.User)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if body.Role != "owner" && body.Role != "employee" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "role must be owner or employee",
		})
		return
	}
	err := uc.userUseCase.Register(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (uc *UserController) TestJWT(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"role":     role,
	})
}

func (uc *UserController) QueryAllUser(c *gin.Context){
	role, _ := c.Get("role")
	if role != "admin"{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "only admin can do this",
		})
		return
	}
	users, err := uc.userUseCase.QueryAllUser()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (uc *UserController) RenameUsername(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "only admin can do this",
		})
		return
	}
	oldUsername := c.Query("oldusername")
	newUsername := c.Query("newusername")
	if err := uc.userUseCase.RenameUsername(oldUsername, newUsername); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (uc *UserController) UpdatePassword(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "only admin can do this",
		})
		return
	}
	username := c.Query("username")
	password := c.Query("password")
	if err := uc.userUseCase.UpdatePassword(username, password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "only admin can do this",
		})
		return
	}
	username := c.Param("username")
	if err := uc.userUseCase.DeleteUser(username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
package handler

import (
	"kuwa72/sample-gorm-txdb-testing/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

func (h *UserHandler) CreateUser(r *gin.Engine) *gin.Engine {
	r.POST("/user/add", func(c *gin.Context) {
		var req CreateUserRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		ret, err := usecase.CreateUser(h.DB, req.Name, req.Email, req.Password)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, ret)
	})
	return r
}

type UpdateUserRequest struct {
	Name     string
	Email    string
	Password string
}

func (h *UserHandler) UpdateUser(r *gin.Engine) *gin.Engine {
	r.POST("/user/update", func(c *gin.Context) {
		var req CreateUserRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		ret, err := usecase.UpdateUser(h.DB, req.Name, req.Email, req.Password)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, ret)
	})
	return r
}

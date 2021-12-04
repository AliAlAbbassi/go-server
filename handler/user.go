package handler

import (
	"net/http"
	"strconv"

	"github.com/alialabbassi/go-server/model"
	"github.com/alialabbassi/go-server/repository"
	"github.com/gin-gonic/gin"
)

//UserPostHandler - handle User post request
func UserPostHandler(c *gin.Context) {
	UserRepo := repository.GetUserRepository()
	var User model.User
	if err := c.ShouldBindJSON(&User); err == nil {
		_, err = UserRepo.Insert(User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//UserGetHandler - handle User get request
func UserGetHandler(c *gin.Context) {
	UserRepo := repository.GetUserRepository()
	Users, err := UserRepo.Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, convertListToArray(Users))
}

//UserPutHandler - handle User put request
func UserPutHandler(c *gin.Context) {
	UserRepo := repository.GetUserRepository()
	var User model.User
	if err := c.ShouldBindJSON(&User); err == nil {
		_, err = UserRepo.Update(User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//UserDeleteHandler - handle User delete request
func UserDeleteHandler(c *gin.Context) {
	UserRepo := repository.GetUserRepository()
	var User model.User
	User.UserID, _ = strconv.Atoi(c.Param("id"))
	_, err := UserRepo.Remove(User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}

package services

import (
	"crudApp/models"
	"crudApp/repositories"
	"crudApp/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Add User Function
func AddUserService(c *gin.Context) {
	log.Println("[UserService] User Add Rest Api Started !")

	var userRequest models.User
	var userCheck models.User
	var errCheckUser error
	var errUserAdd error

	// Bind The Request
	c.Bind(&userRequest)

	log.Printf("[UserService] User Add Request - %v", utils.JsonFormater(userRequest))

	// Check User Is Present Or Not
	errCheckUser, userCheck = repositories.FindUser(userRequest.Email)
	if errCheckUser != nil && !strings.Contains(fmt.Sprint(errCheckUser), utils.RECORD_NOT_FOUND_MESSAGE) {
		c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(errCheckUser), utils.STATUS_OK, nil))
		return
	} else {
		if userCheck != (models.User{}) {
			log.Printf("[UserService] User Check Response - %v", utils.JsonFormater(userCheck))

			c.JSON(http.StatusOK, utils.ResponsePrepare("User Details Alrady Present !", utils.STATUS_OK, nil))
			return
		} else {
			// Save User
			errUserAdd, userRequest = repositories.CreateUser(userRequest)

			if errUserAdd != nil {
				c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(errCheckUser), utils.STATUS_OK, nil))
				return
			}

			log.Printf("[UserService] User AddResponse - %v", utils.JsonFormater(userRequest))

			c.JSON(http.StatusOK, utils.ResponsePrepare("User Added Successfully !", utils.STATUS_OK, userRequest))
			return
		}
	}
}

// Remove User Record Using Id
func RemoveUserService(c *gin.Context, ch chan bool) {
	log.Println("[UserService] User Remove Rest Api Started !")
	id := c.Param("id")
	log.Printf("[UserService] User Remove Id - %v", id)
	userId, _ := strconv.Atoi(id)

	var userCheck models.User
	var errUserCheck error

	// Check Record Is Present Or Not
	errUserCheck, userCheck = repositories.FindUserById(uint(userId))
	if errUserCheck == nil {
		// Remove Record
		err := repositories.RemoveUser(userCheck.ID)
		if err != nil {
			c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(err), utils.STATUS_OK, nil))
			ch <- true
			return
		}

		c.JSON(http.StatusOK, utils.ResponsePrepare(utils.STATUS_MESSAGE, utils.STATUS_OK, nil))
		ch <- true
		return
	} else if strings.Contains(fmt.Sprint(errUserCheck), utils.RECORD_NOT_FOUND_MESSAGE) {
		c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(errUserCheck), utils.STATUS_OK, nil))
		ch <- true
		return
	} else {
		c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(errUserCheck), utils.STATUS_OK, nil))
		ch <- true
	}
}

// Get User By Id
func GetUserByIdService(c *gin.Context, ch chan bool) {
	log.Println("[UserService] User Get By Id Rest Api Started !")
	id := c.Param("id")
	log.Printf("[UserService] User Get Id - %v", id)
	userId, _ := strconv.Atoi(id)

	// Get User Details
	err, user := repositories.FindUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(err), utils.STATUS_OK, nil))
		ch <- true
		return
	}

	c.JSON(http.StatusOK, utils.ResponsePrepare(utils.STATUS_MESSAGE, utils.STATUS_OK, user))
	ch <- true
}

// Update User Bu Id
func UpdateUserById(c *gin.Context, ch chan bool) {
	log.Printf("[UserService] User Update Rest Api Started !")

	// Get Id From Parm
	id := c.Param("id")
	log.Printf("[UserService] User Update Id - %v", id)
	userId, _ := strconv.Atoi(id)

	// Get Request
	var request models.User
	c.Bind(&request)

	log.Printf("[UserService] User Update Request - %v", utils.JsonFormater(request))

	// Get User
	errGetData, userDetails := repositories.FindUserById(uint(userId))

	if errGetData != nil && !strings.Contains(fmt.Sprint(errGetData), utils.RECORD_NOT_FOUND_MESSAGE) {
		c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(errGetData), utils.STATUS_OK, nil))
		ch <- true
		return
	} else {
		if userDetails != (models.User{}) {
			userDetails.FirstName = request.FirstName
			userDetails.LastName = request.LastName

			// Update User
			updateError, userUpdateResponse := repositories.UpdateUser(userDetails)
			if updateError != nil {
				c.JSON(http.StatusOK, utils.ResponsePrepare(fmt.Sprint(errGetData), utils.STATUS_OK, nil))
				ch <- true
				return
			}
			c.JSON(http.StatusOK, utils.ResponsePrepare("User Details Updated", utils.STATUS_OK, userUpdateResponse))
			ch <- true
			return
		} else {
			c.JSON(http.StatusOK, utils.ResponsePrepare("User Not Found", utils.STATUS_OK, nil))
			ch <- true
			return
		}
	}
}

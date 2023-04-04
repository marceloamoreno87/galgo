package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	UserService "github.com/marceloamoreno87/galgo-api/api/users/services"
	UserValidator "github.com/marceloamoreno87/galgo-api/api/users/validators"
	Response "github.com/marceloamoreno87/galgo-api/helpers"
)

func Index(ctx *gin.Context) {
	users, err := UserService.Index(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func Store(ctx *gin.Context) {
	body, err := UserValidator.ValidateStore(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	user, err := UserService.Store(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func Show(ctx *gin.Context) {
	id, err := UserValidator.ValidateShow(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	users, err := UserService.Show(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func Update(ctx *gin.Context) {
	id, body, err := UserValidator.ValidateUpdate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	users, err := UserService.Update(body, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func Destroy(ctx *gin.Context) {
	id, err := UserValidator.ValidateDestroy(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	users, err := UserService.Destroy(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

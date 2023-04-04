package services

import (
	"github.com/gin-gonic/gin"
	UserModel "github.com/marceloamoreno87/galgo-api/api/users/models"
	UserRepository "github.com/marceloamoreno87/galgo-api/api/users/repositories"
	UserTypes "github.com/marceloamoreno87/galgo-api/api/users/types"
	Helper "github.com/marceloamoreno87/galgo-api/helpers"
)

func Index(ctx *gin.Context) (res *Helper.Pagination, err error) {
	page, PageSize := Helper.GetQueryData(ctx)
	list := Helper.Pagination{}
	list.PageSize = PageSize
	list.Page = page
	if res, err = UserRepository.Index(list); err != nil {
		return
	}
	return
}

func Store(body UserTypes.AddUserRequestBody) (res *UserModel.User, err error) {
	user := &UserModel.User{
		Name:     body.Name,
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}
	passwordHashed, err := Helper.HashPassword(body.Password)
	if err != nil {
		return
	}
	user.Password = passwordHashed
	if res, err = UserRepository.Store(user); err != nil {
		return
	}
	return
}

func Show(id string) (res *UserModel.User, err error) {
	if res, err = UserRepository.Show(id); err != nil {
		return
	}
	return
}

func Update(body UserTypes.UpdateUserRequestBody, id string) (first *UserModel.User, err error) {
	first, err = Show(id)
	if err != nil {
		return
	}
	first.Name = body.Name
	first.Username = body.Username
	first.Email = body.Email
	first.Password = body.Password
	if _, err = UserRepository.Update(first, id); err != nil {
		return
	}
	return
}

func Destroy(id string) (first *UserModel.User, err error) {
	first, err = Show(id)
	if err != nil {
		return
	}
	if _, err = UserRepository.Destroy(first); err != nil {
		return
	}
	return
}

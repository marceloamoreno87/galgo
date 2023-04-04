package validators

import (
	"github.com/gin-gonic/gin"
	UserTypes "github.com/marceloamoreno87/galgo-api/api/users/types"
	Helper "github.com/marceloamoreno87/galgo-api/helpers"
)

func ValidateStore(ctx *gin.Context) (body UserTypes.AddUserRequestBody, err error) {
	body = UserTypes.AddUserRequestBody{}
	if err = ctx.BindJSON(&body); err != nil {
		return
	}
	if err = Helper.Validate(&body); err != nil {
		return
	}
	return
}

func ValidateShow(ctx *gin.Context) (id string, err error) {
	uri := UserTypes.ParamIdUser{}
	id = ctx.Param("id")
	if err = ctx.BindUri(&uri); err != nil {
		return
	}
	return
}

func ValidateUpdate(ctx *gin.Context) (id string, body UserTypes.UpdateUserRequestBody, err error) {
	uri := UserTypes.ParamIdUser{}
	id = ctx.Param("id")
	if err = ctx.BindUri(&uri); err != nil {
		return
	}
	body = UserTypes.UpdateUserRequestBody{}
	if err = ctx.BindJSON(&body); err != nil {
		return
	}
	if err = Helper.Validate(body); err != nil {
		return
	}
	return
}

func ValidateDestroy(ctx *gin.Context) (id string, err error) {
	uri := UserTypes.ParamIdUser{}
	id = ctx.Param("id")
	if err = ctx.BindUri(&uri); err != nil {
		return
	}
	return
}

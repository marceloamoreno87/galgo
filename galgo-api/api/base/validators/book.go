package validators

import (
	"github.com/gin-gonic/gin"
	BookType "github.com/marceloamoreno87/galgo-api/api/books/types"
	Helper "github.com/marceloamoreno87/galgo-api/helpers"
)

func ValidateStore(ctx *gin.Context) (body BookType.AddBookRequestBody, err error) {
	body = BookType.AddBookRequestBody{}
	if err = ctx.BindJSON(&body); err != nil {
		return
	}
	if err = Helper.Validate(&body); err != nil {
		return
	}
	return
}

func ValidateShow(ctx *gin.Context) (id string, err error) {
	uri := BookType.ParamIdBook{}
	id = ctx.Param("id")
	if err = ctx.BindUri(&uri); err != nil {
		return
	}
	return
}

func ValidateUpdate(ctx *gin.Context) (id string, body BookType.UpdateBookRequestBody, err error) {
	uri := BookType.ParamIdBook{}
	id = ctx.Param("id")
	if err = ctx.BindUri(&uri); err != nil {
		return
	}
	body = BookType.UpdateBookRequestBody{}
	if err = ctx.BindJSON(&body); err != nil {
		return
	}
	if err = Helper.Validate(body); err != nil {
		return
	}
	return
}

func ValidateDestroy(ctx *gin.Context) (id string, err error) {
	uri := BookType.ParamIdBook{}
	id = ctx.Param("id")
	if err = ctx.BindUri(&uri); err != nil {
		return
	}
	return
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	BookService "github.com/marceloamoreno87/galgo-api/api/books/services"
	BookValidator "github.com/marceloamoreno87/galgo-api/api/books/validators"
	Response "github.com/marceloamoreno87/galgo-api/helpers"
	"github.com/marceloamoreno87/galgo-mail/src/driver/smtp"
	"github.com/marceloamoreno87/galgo-mail/src/mail"
	"github.com/marceloamoreno87/galgo-storage/src/driver/local"
	"github.com/marceloamoreno87/galgo-storage/src/storage"
)

func Index(ctx *gin.Context) {

	// TESTE EMAIL LIB
	message := smtp.MailMessageSmtp{
		To:      []string{"marceloamoreno87@gmail.com"},
		Cc:      []string{"marceloamoreno87@gmail.com"},
		Subject: "teste",
		From:    "marceloamoreno87@gmail.com",
		Body:    "<h1>BODE mano</h1>",
	}
	go mail.Send(&message)
	// TESTE EMAIL LIB

	books, err := BookService.Index(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func Store(ctx *gin.Context) {

	// TESTE STORAGE LIB
	arquivo, _ := ctx.FormFile("file")
	s := &local.StorageLocal{
		File: arquivo,
		Path: "./uploads/produtos/",
	}
	go storage.Save(s)
	// TESTE STORAGE LIB

	//download the file from folder using gin
	ctx.File("./uploads/produtos/")
	ctx.File("./uploads/produtos/" + arquivo.Filename)

	body, err := BookValidator.ValidateStore(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	book, err := BookService.Store(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": book})
}

func Show(ctx *gin.Context) {
	id, err := BookValidator.ValidateShow(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	books, err := BookService.Show(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func Update(ctx *gin.Context) {
	id, body, err := BookValidator.ValidateUpdate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	books, err := BookService.Update(body, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func Destroy(ctx *gin.Context) {
	id, err := BookValidator.ValidateDestroy(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	books, err := BookService.Destroy(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

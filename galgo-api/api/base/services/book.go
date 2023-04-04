package services

import (
	"github.com/gin-gonic/gin"
	Model "github.com/marceloamoreno87/galgo-api/api/books/models"
	BookRepository "github.com/marceloamoreno87/galgo-api/api/books/repositories"
	BookTypes "github.com/marceloamoreno87/galgo-api/api/books/types"
	"github.com/marceloamoreno87/galgo-api/helpers"
)

func Index(ctx *gin.Context) (res *helpers.Pagination, err error) {
	page, PageSize := helpers.GetQueryData(ctx)
	list := helpers.Pagination{}
	list.PageSize = PageSize
	list.Page = page
	if res, err = BookRepository.Index(list); err != nil {
		return
	}
	return
}

func Store(body BookTypes.AddBookRequestBody) (res *Model.Book, err error) {
	book := &Model.Book{
		Title:   body.Title,
		Content: body.Content,
	}
	if res, err = BookRepository.Store(book); err != nil {
		return
	}
	return
}

func Show(id string) (res *Model.Book, err error) {
	if res, err = BookRepository.Show(id); err != nil {
		return
	}
	return
}

func Update(body BookTypes.UpdateBookRequestBody, id string) (first *Model.Book, err error) {
	first, err = Show(id)
	if err != nil {
		return
	}
	first.Title = body.Title
	first.Content = body.Content
	if _, err = BookRepository.Update(first, id); err != nil {
		return
	}
	return
}

func Destroy(id string) (first *Model.Book, err error) {
	first, err = Show(id)
	if err != nil {
		return
	}
	if _, err = BookRepository.Destroy(first); err != nil {
		return
	}
	return
}

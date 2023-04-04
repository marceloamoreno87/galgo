package repositories

import (
	BookModel "github.com/marceloamoreno87/galgo-api/api/books/models"
	Database "github.com/marceloamoreno87/galgo-api/config"
	Paginate "github.com/marceloamoreno87/galgo-api/helpers"
)

func Index(list Paginate.Pagination) (*Paginate.Pagination, error) {
	var books []*BookModel.Book
	if err := Database.DB.Scopes(Paginate.Paginate(books, &list, Database.DB)).Find(&books).Error; err != nil {
		return nil, err
	}
	list.Rows = books
	return &list, nil
}

func Store(new *BookModel.Book) (book *BookModel.Book, err error) {
	if err = Database.DB.Create(&new).Error; err != nil {
		return
	}
	return
}

func Show(id string) (book *BookModel.Book, err error) {
	if err = Database.DB.First(&book, id).Error; err != nil {
		return
	}
	return
}

func Update(update *BookModel.Book, id string) (book *BookModel.Book, err error) {
	if err = Database.DB.Save(&update).Error; err != nil {
		return
	}
	return
}

func Destroy(destroy *BookModel.Book) (book *BookModel.Book, err error) {
	if err = Database.DB.Delete(&destroy).Error; err != nil {
		return
	}
	return
}

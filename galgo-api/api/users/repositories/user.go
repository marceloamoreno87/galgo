package repositories

import (
	UserModel "github.com/marceloamoreno87/galgo-api/api/users/models"
	Database "github.com/marceloamoreno87/galgo-api/config"
	Paginate "github.com/marceloamoreno87/galgo-api/helpers"
)

func Index(list Paginate.Pagination) (*Paginate.Pagination, error) {
	var users []*UserModel.User
	if err := Database.DB.Scopes(Paginate.Paginate(users, &list, Database.DB)).Find(&users).Error; err != nil {
		return nil, err
	}
	list.Rows = users
	return &list, nil
}

func Store(new *UserModel.User) (user *UserModel.User, err error) {
	if err = Database.DB.Create(&new).Error; err != nil {
		return
	}
	return
}

func Show(id string) (user *UserModel.User, err error) {
	if err = Database.DB.First(&user, id).Error; err != nil {
		return
	}
	return
}

func Update(update *UserModel.User, id string) (user *UserModel.User, err error) {
	if err = Database.DB.Save(&update).Error; err != nil {
		return
	}
	return
}

func Destroy(destroy *UserModel.User) (user *UserModel.User, err error) {
	if err = Database.DB.Delete(&destroy).Error; err != nil {
		return
	}
	return
}

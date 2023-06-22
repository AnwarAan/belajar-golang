// package service

// import (
// 	"errors"
// 	entity "module/helper"
// )

// type CategoryService struct {
// 	Respository repository.CategoryRepository
// }

// func (service CategoryRepository) get(id string) (*entity.Category, error) {
// 	category := service.Repository.FindById(id)
// 	if category == nil {
// 		return category, errors.New("not found")
// 	} else {
// 		return category, nil
// 	}
// }

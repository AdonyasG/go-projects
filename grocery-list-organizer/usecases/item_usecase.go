package usecases

import "github.com/AdonyasG/go-projects/grocery-list-organizer/entities"

type ItemUseCase interface {
    AddItem(name string, category string)(entities.Item, error)
    ListItemsByCategory() []entities.Item
    DeleteItem(id int) error
}
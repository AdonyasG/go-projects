package usecases

import (
	"errors"

	"github.com/AdonyasG/go-projects/grocery-list-organizer/entities"
)

type ItemUseCaseImpl struct {
	items []entities.Item
	nextID int
}

func NewItmeUseCase() ItemUseCase {
	return &ItemUseCaseImpl{
		items: []entities.Item{},
		nextID: 1,
	}
}

func (u *ItemUseCaseImpl) AddItem(name string, category string)(entities.Item, error) {
	item := entities.Item{ID: u.nextID, Name: name, Category: category}
	u.items = append(u.items, item )
	u.nextID++
	return item, nil
}

func (u *ItemUseCaseImpl) ListItemsByCategory()[]entities.Item {
	return u.items
}

func (u *ItemUseCaseImpl) DeleteItem(id int)error{

	for i, item := range u.items {
		if item.ID == id {
			u.items = append(u.items[:i], u.items[i+1:]...)
			return nil
		}
	}
	return errors.New("items not found")
}
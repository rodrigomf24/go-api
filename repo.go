package main

import "fmt"

var currentId int
var items Items

//Get some data
func init() {
	RepoCreateItem(Item{Name: "Gun1"})
	RepoCreateItem(Item{Name: "Gun2"})
}

func RepoFindItem(id int) Item {
	for _, i := range items {
		if i.Id == id {
			return i
		}
	}

	// return empty Item if not found
	return Item{}
}

func RepoCreateItem(i Item) Item {
	currentId += 1
	i.Id = currentId
	items = append(items, i)
	return i
}

func RepoDestroyItem(id int) error {
	for index, i := range items {
		if i.Id == id {
			items = append(items[:index], items[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Item with id of %d to delete", id)
}
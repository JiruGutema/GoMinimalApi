//Package models contains the models of necessary entities
package models


// Task type for our tasks
type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	IsCompleted bool `json:"isCompleted"`
}



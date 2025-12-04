
//Package models contains the models of necessary entities
package models

type User struct {
	ID int `json:"id"`
	FirstName string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	
}


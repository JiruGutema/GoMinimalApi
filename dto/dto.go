//Package dto provides a data tranfer object
package dto 

type Login struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

type Response struct {
	Status string `json:"status"`
	Body []any `json:"body"`
}

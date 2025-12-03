//Package data stores the data for the todo
package data

import (
	"gominimalapi/models"
)
//Tasks is the lists of the books
var Tasks = []models.Task {
{ID: 1, Name: "Task One", IsCompleted: false},
{ID: 2, Name: "Task Two", IsCompleted: true},
{ID: 3, Name: "Task Three", IsCompleted: false},
}

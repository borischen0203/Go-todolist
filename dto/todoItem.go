package dto

//TODO: Add request and response struct(Note:request should be formvalue type)
type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

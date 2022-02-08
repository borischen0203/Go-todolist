package dto

//TODO: Add request and response struct(Note:request should be formvalue type)
type TodoItemModel struct {
	Id          int    `gorm:"primary_key";`
	Description string `form:"description"`
	Completed   bool   `form:"completed"`
}

type TodoRequest struct {
	Descripion string `form:"description"`
}

type TodoUpdateRequest struct {
	Id        int  `param:"id" `
	Completed bool `form:"completed"`
}

type TodoResponse struct {
	Id int
}

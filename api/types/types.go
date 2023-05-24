package types

// contains the todo fields, that should be returned in a response
type TodoResponse struct {
	ID          uint   `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// defines the /todo/create
type CreateDTO struct {
	Task string `json:"task" validate:"required, min=3, max=150"`
}

// defiens the /todo/create response
type TodoCreateRespone struct {
	Todo *TodoResponse `json:"todo"`
}

// defines the todolist
type TodosResponse struct {
	Todos *[]TodoResponse `json:"todos"`
}

// defines the payload for the check todo
type CheckTodoDTO struct {
	Completed bool `json:"completed"`
}

// MsgResponse defined the message payload
type MsgResponse struct {
	Message string `json:"message"`
}

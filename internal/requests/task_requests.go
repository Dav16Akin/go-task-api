package requests

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"max=255"`
	Completed   bool   `json:"completed"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title" binding:"omitempty,min=3"`
	Description *string `json:"description" binding:"omitempty,max=255"`
	Completed   *bool   `json:"completed"`
}

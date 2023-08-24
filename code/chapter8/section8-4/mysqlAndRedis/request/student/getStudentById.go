package student

type GetStudentByIdParam struct {
	Id *int `json:"id" binding:"required,gte=0"`
}

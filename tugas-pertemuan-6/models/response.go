package models

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	InternalServerErrorResponse ErrorResponse = ErrorResponse{Error: "Internal server error"}
	NotFoundErrorResponse       ErrorResponse = ErrorResponse{Error: "Not found"}
	UnauthorizedErrorResponse   ErrorResponse = ErrorResponse{Error: "Unauthorized"}
	ForbiddenErrorResponse      ErrorResponse = ErrorResponse{Error: "Forbidden"}
)

type Pagination struct {
	CurrentPage uint `json:"current_page"`
	TotalPages  uint `json:"total_pages"`
	TotalItems  uint `json:"total_items"`
	PageSize    uint `json:"page_size"`
}

type GetAllStudentResponse struct {
	Students   []Student  `json:"students"`
	Pagination Pagination `json:"pagination"`
}

type GetStudentByIdResponse struct {
	Student
}

type CreateStudentResponse struct {
	Student
}

type LoginResponse struct {
	User  UserSafe `json:"user"`
	Token string   `json:"token"`
}

type RegisterResponse struct {
	UserSafe
}

type ProfileResponse struct {
	UserSafe
}

package models

type CreateStudentRequest struct {
	NIM      string `json:"nim"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Major    string `json:"major"`
	Semester int    `json:"semester"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

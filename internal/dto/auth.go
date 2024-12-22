package dto

type LoginRequest struct {
	Email    string `json:"email" example:"abc@mail.ru"`
	Password string `json:"password" example:"1234567"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserSessionResponse struct {
	ID        int32  `json:"id"`
	UserID    int32  `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

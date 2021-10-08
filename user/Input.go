package user

type RegisterUserInput struct {
	ID 			int
	Name 		string `json:"name" binding:"required"`
	NoKta 		string `json:"nokta" binding:"required"`
	Occupation  string `json:"occupation" binding:"required"`
	Email 		string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required"`
}
type LoginInput struct {
	Email 		string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required"`
}
type CheckEmailInput struct {
	Email 		string `json:"email" binding:"required,email"`
}
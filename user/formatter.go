package user

type UserFormatter struct {
	ID 			int    `json:"id" binding:"required"`
	Name 		string `json:"name" binding:"required"`
	NoKta 		string `json:"no_kta" binding:"required"`
	Occupation 	string `json:"occupation" binding:"required"`
	Email 		string `json:"email" binding:"required,email"`
	Token 		string `json:"token" binding:"required"`
}

func FormatUser(user User, token string) UserFormatter {
 formatter := UserFormatter{
	 ID 	: user.ID,
	 Name	: user.Name,
	 NoKta	: user.NoKta,
	 Occupation: user.Occupation,
	 Email: user.Email,
	 Token : token,
 }
 return formatter
}
package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 					int
	Name		 		string
	Occupation  		string
	NoKta 				string
	Email  				string
	PasswordHash		string
	Avatar_file_name 	string
	Role 				string



}

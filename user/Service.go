package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser (input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)

}
// maaping struck input ke struct USer
// sampan struch melalui reposuitory
type service struct {
		repository Repository
}
func NewService(repository Repository) *service {
	return &service{repository}
}


func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.NoKta = input.NoKta
	user.Occupation =  input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, nil
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindById(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0  {
		return user, errors.New("user tidaks adah")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(password))
	if err != nil {
		return user, err
	}
	return user, err
}





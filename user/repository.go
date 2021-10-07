package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User,error)
	FindById(email string) (User, error)
	FindByEmail(Email string) (User, error)
}
//tangkap untuk mengabil datadb
type repository struct {
	db *gorm.DB
}

func (r *repository) FindById(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user , nil
}


func NewRepository(db *gorm.DB) *repository {
	return &repository{db:db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return User{}, err
	 }
	 return user,nil
}

//func (r *repository) FindById(email string) (User, error) {
//	panic("implement me")
//}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user , nil
}







package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User,error)
}
//tangkap untuk mengabil datadb
type repository struct {
	db *gorm.DB
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
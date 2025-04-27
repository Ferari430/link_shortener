package user

import (
	"log"
	"my_project/db"
)

type UserRepository struct {
	Database db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{Database: *database}
}

func (repo UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {

		return nil, result.Error
	}

	return user, nil
}

func (repo UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.First(&user, "email = ?", email)
	if result.Error != nil {
		log.Println("Not found [FINDING BY EMAIL]")
		return nil, result.Error
	}

	return &user, nil
}

package auth

import "gorm.io/gorm"

var db *gorm.DB

func UserAutomigrate(gormDB *gorm.DB) {
	db = gormDB
	db.AutoMigrate(&user{})
}

func dbCreateUser(req CreateUserRequest) (CreateUserResponse, error) {
	user := user{
		email:    req.Email,
		password: req.Password,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return CreateUserResponse{}, result.Error
	}
	return CreateUserResponse{
		ID:    user.Model.ID,
		Email: user.email,
	}, nil
}

func dbGetUser(id uint) (user, error) {
	var user user
	result := db.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

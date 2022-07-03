package auth

func CreateUser(req CreateUserRequest) (CreateUserResponse, error) {
	return dbCreateUser(req)
}

func GetUser(id uint) (user, error) {
	return dbGetUser(id)
}

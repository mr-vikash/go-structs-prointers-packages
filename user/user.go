package user

type User struct {
	Name     string
	Email    string
	Password string
	Age      int
}

func NewUser(name string, email string, password string, age int) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
	}
}

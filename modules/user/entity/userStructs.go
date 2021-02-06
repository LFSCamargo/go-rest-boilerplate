package userStructs

// User - Is the return for the user api
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// AuthenticationOutput - is the return for the login and signup with a new token
type AuthenticationOutput struct {
	Token string `json:"token"`
}

// RegisterInput - Is the struct for converting the []byte to a readable struct
type RegisterInput struct {
	Email    string
	Name     string
	Password string
}

// LoginInput - Is the struct for converting the []byte to a readable struct
type LoginInput struct {
	Email    string
	Password string
}

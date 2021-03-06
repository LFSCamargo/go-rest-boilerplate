package userModel

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// User - is the mgm model for the user inside mongo
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
}

// CreateNewUser - creates a new user inside the database using the mgm
func CreateNewUser(name string, email string, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(hash)
	dbUser := &User{
		Username: name,
		Email:    email,
		Password: hashedPassword,
	}
	error := mgm.Coll(dbUser).Create(dbUser)

	return dbUser, error
}

// FindByEmail - Finds a user by email
func FindByEmail(email string) (*User, error) {
	user := &User{Email: email}
	coll := mgm.Coll(user)
	err := coll.First(bson.M{"email": email}, user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

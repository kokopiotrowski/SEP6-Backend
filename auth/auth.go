package auth

import (
	"database/sql"
	"errors"
	"os"
	"regexp"
	swagger "studies/SEP6-Backend/swagger/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RegisterUser(db *sql.DB, register swagger.Register) (bool, error) {

	var validUsername = regexp.MustCompile(`^(?=.{5,20}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$`)
	var validPassword = regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,}$`)

	if !validUsername.MatchString(register.Username) {
		return false, errors.New("Invalid username. Should contain 5-20 characters")
	}

	if !validPassword.MatchString(register.Password) {
		return false, errors.New("Invalid password. Should contain at least 6 characters. At least one number and one letter")
	}

	hashedPass, err := hashPassword(register.Password)

	query := "INSERT INTO Users(Username, Pass) VALUES (?, ?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(register.Username, hashedPass)

	if err != nil {
		return false, err
	}
	return true, nil
}

func LogIn(db *sql.DB, login swagger.Login) (string, error) {

	var userId uint64
	var username string
	var hashedPassDB string

	err := db.QueryRow("SELECT UserId, Username, Pass FROM Users WHERE Username=?", login.Username).Scan(&userId, &username, &hashedPassDB)

	if err != nil {
		return "", errors.New("Could not login. Didn't find user with this username")
	}

	hp, err := hashPassword(login.Password)
	if hp != hashedPassDB {
		return "", errors.New("Could not login. Incorrect password")
	}
	token, err := CreateToken(userId, username)
	if err != nil {
		return "", errors.New("Could not generate token.")
	}
	return token, nil
}

func CreateToken(userId uint64, username string) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId

	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

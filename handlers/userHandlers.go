package handlers

import (
	"errors"
	"sagara-try/config"
	"sagara-try/helpers"
	"sagara-try/middleware"
	"sagara-try/models"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type H map[string]interface{}
type User models.User

func (h *User) H_Login() (H, error) {

	datum := User{}

	err := config.GetDB().Debug().Where("email = ?", h.Email).Take(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.Login] - email is not exist "+err.Error())
		return nil, errors.New("email not exists")
	}

	err = helpers.VerifyPassword(datum.Password, h.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		helpers.Logger("error", "In Server: invalid verify password")
		return nil, errors.New("Password is failure")
	}

	strUserID := strconv.Itoa(datum.UserID)
	token, err := middleware.CreateToken(strUserID, datum.BlogID, datum.Username, datum.Email, datum.Role)
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.Login] - failure create token "+err.Error())
		return nil, err
	}

	rMsg := H{}
	rMsg["user_id"] = datum.UserID
	rMsg["blog_id"] = datum.BlogID
	rMsg["username"] = datum.Username
	rMsg["email"] = datum.Email
	rMsg["role"] = datum.Role
	access := token["accessToken"]
	refresh := token["refreshToken"]

	return H{"accessToken": access, "refreshToken": refresh, "users": rMsg}, nil
}

func (h *User) H_Register() (H, error) {
	hashedPassword, err := helpers.Hash(h.Password)
	if err != nil {
		helpers.Logger("error", "In Server: [userHandler.Register] - hashed: "+err.Error())
		return nil, err
	}

	datum := User{}
	datum.UserID = h.UserID
	datum.Username = h.Username
	datum.Email = h.Email
	datum.Role = "user"
	datum.Password = string(hashedPassword)
	datum.Created = time.Now()

	err = config.GetDB().Debug().Create(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [userHandler.Register] - failed insert: "+err.Error())
		return nil, err
	}

	strUserID := strconv.Itoa(datum.UserID)
	token, err := middleware.CreateToken(strUserID, datum.BlogID, datum.Username, datum.Email, datum.Role)
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.Register] - failure create token "+err.Error())
		return nil, err
	}

	msg := H{}
	msg["id"] = datum.UserID
	msg["username"] = datum.Username
	msg["email"] = datum.Email
	msg["role"] = datum.Role
	msg["created"] = datum.Created
	access := token["accessToken"]
	refresh := token["refreshToken"]

	return H{"accessToken": access, "refreshToken": refresh, "users": msg}, nil
}

func H_GetOneUser(userId string) (*User, error) {
	var datum User
	err := config.GetDB().Debug().Where("user_id = ?", userId).Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.GetOneUser] - id is not exist "+err.Error())
		return nil, err
	}
	return &datum, nil
}

func (h *User) H_UpdateOneUser(userId int) (H, error) {
	hashedPassword, er := helpers.Hash(h.Password)
	if er != nil {
		helpers.Logger("error", "In Server: [UserHandler.Update] - hashed: "+er.Error())
		return nil, er
	}

	h.Password = string(hashedPassword)

	datum := User{}
	datum.UserID = userId
	datum.Username = h.Username
	datum.Email = h.Email
	datum.Role = "user"
	datum.Password = h.Password
	datum.Updated = time.Now()

	err := config.GetDB().Debug().Model(datum).Where("user_id = ?", userId).Update(&h).Error
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.Update] - failed updated data, "+err.Error())
		return nil, err
	}

	msg := H{}
	msg["id"] = datum.UserID
	msg["username"] = datum.Username
	msg["email"] = datum.Email
	msg["role"] = datum.Role
	msg["updated"] = datum.Updated

	return msg, err
}

func H_DeleteOneUser(userId string) (string, error) {
	rowsAffected := config.GetDB().Debug().Model(User{}).Where("blog_id = ?", userId).Delete(User{}).RowsAffected
	if rowsAffected == 0 {
		helpers.Logger("error", "In Server: [BlogHandler.Deleted] - id is not exist")
		return "", errors.New("id is not exist")
	}
	return "success to deleted", nil
}

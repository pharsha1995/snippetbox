package mocks

import (
	"time"

	"github.com/pharsha1995/snippetbox/internal/models"
)

var mockUser = models.User{
	Name:  "alice",
	Email: "alice@example.com",
	Created: time.Now(),
}

var mockPassword = "password"

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == mockPassword {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return &mockUser, nil
	default:
		return &models.User{}, models.ErrNoRecord
	}
}

func (m *UserModel) PasswordUpdate(id int, curPassword, newPassword string) error {
	switch id {
	case 1:
		if curPassword == mockPassword {
			mockPassword = newPassword
			return nil
		}
		return models.ErrInvalidCredentials
	default:
		return models.ErrNoRecord
	}
}
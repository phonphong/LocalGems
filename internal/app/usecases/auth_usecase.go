package usecases

import (
	"locall-gems-server/internal/core/entity"
	"locall-gems-server/internal/core/errors"
	"locall-gems-server/internal/infra/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v4"
)

type AuthUsecase interface {
	Register(user *entity.User) error
	Login(credentials *entity.Credentials) (string, error)
}

type authUsecase struct {
	userRepo repositories.UserRepository
	secret   string
}

func NewAuthUsecase(repo repositories.UserRepository, secret string) AuthUsecase {
	return &authUsecase{userRepo: repo, secre
		t: secret}
}

func (u *authUsecase) Register(user *entity.User) error {
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.userRepo.Create(user)
}

func (u *authUsecase) Login(credentials *entity.Credentials) (string, error) {
	user, err := u.userRepo.FindByEmail(credentials.Email)
	if err != nil {
		return "", errors.NewAuthenticationError("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.NewAuthenticationError("invalid email or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(u.secret))
}

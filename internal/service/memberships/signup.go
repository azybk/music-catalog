package memberships

import (
	"errors"
	"log"

	"github.com/azybk/music-catalog/internal/models/memberships"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(request memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatalf("error get user from database, err:%+v", err)
		return err
	}

	if existingUser != nil {
		log.Fatal("username or email already exists")
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error hash password, err: %+v", err)
		return err
	}

	model := memberships.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(pass),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}

	return s.repository.CreateUser(model)
}

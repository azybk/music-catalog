package memberships

import (
	"errors"
	"log"

	"github.com/azybk/music-catalog/internal/models/memberships"
	"github.com/azybk/music-catalog/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(req memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(req.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("error get user")
		return "", err
	}

	if userDetail == nil {
		log.Println("email not exist")
		return "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password))
	if err != nil {
		log.Println("email and password not match")
		return "", errors.New("email and password not match")
	}

	accessToken, err := jwt.CreateToken(int64(userDetail.ID), userDetail.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Println("failed create JWT token")
		return "", err
	}

	return accessToken, nil
}

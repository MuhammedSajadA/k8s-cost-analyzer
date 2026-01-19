package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/models"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret string
}

func NewAuthService(db *gorm.DB, secret string) *AuthService {
	return &AuthService{db: db, jwtSecret: secret}
}

func (s *AuthService) Register(email, password string) error {
	var existing models.User
	if err := s.db.Where("email = ?", email).First(&existing).Error; err == nil {
		return errors.New("user already exists")
	} else if err != gorm.ErrRecordNotFound {
		return errors.New("failed to check user existence")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		Email:    email,
		Password: string(hash),
	}

	return s.db.Create(&user).Error
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("invalid credentials")
		}
		return "", errors.New("failed to query user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", errors.New("failed to sign token")
	}
	return signed, nil
}

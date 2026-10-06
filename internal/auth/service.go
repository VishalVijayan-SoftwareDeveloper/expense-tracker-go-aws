package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/config"
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      Repository
	jwtSecret string
}

func NewService(
	repo Repository,
	cfg *config.Config,
) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: cfg.JWTSecret,
	}
}

func (s *Service) Register(
	ctx context.Context,
	req RegisterRequest,
) error {

	req.Email = strings.TrimSpace(req.Email)

	if req.Email == "" {
		return errors.New("email is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	existingUser, err := s.repo.GetByEmail(
		ctx,
		req.Email,
	)

	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := &user.User{
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	return s.repo.CreateUser(
		ctx,
		user,
	)
}

func (s *Service) Login(
	ctx context.Context,
	req LoginRequest,
) (string, error) {

	user, err := s.repo.GetByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := GenerateToken(
		user.ID,
		s.jwtSecret,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

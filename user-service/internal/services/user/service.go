package user

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/constants"
	errConstant "user-service/internal/constants/error"
	"user-service/internal/domain/model"
	"user-service/internal/repository"
	"user-service/internal/repository/role"
	"user-service/internal/repository/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository     user.Repository
	roleRepository     role.Repository
	transactionManager repository.TransactionManager
}

type Service interface {
	Register(context.Context, CreateInput) error
}

func NewService(userRepository user.Repository, roleRepository role.Repository, transactionManager repository.TransactionManager) Service {
	return &service{
		userRepository:     userRepository,
		roleRepository:     roleRepository,
		transactionManager: transactionManager,
	}
}

func (s *service) Register(ctx context.Context, input CreateInput) error {
	exist, err := s.userRepository.ExistByEmail(ctx, input.Email)
	if err != nil {
		return err
	}
	if exist {
		return errConstant.ErrAlreadyExists
	}

	customerRole, err := s.roleRepository.FindByCode(ctx, constants.Customer)
	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("user register : %w", err)
	}

	rawToken, err := generateToken()
	if err != nil {
		return err
	}

	// for testing
	fmt.Println("token :", rawToken)

	hashToken := hashToken(rawToken)
	expiresAt := time.Now().UTC().Add(emailVerificationTokenTTL)

	err = s.transactionManager.WithinTransaction(ctx, func(repositories *repository.Registry) error {
		newUser := &model.User{
			UUID:         uuid.New(),
			Name:         input.Name,
			Username:     input.Username,
			PasswordHash: string(passwordHash),
			Email:        input.Email,
			RoleID:       customerRole.ID,
		}

		if err := repositories.User.Create(ctx, newUser); err != nil {
			return err
		}

		if err := repositories.VerificationToken.Create(ctx, &model.VerificationToken{
			UserID:    newUser.ID,
			HashToken: hashToken,
			TokenType: constants.TokenTypeEmailVerification,
			ExpiresAt: expiresAt,
		}); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

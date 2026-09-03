package user

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"user-service/internal/constants"
	errConstant "user-service/internal/constants/error"
	"user-service/internal/domain/model"
	"user-service/internal/repository"
	"user-service/internal/repository/role"
	"user-service/internal/repository/user"
	"user-service/internal/repository/verificationtoken"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository              user.Repository
	roleRepository              role.Repository
	verificationTokenRepository verificationtoken.Repository
	transactionManager          repository.TransactionManager
}

type Service interface {
	Register(context.Context, CreateInput) error
	VerifyAccount(context.Context, string) error
	Authenticate(context.Context, AuthenticateInput) (*AuthenticatedUser, error)
}

func NewService(userRepository user.Repository, roleRepository role.Repository, verificationTokenRepository verificationtoken.Repository, transactionManager repository.TransactionManager) Service {
	return &service{
		userRepository:              userRepository,
		roleRepository:              roleRepository,
		verificationTokenRepository: verificationTokenRepository,
		transactionManager:          transactionManager,
	}
}

func (s *service) Register(ctx context.Context, input CreateInput) error {
	email := normalizeEmail(input.Email)

	exist, err := s.userRepository.ExistByEmail(ctx, email)
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
			Email:        email,
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

func (s *service) VerifyAccount(ctx context.Context, token string) error {
	hashToken := hashToken(token)

	verifyToken, err := s.verificationTokenRepository.FindByHashToken(ctx, hashToken)
	if err != nil {
		return err
	}

	if verifyToken.TokenType != constants.TokenTypeEmailVerification {
		slog.Error("invalid token type", "error", err, "hashToken", hashToken)
		return errConstant.ErrInvalidToken
	}

	now := time.Now().UTC()

	if !now.Before(verifyToken.ExpiresAt) {
		slog.Error("invalid token expired", "error", err, "hashToken", hashToken)
		return errConstant.ErrInvalidToken
	}

	err = s.transactionManager.WithinTransaction(ctx, func(repositories *repository.Registry) error {
		if err := repositories.User.MarkVerified(ctx, verifyToken.UserID); err != nil {
			return err
		}

		if err := repositories.VerificationToken.Delete(ctx, hashToken); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) Authenticate(ctx context.Context, input AuthenticateInput) (*AuthenticatedUser, error) {
	email := normalizeEmail(input.Email)
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return nil, errConstant.ErrPasswordNotMatch
	}

	if !user.IsVerified {
		return nil, errConstant.ErrAccountNotVerified
	}

	return &AuthenticatedUser{
		UUID:     user.UUID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

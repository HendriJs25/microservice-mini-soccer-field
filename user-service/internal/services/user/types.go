package user

import "github.com/google/uuid"

type CreateInput struct {
	Name     string
	Username string
	Password string
	Email    string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResult struct {
	Token    string
	RoleName string
	User     AuthenticatedUser
}

type AuthenticateInput struct {
	Email    string
	Password string
}

type AuthenticatedUser struct {
	UUID     uuid.UUID
	Name     string
	Username string
	Email    string
}

package service

import (
	"context"
	"go-backend-template/internal/model"
	"go-backend-template/internal/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *model.UserCreate) (*model.User, error) {
	id, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, req *model.UserUpdate) (*model.User, error) {
	if err := s.repo.Update(ctx, id, req); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

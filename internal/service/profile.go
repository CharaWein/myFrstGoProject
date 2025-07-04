package service

import (
	"context"
	"go-backend-template/internal/model"
	"go-backend-template/internal/repository"

	"github.com/google/uuid"
)

type ProfileService struct {
	repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(ctx context.Context, req *model.ProfileCreate) (*model.Profile, error) {
	id, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

func (s *ProfileService) GetProfile(ctx context.Context, id uuid.UUID) (*model.Profile, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProfileService) GetAllProfiles(ctx context.Context) ([]model.Profile, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProfileService) UpdateProfile(ctx context.Context, id uuid.UUID, req *model.ProfileUpdate) (*model.Profile, error) {
	if err := s.repo.Update(ctx, id, req); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

func (s *ProfileService) DeleteProfile(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProfileService) GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*model.Profile, error) {
	return s.repo.GetByUserID(ctx, userID)
}

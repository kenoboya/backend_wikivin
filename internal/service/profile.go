package service

import (
	"context"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
)

type ProfilesService struct {
	profileRepo  repo.Profiles
}

func NewProfileService(profileRepo repo.Profiles) *ProfilesService {
	return &ProfilesService{
		profileRepo:  profileRepo,
	}
}

func (s *ProfilesService) GetBriefInfoProfile(ctx context.Context, userID int) (model.BriefInfoProfile, error){
	return s.profileRepo.GetBriefInfoProfile(ctx, userID)
}

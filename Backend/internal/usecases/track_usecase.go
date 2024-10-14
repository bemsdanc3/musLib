package usecases

import (
	"Backend/internal/entities"
	"Backend/internal/repository"
)

type TrackUsecase interface {
	GetTracksByUserID(userID int) ([]*entities.Track, error)
}

type trackUsecase struct {
	repo repository.TrackRepository
}

func NewTrackUseccase(repo repository.TrackRepository) TrackUsecase {
	return &trackUsecase{repo: repo}
}

func (t *trackUsecase) GetTracksByUserID(userID int) ([]*entities.Track, error) {
	return t.repo.GetTracksByUserID(userID)
}

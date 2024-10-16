package usecases

import (
	"Backend/internal/entities"
	"Backend/internal/repository"
)

type TrackUsecase interface {
	GetTracksByUserID(userID int) ([]*entities.Track, error)
	GetAllTracks() ([]*entities.Track, error)
	CreateTrack(track *entities.Track) error
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

func (t *trackUsecase) GetAllTracks() ([]*entities.Track, error) {
	return t.repo.GetAllTracks()
}

func (t *trackUsecase) CreateTrack(track *entities.Track) error {
	return t.repo.CreateTrack(track)
}

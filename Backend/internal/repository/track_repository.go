package repository

import (
	"Backend/internal/entities"
	"database/sql"
	"errors"
)

type TrackRepository interface {
	GetTracksByUserID(userID int) ([]*entities.Track, error)
	GetAllTracks() ([]*entities.Track, error)
}

type trackRepository struct {
	db *sql.DB
}

func NewTrackRepository(db *sql.DB) TrackRepository {
	return &trackRepository{db: db}
}

func (r *trackRepository) GetTracksByUserID(userID int) ([]*entities.Track, error) {
	query := `SELECT id, title, description, uploaddate, filelink FROM tracks WHERE userid = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*entities.Track
	for rows.Next() {
		track := &entities.Track{}
		err := rows.Scan(&track.ID, &track.Title, &track.Description, &track.UploadDate, &track.FileLink)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	if len(tracks) == 0 {
		return nil, errors.New("no tracks found for this user")
	}
	return tracks, nil
}

func (r *trackRepository) GetAllTracks() ([]*entities.Track, error) {
	query := `SELECT id, title, description, userid, uploaddate, filelink FROM tracks`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*entities.Track
	for rows.Next() {
		track := &entities.Track{}
		err := rows.Scan(&track.ID, &track.Title, &track.Description, &track.UserID, &track.UploadDate, &track.FileLink)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

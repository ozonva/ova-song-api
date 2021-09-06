package repo

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-song-api/internal/models"
)

type Repo interface {
	AddSong(song models.Song) (int64, error)
	AddSongs(songs []models.Song) (int64, error)
	ListSongs(limit, offset uint64) ([]models.Song, error)
	DescribeSong(songId uint64) (*models.Song, error)
	UpdateSong(song models.Song) (bool, error)
	RemoveSong(songId uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	tableName string
}

func (r *repo) AddSong(songs models.Song) (int64, error) {
	query := squirrel.Insert(r.tableName).
		Columns("name", "author", "year").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	query = query.Values(songs.Name, songs.Author, songs.Year)

	var insertedID int64
	err := query.QueryRow().Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (r *repo) AddSongs(songs []models.Song) (int64, error) {
	query := squirrel.Insert(r.tableName).
		Columns("name", "author", "year").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for i := range songs {
		query = query.Values(songs[i].Name, songs[i].Author, songs[i].Year)
	}

	rows, err := query.Query()
	if err != nil {
		return 0, err
	}

	var id int64
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) ListSongs(limit, offset uint64) ([]models.Song, error) {
	query := squirrel.Select("id", "name", "author", "year").
		From(r.tableName).
		Limit(limit).
		Offset(offset).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.Id, &song.Name, &song.Author, &song.Year)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *repo) DescribeSong(songId uint64) (*models.Song, error) {
	query := squirrel.Select("id", "name", "author", "year").
		From(r.tableName).
		Where(squirrel.Eq{"id": songId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	row := query.QueryRow()

	var song models.Song
	if err := row.Scan(&song.Id, &song.Name, &song.Author, &song.Year); err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *repo) UpdateSong(song models.Song) (bool, error) {
	query := squirrel.Update(r.tableName).
		Set("name", song.Name).
		Set("author", song.Author).
		Set("year", song.Year).
		Where(squirrel.Eq{"id": song.Id}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	result, err := query.Exec()
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func (r *repo) RemoveSong(songId uint64) (bool, error) {
	query := squirrel.Delete(r.tableName).
		Where(squirrel.Eq{"id": songId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	result, err := query.Exec()

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected != 0, nil
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		db: db, tableName: "songs",
	}
}

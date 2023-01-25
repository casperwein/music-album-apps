package psql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/casperwein/go-edspert/album-app/entity"
)

// Create is function to create artist to database
func (repo *artistConnection) Create(ctx context.Context, artist *entity.Artist) (int64, error) {
	// The query insert
	query := `
        INSERT INTO public.artists (name) 
        VALUES ($1)
        RETURNING id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, artist.Name).Scan(&artist.Id)
	if err != nil {
		return 0, err
	}

	return artist.Id, nil
}

// Get is function to get specific artist by id from database
func (repo *artistConnection) Get(ctx context.Context, id int64) (*entity.Artist, error) {
	// The query select
	query := `
        SELECT id, name
        FROM artists
        WHERE id = $1`

	var artist entity.Artist

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query and find the specific Artist and then set the result to Artist variable
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&artist.Id,
		&artist.Name,
	)

	// If any error
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

// GetAllArtist is function to get all Artists from database
func (repo *artistConnection) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	// The query select
	query := `
		SELECT id, name
		FROM artists`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var artists []entity.Artist

	// Run the query to get all artists
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the artists is not empty
	for rows.Next() {
		var artist entity.Artist

		// Set to the artist variable
		err := rows.Scan(
			&artist.Id,
			&artist.Name,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the artist to artist variable
		artists = append(artists, artist)
	}

	return artists, nil
}

// BatchCreate is function to insert some albums in once to database
func (repo *artistConnection) BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error) {
	var IDs []int64

	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return IDs, nil
	}
	// If any error, the transaction will be rollback
	defer tx.Rollback()

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query insert
	query := `INSERT INTO artists (name) VALUES ($1) RETURNING id`

	// Loop every album
	for _, artist := range artists {
		var id int64

		// Run query insert of every Artist to database
		err := tx.QueryRowContext(ctx, query, artist.Name).Scan(&id)
		if err != nil {
			log.Printf("error execute insert err: %v", err)
			continue
		}

		// Add the new id to IDs variable
		IDs = append(IDs, id)
	}

	// Commit the transaction
	err = tx.Commit()
	// If any error
	if err != nil {
		return IDs, err
	}

	return IDs, nil
}

// Update is function to update album in database
func (repo *artistConnection) Update(ctx context.Context, artist entity.Artist) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE artists set name=$1 WHERE id=$2`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, artist.Name, artist.Id)
	if err != nil {
		return err
	}

	// Get how many data has been updated
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)
	return nil
}

// Delete is function to delete album in database
func (repo *artistConnection) Delete(ctx context.Context, id int64) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from artists WHERE id=$1`

	// Run the delete query
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Get how many data has been deleted
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected delete : %d", rows)
	return nil
}

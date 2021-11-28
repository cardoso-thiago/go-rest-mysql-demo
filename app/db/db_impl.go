package db

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/cardoso-thiago/go-rest-mysql-demo/app/model"

	_ "github.com/go-sql-driver/mysql"
)

func (d *DB) Open() error {
	db, err := sqlx.Open("mysql", dbConnectionStr)
	if err != nil {
		return err
	}
	log.Println("Conectado na base de dados.")

	db.MustExec(dropSchema)
	db.MustExec(createSchema)

	d.db = db
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) CreateArtist(p *model.Artist) (model.Artist, error) {
	var artist model.Artist
	var insertedId int
	err := d.db.QueryRow(insertArtistQuery, p.Name, p.Genre, p.Country).Scan(&insertedId)
	if err != nil {
		return artist, err
	}

	d.db.QueryRow(getArtistQuery, insertedId).Scan(&artist.ID, &artist.Name, &artist.Genre, &artist.Country)
	return artist, nil
}

func (d *DB) GetArtists() ([]*model.Artist, error) {
	var artists []*model.Artist
	err := d.db.Select(&artists, getAllArtistsQuery)
	if err != nil {
		return artists, err
	}

	return artists, nil
}

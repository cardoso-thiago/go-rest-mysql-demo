package db

import (
	"github.com/jmoiron/sqlx"

	"github.com/cardoso-thiago/go-rest-mysql-demo/app/model"

	_ "github.com/go-sql-driver/mysql"
)

type ArtistDB interface {
	Open() error
	Close() error
	CreateArtist(p *model.Artist) (model.Artist, error)
	GetArtists() ([]*model.Artist, error)
}

type DB struct {
	db *sqlx.DB
}

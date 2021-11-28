package db

import "fmt"

var (
	user            = "cardoso"
	pass            = "112358"
	prot            = "tcp"
	addr            = "localhost:3306"
	dbname          = "music_catalog"
	netAddr         = fmt.Sprintf("%s(%s)", prot, addr)
	dbConnectionStr = fmt.Sprintf("%s:%s@%s/%s?timeout=30s", user, pass, netAddr, dbname)
)

const createSchema = `
CREATE TABLE IF NOT EXISTS artist
(
	id SERIAL PRIMARY KEY,
    name varchar(50) NOT NULL,
    genre varchar(50) NOT NULL,
    country varchar(50) NOT NULL
)
`

const dropSchema = "DROP TABLE IF EXISTS artist"

var insertArtistQuery = "INSERT INTO artist (name, genre, country) VALUES (?, ?, ?) RETURNING id"

const getArtistQuery = "SELECT * FROM artist WHERE ID = ?"

const getAllArtistsQuery = "SELECT * FROM artist"

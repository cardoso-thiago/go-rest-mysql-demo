package app

import (
	"log"
	"net/http"

	"github.com/cardoso-thiago/go-rest-mysql-demo/app/model"
)

func (app *App) CreateArtistHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.Artist{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Erro ao formatar o json do corpo. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		artist := &model.Artist{
			ID:      0,
			Name:    req.Name,
			Genre:   req.Genre,
			Country: req.Country,
		}

		insertedArtist, err := app.DB.CreateArtist(artist)
		if err != nil {
			log.Printf("Não foi possível salvar o artista na base. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		sendResponse(w, r, insertedArtist, http.StatusOK)

	}
}

func (app *App) GetArtistsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		artists, err := app.DB.GetArtists()
		if err != nil {
			log.Printf("Não foi possível obter os artistas., err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]model.ArtistGet, len(artists))
		for idx, artist := range artists {
			resp[idx] = mapArtistToJSON(artist)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}

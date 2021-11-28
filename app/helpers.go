package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cardoso-thiago/go-rest-mysql-demo/app/model"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Erro ao formatar json. err=%v\n", err)
	}
}

func mapArtistToJSON(p *model.Artist) model.ArtistGet {
	return model.ArtistGet{
		Name:    p.Name,
		Genre:   p.Genre,
		Country: p.Country,
	}
}

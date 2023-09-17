package queries

import (
	"github.com/misterek/steampipe-plugin-starwars/starwars/models"
)

type StarshipQuery struct {
	AllStarships struct {
		Starships []models.Starship
	}
}

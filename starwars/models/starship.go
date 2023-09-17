package models

import (
	"github.com/shurcooL/graphql"
)

type Starship struct {
	CargoCapacity        graphql.Float
	Consumables          graphql.String
	CostInCredits        graphql.Float
	Crew                 graphql.String
	HyperdriveRating     graphql.Float
	Id                   graphql.String
	Length               graphql.Float
	Manufacturers        []graphql.String
	MaxAtmospheringSpeed graphql.Float
	Model                graphql.String
	Name                 graphql.String
	Passengers           graphql.String
	StarshipClass        graphql.String
}

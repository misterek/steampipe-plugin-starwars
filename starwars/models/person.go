package models

import (
	"github.com/shurcooL/graphql"
)

type Person struct {
	Height    graphql.Int
	Name      graphql.String
	Id        graphql.String
	BirthYear graphql.String
	EyeColor  graphql.String
	Gender    graphql.String
	HairColor graphql.String
	SkinColor graphql.String
}

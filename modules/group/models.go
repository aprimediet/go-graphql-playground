package group

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// Group is group data type
type Group struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}

// GroupType is group type schema for graphql
var GroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Group",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				source := params.Source.(Group)

				return source.ID.Hex(), nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

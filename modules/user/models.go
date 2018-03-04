package user

import (
	"github.com/aprimediet/gographql/db"
	"github.com/aprimediet/gographql/modules/group"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// User is user data type
type User struct {
	ID     bson.ObjectId   `bson:"_id" json:"id"`
	Name   string          `bson:"name" json:"name"`
	Email  string          `bson:"email" json:"email"`
	Groups []bson.ObjectId `bson:"groups" json:"groups"`
}

// UserType is users type schema for graphql
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				source := params.Source.(User)

				return source.ID.Hex(), nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(group.GroupType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				source := params.Source.(User)
				groups := []group.Group{}

				err := db.DB.C("groups").Find(bson.M{"_id": bson.M{"$in": source.Groups}}).All(&groups)

				return groups, err
			},
		},
	},
})

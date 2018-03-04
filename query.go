package main

import (
	"github.com/aprimediet/gographql/modules/group"
	"github.com/aprimediet/gographql/modules/user"
	"github.com/graphql-go/graphql"
)

// RootQuery is Root Query
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users":  user.List,
		"user":   user.Retrieve,
		"groups": group.List,
		"group":  group.Retrieve,
	},
})

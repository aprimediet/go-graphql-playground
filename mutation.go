package main

import (
	"github.com/aprimediet/gographql/modules/group"
	"github.com/aprimediet/gographql/modules/user"
	"github.com/graphql-go/graphql"
)

// RootMutation is graphql root mutation
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"userCreate":  user.Create,
		"userUpdate":  user.Update,
		"userDelete":  user.Delete,
		"groupCreate": group.Create,
		"groupUpdate": group.Update,
		"groupDelete": group.Delete,
	},
})

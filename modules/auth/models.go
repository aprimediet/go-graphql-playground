package auth

import "gopkg.in/mgo.v2/bson"

// Token is Token model for mongodb
type Token struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	UserID   bson.ObjectId `bson:"userId" json:"userId"`
	Created  string        `bson:"created" json:"created"`
	Expired  string        `bson:"expired" json:"expired"`
	IsActive bool          `bson:"isActive" json:"isActive"`
}

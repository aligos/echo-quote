package quotemodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Quote struct {
	Id           bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	QuoteMessage string        `json:"quoteMessage,omitempty" bson:"quoteMessage"`
	Author       string        `json:"author,omitempty" bson:"author"`
	CreatedAt    time.Time     `json:"createdAt,omitempty" bson:"createdAt"`
}

func (q Quote) IsValid() bool {
	if l := len(q.QuoteMessage); l > 4 {
		return true
	}

	return false
}

type Quotes []Quote

package quoteactions

import (
	"errors"
	"math/rand"
	"time"

	quote "github.com/aligos/echo-quote/api/quote/model"
	"github.com/aligos/echo-quote/config"
	"gopkg.in/mgo.v2/bson"
)

const col string = "quotes"

func GetRandomQuote() (quote.Quote, error) {
	var q quote.Quote
	db := dbconfig.DB{}
	qs := quote.Quotes{}

	s, err := db.DoDial()

	if err != nil {
		return q, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{}).All(&qs)
	rand.Seed(time.Now().UnixNano())
	for i := len(qs) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		q = qs[j]
	}

	if err != nil {
		return q, errors.New("There was an error trying to find the quotes.")
	}

	return q, err
}

func All() (quote.Quotes, error) {
	db := dbconfig.DB{}
	qs := quote.Quotes{}

	s, err := db.DoDial()

	if err != nil {
		return qs, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{}).All(&qs)

	if err != nil {
		return qs, errors.New("There was an error trying to find the quotes.")
	}

	return qs, err
}

func GetById(id string) (quote.Quote, error) {
	db := dbconfig.DB{}
	q := quote.Quote{}

	s, err := db.DoDial()

	if err != nil {
		return q, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&q)

	if err != nil {
		return q, errors.New("There was an error trying to find the quotes.")
	}

	return q, err
}

func NewQuote(q quote.Quote) (quote.Quote, error) {
	db := dbconfig.DB{}
	q.Id = bson.NewObjectId()
	q.CreatedAt = time.Now()

	s, err := db.DoDial()

	if err != nil {
		return q, errors.New("There was an error trying to connect to the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Insert(q)

	if err != nil {
		return q, errors.New("There was an error trying to insert the doc to the DB.")
	}

	return q, err
}

func DeleteQuote(id string) error {
	db := dbconfig.DB{}

	s, err := db.DoDial()

	if err != nil {
		return errors.New("There was an error trying to connect with the DB.")
	}

	idoi := bson.ObjectIdHex(id)

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.RemoveId(idoi)

	if err != nil {
		return errors.New("There was an error trying to remove the quote.")
	}

	return err
}

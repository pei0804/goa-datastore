package model

import "cloud.google.com/go/datastore"

type Post struct {
	_kind   string         `goon:"kind,Post"`
	ID      string         `datastore:"-" goon:"id"`
	Content string         `datastore:"content"`
	User    *datastore.Key `goon:"User" datastore:"-"`
}

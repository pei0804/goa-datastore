package model

type Account struct {
	ID   int64 `datastore:"-" goon:"id"`
	Name string
}

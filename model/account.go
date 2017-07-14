package model

import (
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Account struct {
	ID   int    `datastore:"-" goon:"id"`
	Name string `datastore:"name"`
}

type AccountDB struct {
}

func (db *AccountDB) GetByName(ctx context.Context, name string) (*[]Account, error) {
	g := goon.FromContext(ctx)

	as := &[]Account{}
	_, err := g.GetAll(datastore.NewQuery("Account").Filter("name =", name), &as)
	if err != nil {
		log.Debugf(ctx, "%v", err)
		return nil, err
	}
	return as, nil
}

func (db *AccountDB) Get(ctx context.Context, ID int) (*Account, error) {
	g := goon.FromContext(ctx)
	a := &Account{ID: ID}
	if err := g.Get(a); err != nil {
		log.Debugf(ctx, "%v", err)
		return nil, err
	}
	return a, nil
}

func (db *AccountDB) Add(ctx context.Context, account *Account) error {
	g := goon.FromContext(ctx)
	_, err := g.Put(&account)
	if err != nil {
		return err
	}
	return nil
}

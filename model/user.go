package model

import (
	"fmt"

	"github.com/mjibson/goon"
	"github.com/pei0804/goa-datastore/app"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type User struct {
	_kind string `goon:"kind,User"`
	ID    int64  `datastore:"-" goon:"id"`
	IDStr string
	Name  string `datastore:"name"`
}

type UserDB struct {
}

func (db *UserDB) GetFindByName(ctx context.Context, name string) ([]*app.User, error) {
	g := goon.FromContext(ctx)
	as := []*User{}
	_, err := g.GetAll(datastore.NewQuery(g.Kind(new(User))).Filter("name =", name), &as)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return nil, err
	}
	var appAs []*app.User
	for _, t := range as {
		appAs = append(appAs, t.UserToUser())
	}

	return appAs, nil
}

func (db *UserDB) Get(ctx context.Context, id int64) (*User, error) {
	g := goon.FromContext(ctx)
	u := User{
		ID: id,
	}
	err := g.Get(&u)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return nil, err
	}
	return &u, nil
}

func (db *UserDB) Add(ctx context.Context, user *User) (*User, error) {
	g := goon.FromContext(ctx)
	_, err := g.Put(user)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return nil, err
	}
	err = g.Get(user)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return nil, err
	}
	return user, nil
}

func (db *UserDB) Delete(ctx context.Context, id int64) error {
	g := goon.FromContext(ctx)
	u := &User{
		ID: id,
	}
	err := g.Get(u)
	if err != nil {
		return err
	}
	err = g.Delete(g.Key(u))
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return err
	}
	return nil
}

func (db *UserDB) Update(ctx context.Context, id int64, updateUser *User) (*User, error) {
	g := goon.FromContext(ctx)
	findUser := &User{
		ID: id,
	}
	err := g.RunInTransaction(func(g *goon.Goon) error {
		err := g.Get(findUser)
		if err != nil {
			log.Errorf(ctx, "%v", err)
			return err
		}
		updateUser.ID = findUser.ID
		_, err = g.Put(updateUser)
		if err != nil {
			log.Errorf(ctx, "%v", err)
			return err
		}
		return nil
	}, nil)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return nil, err
	}
	return updateUser, nil
}

func (u *User) UserToUser() *app.User {
	user := &app.User{}
	user.ID = u.ID
	user.IDStr = fmt.Sprintf("%v", u.ID)
	user.Name = u.Name
	return user
}

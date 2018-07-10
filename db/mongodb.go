package db

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
)

const (
	host      = "127.0.0.1:27017"
	authdb    = "admin"
	user      = "user"
	pass      = "123456"
	timeout   = 60 * time.Second
	poollimit = 4096
)

var globalS *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{host},
		Timeout:   timeout,
		Source:    authdb,
		Username:  user,
		Password:  pass,
		PoolLimit: poollimit,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal("create session error", err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalS.Copy()
	c := ms.DB(db).C(collection)
	return ms, c
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func FindAllSort(db, collection, sort string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Sort(sort).All(result)
}

func FindWithPage(db, collection, sort string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Sort(sort).Skip(page * limit).Limit(limit).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(selector, update)
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	_, err := c.UpdateAll(selector, update)
	return err
}

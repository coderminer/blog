package db

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
)

const (
	host      = "127.0.0.1:27017"
	authdb    = "Blog"
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

// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

// core/drive
package mongo

import (
	"gopkg.in/mgo.v2"
)

const (
	Host     = "mongo"
	Username = "user.gerbo"
	Password = "12345"
	Database = "gerbo"
)

func session() *mgo.Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{Host},
		Username: Username,
		Password: Password,
		Database: Database,
	})
	if err != nil {
		panic(err)
	}
	defer session.Clone()
	return session
}

func GetSession(collection string) *mgo.Collection {
	session := session()
	coll := session.DB(Database).C(collection)
	return coll
}

func GetDB() *mgo.Database {
	session := session()
	db := session.DB(Database)
	return db
}
package mgo

import (
	"errors"
	"log"
)

func fatal(err error) {
	log.Fatalf("mgo: %s", err.Error())
}

func checkAgainstNil(val interface{}) {
	if val == nil {
		log.Fatalf("mgo: %#v is nil", val)
	}
}

func checkAgainstEmpty(key, val string) {
	if val == "" {
		log.Fatalf("mgo: %s is empty", key)
	}
}

func ThrowNilCollectionError() error {
	return errors.New("mgo: collection is nil")
}

func ThrowDecodeModelIsNilError() error {
	return errors.New("mgo: decode model is nil")
}

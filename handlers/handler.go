package handlers

import mgo "gopkg.in/mgo.v2"

type Handler struct {
	DBSession *mgo.Session
}

package domain

import "time"

type Log struct {
	Id        string    `bson:"_id"`
	Request   string    `bson:"METHOD"`
	Timestamp time.Time `bson:"timestamp"`
}

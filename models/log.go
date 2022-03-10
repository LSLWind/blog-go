package models

import "time"

type Log struct {
	Id         uint
	Type       int
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
}

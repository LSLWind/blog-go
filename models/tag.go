package models

import "time"

type Tag struct {
	Id         uint
	Name       string
	Icon       string
	CreateTime time.Time
	UpdateTime time.Time
}

package model

import "time"

type Currency struct {
	Id    uint
	Title string
	Code  string
	Value string
	Date  time.Time
}

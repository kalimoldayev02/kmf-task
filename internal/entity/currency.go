package entity

import "time"

type Currency struct {
	ID    uint
	Title string
	Code  string
	Value string
	Date  time.Time
}

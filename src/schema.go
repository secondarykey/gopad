package main

//+AR
type Memo struct {
	Id    int `db:"pk"`
	Title string
	Memo  string
}

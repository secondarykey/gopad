package main

//+AR
type Memo struct {
	Id    string `db:"pk"`
	Title string
	Memo  string
}

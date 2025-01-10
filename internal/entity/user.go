package entity

type User struct {
	UUID    string `db:"uuid"`
	Balance int    `db:"balance"`
}

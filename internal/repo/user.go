package repo

type User interface{
	GetBalance(uuid string)
}
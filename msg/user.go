package msg

type Account struct {
	ID      uint
	Account string
	Passwd  string
}

type User struct {
	ID   int64
	Name string
}

type CreateRole struct {
	Name string
	Sex  int8
}

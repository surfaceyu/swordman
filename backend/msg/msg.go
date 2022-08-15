package msg

type Account struct {
	ID      uint
	Account string
	Passwd  string
}

type Server struct {
	ID   int16
	Name string
	Host string
	Port int16
}

type User struct {
	ID   int64
	Name string
}

type CreateRole struct {
	Name string
	Sex  int8
}

type SendChatMsg struct {
	Channel string
	To      int64
	Data    string
}

type GetChatMsg struct {
	Channel string
	To      int64
	Limit   int
}

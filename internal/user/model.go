package user

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

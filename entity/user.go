package entity

type User struct {
	NickName string `json:"nick_name,omitempty"`
	HeadImg  string `json:"head_img,omitempty"`
}

func (this *User) CollName() string {
	return "user_data"
}

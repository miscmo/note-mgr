package entity

type Comment struct {
	Content     string   `bson:"content,omitempty"`
	UserId      string   `bson:"user_id,omitempty"`
	CreateTime  string   `bson:"create_time,omitempty"`
	UpdateTime  string   `bson:"update_time,omitempty"`
	SubComments []string `bson:"sub_comments,omitempty"`
}

func (this *Comment) CollName() string {
	return "comment_data"
}

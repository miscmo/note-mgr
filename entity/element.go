package entity

type Element struct {
	Type        int      `bson:"type,omitempty"`
	Content     string   `bson:"content,omitempty"`
	SubElements []string `bson:"sub_elements,omitempty"`
	Comment     string   `bson:"comment,omitempty"`
	CreateTime  string   `bson:"create_time,omitempty"`
	UpdateTime  string   `bson:"update_time,omitempty"`
}

func (this *Element) CollName() string {
	return "element_data"
}

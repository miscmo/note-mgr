package entity

type Element struct {
	ID         string   `bson:"_id,omitempty"`
	Type       int      `bson:"type,omitempty"`
	Content    string   `bson:"content,omitempty"`
	SubElems   []string `bson:"sub_elements,omitempty"`
	RefElems   []string `bson:"ref_elements,omitempty"`
	Comment    string   `bson:"comment,omitempty"`
	CreateTime string   `bson:"create_time,omitempty"`
	UpdateTime string   `bson:"update_time,omitempty"`
}

func (this *Element) CollName() string {
	return "element_data"
}

func (this *Element) ToProto() *ElemInfo {
	return &ElemInfo{
		ID:         this.ID,
		Content:    this.Content,
		CreateTime: this.CreateTime,
		UpdateTime: this.UpdateTime,
	}
}

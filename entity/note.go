package entity

const (
	ElemData = iota
	ElemFake
)

type Note struct {
	Name       string `bson:"name,omitempty"`
	Root       string `bson:"root,omitempty"`
	Content    string `bson:"content,omitempty"`
	CreateTime string `bson:"create_time,omitempty"`
	UpdateTime string `bson:"update_time,omitempty"`
}

func (this *Note) CollName() string {
	return "note_data"
}

package entity

const (
	ElemData = iota
	ElemFake
)

type Note struct {
	ID         string `bson:"_id,omitempty"`
	Name       string `bson:"name,omitempty"`
	Root       string `bson:"root,omitempty"`
	Content    string `bson:"content,omitempty"`
	CreateTime string `bson:"create_time,omitempty"`
	UpdateTime string `bson:"update_time,omitempty"`
}

func (this *Note) CollName() string {
	return "note_data"
}

func (this *Note) ToProto() *NoteInfo {
	return &NoteInfo{
		ID:         this.ID,
		CreateTime: this.CreateTime,
		UpdateTime: this.UpdateTime,
	}
}

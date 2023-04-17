package entity

type SaveNoteReq struct {
	Content string `json:"content,omitempty"`
}

type SaveNoteRsp struct {
	Code int32  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type UpdateNoteReq struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type UpdateNoteRsp struct {
	Code int32  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type ElemInfo struct {
	ID         string     `json:"id,omitempty"`
	Content    string     `json:"content,omitempty"`
	CreateTime string     `json:"create_time,omitempty"`
	UpdateTime string     `json:"update_time,omitempty"`
	SubElems   []ElemInfo `json:"sub_elems,omitempty"`
}

type NoteInfo struct {
	ID         string   `json:"id,omitempty"`
	RootElem   ElemInfo `json:"root_elem,omitempty"`
	CreateTime string   `json:"create_time,omitempty"`
	UpdateTime string   `json:"update_time,omitempty"`
}

type GetNoteReq struct {
	ID string `json:"id,omitempty"`
}

type GetNoteRsp struct {
	Items []NoteInfo `json:"items,omitempty"`
	Total int32      `json:"total,omitempty"`
	Code  int32      `json:"code,omitempty"`
	Msg   string     `json:"msg,omitempty"`
}

type GetElemsReq struct {
}

type GetElemsRsp struct {
}

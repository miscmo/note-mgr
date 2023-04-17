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

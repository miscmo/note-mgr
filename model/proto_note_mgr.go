package model

type SaveNoteReq struct {
	ID      int64  `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type SaveNoteRsp struct {
	Code int32  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

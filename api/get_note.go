package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miscmo/note-mgr/entity"
	"github.com/miscmo/note-mgr/repository"
)

func GetNotes(c *gin.Context) {
	req := entity.GetNoteReq{}
	rsp := entity.GetNoteRsp{}

	var (
		err  error
		list []entity.Note
	)

	body := c.Request.Body

	decoder := json.NewDecoder(body)
	if err = decoder.Decode(&req); err != nil {
		rsp.Code = entity.ErrParamInvaild
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("decode failed, err: %+v, rsp: %v\n", err, rsp)
		return
	}

	fmt.Printf("req: %+v", req)

	list, err = repository.GetNoteMgr().Find()
	if err != nil {
		rsp.Code = entity.ErrFailed
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("get notes failed, rsp: %v\n", rsp)
		return

	}

	fmt.Printf("get note list: %+v", list)

	var result []entity.NoteInfo

	for _, v := range list {
		root, err := GetElems(v.Root)
		if err != nil {
			fmt.Printf("get root elems failed, err: %v", err)
		}

		noteInfo := entity.NoteInfo{
			ID:         v.ID,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			RootElem:   *root,
		}

		result = append(result, noteInfo)
	}

	rsp.Items = result

	fmt.Printf("get note rsp: %+v", rsp)

	c.JSON(http.StatusOK, rsp)
}

func GetElems(id string) (*entity.ElemInfo, error) {
	var ret entity.ElemInfo

	if len(id) <= 0 {
		return &ret, nil
	}

	elem, err := repository.GetElementMgr().GetElem(id)
	if err != nil {
		return &ret, err
	}

	var subElems []entity.ElemInfo

	for _, e := range elem.SubElems {

		e, _ := GetElems(e)

		subElems = append(subElems, *e)
	}

	ret.SubElems = subElems

	return &ret, nil
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miscmo/note-mgr/entity"
	"github.com/miscmo/note-mgr/repository"
)

func UpdateNote(c *gin.Context) {
	req := entity.UpdateNoteReq{}
	rsp := entity.UpdateNoteRsp{}

	var (
		err error
	)

	body := c.Request.Body

	defer c.JSON(http.StatusOK, rsp)

	decoder := json.NewDecoder(body)
	if err = decoder.Decode(&req); err != nil {
		rsp.Code = entity.ErrParamInvaild
		rsp.Msg = err.Error()
		fmt.Printf("decode failed, err: %+v, rsp: %v\n", err, rsp)
		return
	}

	fmt.Printf("req: %+v", req)

	note := entity.Note{
		Content: req.Content,
	}

	id, err := repository.GetNoteMgr().InsertOne(note)
	if err != nil {
		rsp.Code = entity.ErrFailed
		rsp.Msg = err.Error()
		fmt.Printf("decode failed, rsp: %v\n", rsp)
		return

	}

	fmt.Printf("insert id: %v", id)

}

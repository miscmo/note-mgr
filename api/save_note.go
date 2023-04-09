package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miscmo/note-mgr/model"
)

// @Summary
//

func SaveNote(c *gin.Context) {
	req := model.SaveNoteReq{}
	rsp := model.SaveNoteRsp{}

	var (
		err error
	)

	body := c.Request.Body

	decoder := json.NewDecoder(body)
	if err = decoder.Decode(&req); err != nil {
		rsp.Code = model.ErrParamInvaild
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("decode failed, err: %+v, rsp: %v\n", err, rsp)
		return
	}

	fmt.Printf("req: %+v", req)

	// _, err = dao.InsertNode(req.Node)
	// if err != nil {
	// 	rsp.Code = model.UpdateDBFailed
	// 	rsp.Msg = err.Error()
	// 	c.JSON(http.StatusOK, rsp)
	// 	fmt.Printf("decode failed, rsp: %v\n", rsp)
	// 	return
	// }

	c.JSON(http.StatusOK, rsp)
}

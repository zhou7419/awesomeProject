package note

import (
	"awesomeProject/app/request/noteReq"
	"awesomeProject/app/service/note"
	pkgJwt "awesomeProject/pkg/jwt"
	"awesomeProject/pkg/response"
	"github.com/gin-gonic/gin"
)

var ns = new(noteService.NoteService)

func Add(c *gin.Context) {
	var addReq noteReq.Add
	if err := c.Bind(&addReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	jwtClaims, exists := c.Get("claims")
	if !exists {
		response.Unauthorized(c, "token解析失败")
		return
	}

	claims, ok := jwtClaims.(*pkgJwt.CustomClaims)
	if !ok {
		response.Unauthorized(c, "token解析失败2")
		return
	}

	if err := ns.Add(claims.UserID, addReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

func Index(c *gin.Context) {
	var indexReq noteReq.Index
	if err := c.Bind(&indexReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	list, total, err := ns.Index(indexReq)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Data(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func Detail(c *gin.Context) {
	var detailReq noteReq.Detail
	if err := c.Bind(&detailReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	data, err := ns.GetById(detailReq.ID)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Data(c, data)
}

func Update(c *gin.Context) {
	var updateReq noteReq.Update
	if err := c.Bind(&updateReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := ns.Update(updateReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

func Delete(c *gin.Context) {
	var deleteReq noteReq.Delete
	if err := c.Bind(&deleteReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := ns.Delete(deleteReq.ID); err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

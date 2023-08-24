package controller

import (
	"github.com/gin-gonic/gin"
	"mysqlAndRedis/biz"
	req "mysqlAndRedis/request/student"
	"mysqlAndRedis/resp"
	"net/http"
)

func GetStudentById(c *gin.Context) {
	param := req.GetStudentByIdParam{}
	response := &resp.Response{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Code = 10001
		response.Message = "bind param failed: " + err.Error()
		response.Data = []interface{}{}
		c.JSON(http.StatusOK, response)
		return
	}

	id := *param.Id
	studentBiz := &biz.Student{}
	err = studentBiz.GetById(id)

	// 查询出错
	if err != nil {
		response.Code = 10002
		response.Message = "get student by id failed: " + err.Error()
		response.Data = []interface{}{}
		c.JSON(http.StatusOK, response)
		return
	}

	// 没查到数据
	if studentBiz.Id == 0 {
		response.Code = 10003
		response.Message = "student not found"
		response.Data = []interface{}{}
		c.JSON(http.StatusOK, response)
		return
	}

	response.Code = 200
	response.Message = "success"
	response.Data = []interface{}{studentBiz}
	c.JSON(http.StatusOK, response)
	return
}

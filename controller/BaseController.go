package controller

import (
	"ycoem/lib/request"
	"ycoem/lib/response"
)

type BaseController struct {
	response.Response
	request.Request
}

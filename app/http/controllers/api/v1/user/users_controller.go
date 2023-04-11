package user

import (
	v1 "mall/app/http/controllers/api/v1"
	"mall/app/models/user"
	"mall/app/requests"
	"mall/pkg/auth"
	"mall/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	v1.BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	// data := user.All()
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c,&request,requests.Pagination);!ok{
		return 
	}
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
        "data":  data,
        "pager": pager,
    })
	
}

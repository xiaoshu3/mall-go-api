package auth

import (
	v1 "mall/app/http/controllers/api/v1"
	"mall/app/models/user"
	"mall/app/requests"
	"mall/pkg/helpers"
	"mall/pkg/jwt"
	"mall/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// panic("panic 测试")
	// 	请求对象
	// type PhoneExistRequest struct {
	// 	Phone string `json:"phone"`
	// }

	request := requests.SignupPhoneExistRequest{}

	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// //表单验证

	// errs := requests.ValidateSignupPhoneExist(&request, c)
	// // errs 返回长度等于零即通过，大于 0 即有错误发生
	// if len(errs) > 0 {
	// 	// 验证失败，返回 422 状态码和错误信息
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
	// 		"errors": errs,
	// 	})
	// 	return
	// }

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"exist": user.IsPhoneExist(request.Phone),
	// })

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// logger.Du
	// 2. 验证成功，创建数据
	userModel := user.User{
		// Name:     request.Name,
		Name:     helpers.PhoneToName(request.Phone),
		Phone:    request.Phone,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		// response.Created(c, userModel)
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}

}

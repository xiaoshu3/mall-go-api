// Package factories 存放工厂方法
package factories

import (
	"mall/app/models/user"
	"mall/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name: faker.Username(),
			// Email:    faker.Email(),
			Phone: helpers.RandomNumber(11),
			// Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
			Password: "123456",
		}
		objs = append(objs, model)
	}

	return objs
}

package handler

import (
	"github.com/hhhhp52/webtest/src/error"
	"github.com/hhhhp52/webtest/src/service"
	"github.com/kataras/iris/v12"
)

// GetAllMemberHandler return hello message
func GetAllMemberHandler(ctx iris.Context) {

	result, err := service.GetAllMember()

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}

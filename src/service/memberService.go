package service

import (
	dto "github.com/hhhhp52/webtest/src/dto/memberDto"
	"github.com/hhhhp52/webtest/src/error"
	gormdao "github.com/hhhhp52/webtest/src/persistence/gorm"
	memberdao "github.com/hhhhp52/webtest/src/persistence/gorm/memberDao"
	"github.com/hhhhp52/webtest/src/utils/logger"
)

//GetAllMember is to get all user
func GetAllMember() (result []map[string]interface{}, e *error.Error) {
	tx := gormdao.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	memberlist := memberdao.GetAll(tx)
	if len(memberlist) == 0 {
		return nil, error.DataNotFoundError("no member")
	}

	result = make([]map[string]interface{}, 0)

	result = dto.MemberDto(&memberlist)
	return result, nil
}

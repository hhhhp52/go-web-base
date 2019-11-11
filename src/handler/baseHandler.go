package handler

import (
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/hhhhp52/webtest/src/error"
	"github.com/hhhhp52/webtest/src/utils/time"
	"github.com/hhhhp52/webtest/src/utils/typecast"
	"github.com/kataras/iris/v12"
)

// custome validator
func init() {
	// validate minimum
	govalidator.ParamTagMap["min"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		min := typecast.StringToFloat64(params[0])
		number := typecast.StringToFloat64(str)
		return number >= min
	})
	govalidator.ParamTagRegexMap["min"] = regexp.MustCompile("^min\\((\\w+)\\)$")

	// validate maximum
	govalidator.ParamTagMap["max"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		max := typecast.StringToFloat64(params[0])
		number := typecast.StringToFloat64(str)
		return number <= max
	})
	govalidator.ParamTagRegexMap["max"] = regexp.MustCompile("^max\\((\\w+)\\)$")

	// validate password
	// https://github.com/asaskevich/govalidator/issues/261
	// https://github.com/StefanSchroeder/Golang-Regex-Tutorial/issues/11
	govalidator.CustomTypeTagMap.Set("password", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		hasAlphabetic := govalidator.Matches(i.(string), "^.*[a-zA-Z]{1,}.*$")
		hasNumber := govalidator.Matches(i.(string), "^.*[0-9]{1,}.*$")
		checkLength := govalidator.StringLength(i.(string), "6", "16")
		return checkLength && hasAlphabetic && hasNumber
	}))

	// validate timestamp
	govalidator.CustomTypeTagMap.Set("timestamp", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		t, ok := i.(time.Time)

		if !ok {
			return false
		}
		return t.Error == nil
	}))
}

// success
func success(ctx iris.Context, data interface{}) {
	ctx.JSON(iris.Map{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

// failed
func failed(ctx iris.Context, err *error.Error) {
	ctx.JSON(iris.Map{
		"code":    err.Code(),
		"message": err.Error(),
		"data":    []string{},
	})
}

//raw
func raw(ctx iris.Context, data interface{}) {
	ctx.Text(fmt.Sprintf("%v", data))
}

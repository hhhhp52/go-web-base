package dto

import "github.com/hhhhp52/webtest/src/domain"

//MemberDto to set the data to front-end
func MemberDto(members *[]domain.User) (result []map[string]interface{}) {
	for _, member := range *members {
		result = append(
			result,
			map[string]interface{}{
				"Account":  member.Account,
				"Name":     member.Name,
				"NickName": member.Nickname,
			},
		)
	}
	return result
}

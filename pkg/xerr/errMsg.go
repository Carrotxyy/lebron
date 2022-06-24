/**
 * @Author: Jacky
 * @Description: 错误描述
 * @File: errMsg
 * @Version: 1.0.0
 * @Date: 2022/6/24 16:32
 */
package xerr

import "fmt"

var message = make(map[uint32]string)

func init() {
	set(OK, "SUCCESS")
	set(ServerCommonError, "服务器开小差啦,稍后再来试一试")
}

func set(errCode uint32, errMsg string) uint32 {
	if _, ok := message[errCode]; ok {
		panic(fmt.Sprintf("错误码:%d已存在", errCode))
	}
	message[errCode] = errMsg
	return errCode
}

func GetErrorMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

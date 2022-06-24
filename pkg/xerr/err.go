/**
 * @Author: Jacky
 * @Description: 错误体
 * @File: err
 * @Version: 1.0.0
 * @Date: 2022/6/24 16:36
 */
package xerr

import "fmt"

type Xerr struct {
	errCode uint32
	errMsg  string
}

func (e *Xerr) GetErrCode() uint32 {
	return e.errCode
}

func (e *Xerr) GetErrMsg() string {
	return e.errMsg
}

func (e *Xerr) Error() string {
	return fmt.Sprintf("ErrCode:%d,ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *Xerr {
	return &Xerr{
		errCode: errCode,
		errMsg:  errMsg,
	}
}

func NewErrCode(errCode uint32) *Xerr {
	return &Xerr{
		errCode: errCode,
		errMsg:  GetErrorMsg(errCode),
	}
}

func NewErrMsg(errMsg string) *Xerr {
	return &Xerr{
		errCode: ServerCommonError,
		errMsg:  errMsg,
	}
}

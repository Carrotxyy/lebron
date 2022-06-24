/**
 * @Author: Jacky
 * @Description:
 * @File: response
 * @Version: 1.0.0
 * @Date: 2022/6/24 17:28
 */
package response

import (
	"lebron/pkg/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"google.golang.org/grpc/status"

	"github.com/pkg/errors"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//
//  Response
//  @Description: 包装响应数据
//  @param r
//  @param w
//  @param data
//  @param err
//
func Response(r *http.Request, w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		ResponseErr(r, w, err)
	} else {
		ResponseSuccess(w, data)
	}
}

//
//  ResponseNotWrap
//  @Description: 不包装
//  @param r
//  @param w
//  @param body
//  @param err
//
func ResponseNotWrap(r *http.Request, w http.ResponseWriter, body interface{}, err error) {
	if err != nil {
		ResponseErr(r, w, err)
	} else {
		httpx.OkJson(w, body)
	}
}

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	body := Body{
		Code: xerr.OK,
		Msg:  xerr.GetErrorMsg(xerr.OK),
		Data: data,
	}

	httpx.OkJson(w, body)
}

func ResponseErr(r *http.Request, w http.ResponseWriter, err error) {
	body := Body{
		Code: xerr.ServerCommonError,
		Msg:  xerr.GetErrorMsg(xerr.ServerCommonError),
	}

	causeErr := errors.Cause(err)
	if e, ok := causeErr.(*xerr.Xerr); ok {
		body.Code = e.GetErrCode()
		body.Msg = e.GetErrMsg()
	} else {
		// grpc
		if gstatus, ok := status.FromError(causeErr); ok {
			body.Code = uint32(gstatus.Code())
			body.Msg = gstatus.Message()
		}
	}

	logx.WithContext(r.Context()).Errorf("API-ERR : %+v", err)
	httpx.WriteJson(w, http.StatusBadRequest, body)
}

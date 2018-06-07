package iot

import (
	"github.com/morlay/goaliyun"
)

type BatchRegisterDeviceWithApplyIdRequest struct {
	ApplyId    int64  `position:"Query" name:"ApplyId"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *BatchRegisterDeviceWithApplyIdRequest) Invoke(client goaliyun.Client) (*BatchRegisterDeviceWithApplyIdResponse, error) {
	resp := &BatchRegisterDeviceWithApplyIdResponse{}
	err := client.Request("iot", "BatchRegisterDeviceWithApplyId", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchRegisterDeviceWithApplyIdResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         BatchRegisterDeviceWithApplyIdData
}

type BatchRegisterDeviceWithApplyIdData struct {
	ApplyId goaliyun.Integer
}

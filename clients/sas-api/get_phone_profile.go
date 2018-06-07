package sas_api

import (
	"github.com/morlay/goaliyun"
)

type GetPhoneProfileRequest struct {
	Phone        string `position:"Query" name:"Phone"`
	SensType     int64  `position:"Query" name:"SensType"`
	DataVersion  string `position:"Query" name:"DataVersion"`
	BusinessType int64  `position:"Query" name:"BusinessType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *GetPhoneProfileRequest) Invoke(client goaliyun.Client) (*GetPhoneProfileResponse, error) {
	resp := &GetPhoneProfileResponse{}
	err := client.Request("sas-api", "GetPhoneProfile", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPhoneProfileResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	Success   bool
	RequestId goaliyun.String
	Data      GetPhoneProfileData
}

type GetPhoneProfileData struct {
	Phone goaliyun.String
	Info  goaliyun.String
}

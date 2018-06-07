package sas_api

import (
	"github.com/morlay/goaliyun"
)

type GetAccountProfileRequest struct {
	DeviceIdMd5     string `position:"Query" name:"DeviceIdMd.5"`
	Carrier         int64  `position:"Query" name:"Carrier"`
	Os              string `position:"Query" name:"Os"`
	Phone           string `position:"Query" name:"Phone"`
	RequestUrl      string `position:"Query" name:"RequestUrl"`
	Ip              string `position:"Query" name:"Ip"`
	UserAgent       string `position:"Query" name:"UserAgent"`
	ConnectionType  int64  `position:"Query" name:"ConnectionType"`
	SensType        int64  `position:"Query" name:"SensType"`
	DeviceType      int64  `position:"Query" name:"DeviceType"`
	AccessTimestamp int64  `position:"Query" name:"AccessTimestamp"`
	BusinessType    int64  `position:"Query" name:"BusinessType"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetAccountProfileRequest) Invoke(client goaliyun.Client) (*GetAccountProfileResponse, error) {
	resp := &GetAccountProfileResponse{}
	err := client.Request("sas-api", "GetAccountProfile", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAccountProfileResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	Success   bool
	RequestId goaliyun.String
	Data      GetAccountProfileData
}

type GetAccountProfileData struct {
	Ip        goaliyun.String
	Phone     goaliyun.String
	IpInfo    goaliyun.String
	PhoneInfo goaliyun.String
}

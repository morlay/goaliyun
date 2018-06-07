package sas_api

import (
	"github.com/morlay/goaliyun"
)

type GetIpProfileRequest struct {
	DeviceIdMd5    string `position:"Query" name:"DeviceIdMd.5"`
	Carrier        int64  `position:"Query" name:"Carrier"`
	Os             string `position:"Query" name:"Os"`
	RequestUrl     string `position:"Query" name:"RequestUrl"`
	Ip             string `position:"Query" name:"Ip"`
	UserAgent      string `position:"Query" name:"UserAgent"`
	ConnectionType int64  `position:"Query" name:"ConnectionType"`
	SensType       int64  `position:"Query" name:"SensType"`
	DeviceType     int64  `position:"Query" name:"DeviceType"`
	BusinessType   int64  `position:"Query" name:"BusinessType"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *GetIpProfileRequest) Invoke(client goaliyun.Client) (*GetIpProfileResponse, error) {
	resp := &GetIpProfileResponse{}
	err := client.Request("sas-api", "GetIpProfile", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetIpProfileResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	Success   bool
	RequestId goaliyun.String
	Data      GetIpProfileData
}

type GetIpProfileData struct {
	Ip   goaliyun.String
	Info goaliyun.String
}

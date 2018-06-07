package ccc

import (
	"github.com/morlay/goaliyun"
)

type RequestLoginInfoRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RequestLoginInfoRequest) Invoke(client goaliyun.Client) (*RequestLoginInfoResponse, error) {
	resp := &RequestLoginInfoResponse{}
	err := client.Request("ccc", "RequestLoginInfo", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RequestLoginInfoResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	LoginInfo      RequestLoginInfoLoginInfo
}

type RequestLoginInfoLoginInfo struct {
	UserName       goaliyun.String
	DisplayName    goaliyun.String
	PhoneNumber    goaliyun.String
	Region         goaliyun.String
	WebRtcUrl      goaliyun.String
	AgentServerUrl goaliyun.String
	Extension      goaliyun.String
	TenantId       goaliyun.String
	Signature      goaliyun.String
	SignData       goaliyun.String
}

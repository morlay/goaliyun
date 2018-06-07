package ccc

import (
	"github.com/morlay/goaliyun"
)

type GetConfigRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	ObjectType string `position:"Query" name:"ObjectType"`
	ObjectId   string `position:"Query" name:"ObjectId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetConfigRequest) Invoke(client goaliyun.Client) (*GetConfigResponse, error) {
	resp := &GetConfigResponse{}
	err := client.Request("ccc", "GetConfig", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetConfigResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	ConfigItem     GetConfigConfigItem
}

type GetConfigConfigItem struct {
	Name  goaliyun.String
	Value goaliyun.String
}

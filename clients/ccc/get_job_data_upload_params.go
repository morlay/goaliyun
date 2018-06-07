package ccc

import (
	"github.com/morlay/goaliyun"
)

type GetJobDataUploadParamsRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	FileName   string `position:"Query" name:"FileName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetJobDataUploadParamsRequest) Invoke(client goaliyun.Client) (*GetJobDataUploadParamsResponse, error) {
	resp := &GetJobDataUploadParamsResponse{}
	err := client.Request("ccc", "GetJobDataUploadParams", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetJobDataUploadParamsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	UploadParams   GetJobDataUploadParamsUploadParams
}

type GetJobDataUploadParamsUploadParams struct {
	AccessId  goaliyun.String
	Policy    goaliyun.String
	Signature goaliyun.String
	Folder    goaliyun.String
	Host      goaliyun.String
	Expire    goaliyun.Integer
}

package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetApgroupScanConfigSaveProgressRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetApgroupScanConfigSaveProgressRequest) Invoke(client goaliyun.Client) (*GetApgroupScanConfigSaveProgressResponse, error) {
	resp := &GetApgroupScanConfigSaveProgressResponse{}
	err := client.Request("cloudwf", "GetApgroupScanConfigSaveProgress", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetApgroupScanConfigSaveProgressResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}

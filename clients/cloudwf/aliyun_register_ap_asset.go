package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type AliyunRegisterApAssetRequest struct {
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	Mac       string `position:"Query" name:"Mac"`
	SerialNo  string `position:"Query" name:"SerialNo"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *AliyunRegisterApAssetRequest) Invoke(client goaliyun.Client) (*AliyunRegisterApAssetResponse, error) {
	resp := &AliyunRegisterApAssetResponse{}
	err := client.Request("cloudwf", "AliyunRegisterApAsset", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AliyunRegisterApAssetResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}

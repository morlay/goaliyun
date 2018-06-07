package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApDetailInfoRequest struct {
	ApAssetId int64  `position:"Query" name:"ApAssetId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListApDetailInfoRequest) Invoke(client goaliyun.Client) (*ListApDetailInfoResponse, error) {
	resp := &ListApDetailInfoResponse{}
	err := client.Request("cloudwf", "ListApDetailInfo", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApDetailInfoResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}

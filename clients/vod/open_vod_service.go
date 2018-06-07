package vod

import (
	"github.com/morlay/goaliyun"
)

type OpenVodServiceRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *OpenVodServiceRequest) Invoke(client goaliyun.Client) (*OpenVodServiceResponse, error) {
	resp := &OpenVodServiceResponse{}
	err := client.Request("vod", "OpenVodService", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OpenVodServiceResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}

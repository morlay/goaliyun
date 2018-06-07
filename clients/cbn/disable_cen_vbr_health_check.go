package cbn

import (
	"github.com/morlay/goaliyun"
)

type DisableCenVbrHealthCheckRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	VbrInstanceOwnerId   int64  `position:"Query" name:"VbrInstanceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrInstanceId        string `position:"Query" name:"VbrInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VbrInstanceRegionId  string `position:"Query" name:"VbrInstanceRegionId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DisableCenVbrHealthCheckRequest) Invoke(client goaliyun.Client) (*DisableCenVbrHealthCheckResponse, error) {
	resp := &DisableCenVbrHealthCheckResponse{}
	err := client.Request("cbn", "DisableCenVbrHealthCheck", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DisableCenVbrHealthCheckResponse struct {
	RequestId goaliyun.String
}

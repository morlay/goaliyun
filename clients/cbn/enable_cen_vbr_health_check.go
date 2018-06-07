package cbn

import (
	"github.com/morlay/goaliyun"
)

type EnableCenVbrHealthCheckRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	HealthCheckSourceIp  string `position:"Query" name:"HealthCheckSourceIp"`
	VbrInstanceOwnerId   int64  `position:"Query" name:"VbrInstanceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrInstanceId        string `position:"Query" name:"VbrInstanceId"`
	HealthCheckTargetIp  string `position:"Query" name:"HealthCheckTargetIp"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VbrInstanceRegionId  string `position:"Query" name:"VbrInstanceRegionId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *EnableCenVbrHealthCheckRequest) Invoke(client goaliyun.Client) (*EnableCenVbrHealthCheckResponse, error) {
	resp := &EnableCenVbrHealthCheckResponse{}
	err := client.Request("cbn", "EnableCenVbrHealthCheck", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnableCenVbrHealthCheckResponse struct {
	RequestId goaliyun.String
}

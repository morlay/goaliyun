package ess

import (
	"github.com/morlay/goaliyun"
)

type DeactivateScalingConfigurationRequest struct {
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *DeactivateScalingConfigurationRequest) Invoke(client goaliyun.Client) (*DeactivateScalingConfigurationResponse, error) {
	resp := &DeactivateScalingConfigurationResponse{}
	err := client.Request("ess", "DeactivateScalingConfiguration", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeactivateScalingConfigurationResponse struct {
	RequestId goaliyun.String
}

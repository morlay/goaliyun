package ess

import (
	"github.com/morlay/goaliyun"
)

type DeleteScalingConfigurationRequest struct {
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *DeleteScalingConfigurationRequest) Invoke(client goaliyun.Client) (*DeleteScalingConfigurationResponse, error) {
	resp := &DeleteScalingConfigurationResponse{}
	err := client.Request("ess", "DeleteScalingConfiguration", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteScalingConfigurationResponse struct {
	RequestId goaliyun.String
}

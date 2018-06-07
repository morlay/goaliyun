package afs

import (
	"github.com/morlay/goaliyun"
)

type CreateConfigurationRequest struct {
	ResourceOwnerId     int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp            string `position:"Query" name:"SourceIp"`
	ConfigurationName   string `position:"Query" name:"ConfigurationName"`
	MaxPV               string `position:"Query" name:"MaxPV"`
	ConfigurationMethod string `position:"Query" name:"ConfigurationMethod"`
	ApplyType           string `position:"Query" name:"ApplyType"`
	Scene               string `position:"Query" name:"Scene"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *CreateConfigurationRequest) Invoke(client goaliyun.Client) (*CreateConfigurationResponse, error) {
	resp := &CreateConfigurationResponse{}
	err := client.Request("afs", "CreateConfiguration", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateConfigurationResponse struct {
	RequestId goaliyun.String
	BizCode   goaliyun.String
}

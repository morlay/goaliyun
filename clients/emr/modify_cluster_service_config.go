package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyClusterServiceConfigRequest struct {
	ResourceOwnerId    int64  `position:"Query" name:"ResourceOwnerId"`
	CustomConfigParams string `position:"Query" name:"CustomConfigParams"`
	GroupId            int64  `position:"Query" name:"GroupId"`
	ServiceName        string `position:"Query" name:"ServiceName"`
	Comment            string `position:"Query" name:"Comment"`
	ClusterId          string `position:"Query" name:"ClusterId"`
	ConfigParams       string `position:"Query" name:"ConfigParams"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *ModifyClusterServiceConfigRequest) Invoke(client goaliyun.Client) (*ModifyClusterServiceConfigResponse, error) {
	resp := &ModifyClusterServiceConfigResponse{}
	err := client.Request("emr", "ModifyClusterServiceConfig", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyClusterServiceConfigResponse struct {
	RequestId goaliyun.String
}

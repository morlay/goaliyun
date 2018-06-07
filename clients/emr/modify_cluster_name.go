package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyClusterNameRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyClusterNameRequest) Invoke(client goaliyun.Client) (*ModifyClusterNameResponse, error) {
	resp := &ModifyClusterNameResponse{}
	err := client.Request("emr", "ModifyClusterName", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyClusterNameResponse struct {
	RequestId goaliyun.String
}

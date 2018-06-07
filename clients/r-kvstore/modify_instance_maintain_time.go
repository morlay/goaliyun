package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceMaintainTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	MaintainStartTime    string `position:"Query" name:"MaintainStartTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MaintainEndTime      string `position:"Query" name:"MaintainEndTime"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceMaintainTimeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceMaintainTimeResponse, error) {
	resp := &ModifyInstanceMaintainTimeResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceMaintainTime", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceMaintainTimeResponse struct {
	RequestId goaliyun.String
}

package dds

import (
	"github.com/morlay/goaliyun"
)

type ModifyNodeSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	NodeStorage          int64  `position:"Query" name:"NodeStorage"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeClass            string `position:"Query" name:"NodeClass"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	NodeId               string `position:"Query" name:"NodeId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyNodeSpecRequest) Invoke(client goaliyun.Client) (*ModifyNodeSpecResponse, error) {
	resp := &ModifyNodeSpecResponse{}
	err := client.Request("dds", "ModifyNodeSpec", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyNodeSpecResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}

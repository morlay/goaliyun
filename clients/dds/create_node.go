package dds

import (
	"github.com/morlay/goaliyun"
)

type CreateNodeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	NodeType             string `position:"Query" name:"NodeType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	NodeStorage          int64  `position:"Query" name:"NodeStorage"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeClass            string `position:"Query" name:"NodeClass"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateNodeRequest) Invoke(client goaliyun.Client) (*CreateNodeResponse, error) {
	resp := &CreateNodeResponse{}
	err := client.Request("dds", "CreateNode", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNodeResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}

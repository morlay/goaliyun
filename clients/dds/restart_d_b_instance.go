package dds

import (
	"github.com/morlay/goaliyun"
)

type RestartDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RestartDBInstanceRequest) Invoke(client goaliyun.Client) (*RestartDBInstanceResponse, error) {
	resp := &RestartDBInstanceResponse{}
	err := client.Request("dds", "RestartDBInstance", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestartDBInstanceResponse struct {
	RequestId goaliyun.String
}

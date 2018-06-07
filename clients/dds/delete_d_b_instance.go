package dds

import (
	"github.com/morlay/goaliyun"
)

type DeleteDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteDBInstanceRequest) Invoke(client goaliyun.Client) (*DeleteDBInstanceResponse, error) {
	resp := &DeleteDBInstanceResponse{}
	err := client.Request("dds", "DeleteDBInstance", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDBInstanceResponse struct {
	RequestId goaliyun.String
}

package rds

import (
	"github.com/morlay/goaliyun"
)

type DegradeDBInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int64  `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DegradeDBInstanceSpecRequest) Invoke(client goaliyun.Client) (*DegradeDBInstanceSpecResponse, error) {
	resp := &DegradeDBInstanceSpecResponse{}
	err := client.Request("rds", "DegradeDBInstanceSpec", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DegradeDBInstanceSpecResponse struct {
	RequestId goaliyun.String
}

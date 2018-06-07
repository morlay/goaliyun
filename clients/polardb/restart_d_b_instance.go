package polardb

import (
	"github.com/morlay/goaliyun"
)

type RestartDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RestartDBInstanceRequest) Invoke(client goaliyun.Client) (*RestartDBInstanceResponse, error) {
	resp := &RestartDBInstanceResponse{}
	err := client.Request("polardb", "RestartDBInstance", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestartDBInstanceResponse struct {
	RequestId goaliyun.String
}

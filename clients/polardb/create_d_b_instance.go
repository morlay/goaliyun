package polardb

import (
	"github.com/morlay/goaliyun"
)

type CreateDBInstanceRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CreateDBInstanceRequest) Invoke(client goaliyun.Client) (*CreateDBInstanceResponse, error) {
	resp := &CreateDBInstanceResponse{}
	err := client.Request("polardb", "CreateDBInstance", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDBInstanceResponse struct {
	RequestId    goaliyun.String
	DBClusterId  goaliyun.String
	DBInstanceId goaliyun.String
	OrderId      goaliyun.String
}

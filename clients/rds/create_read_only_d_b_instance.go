package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateReadOnlyDBInstanceRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int64  `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	VPCId                 string `position:"Query" name:"VPCId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CreateReadOnlyDBInstanceRequest) Invoke(client goaliyun.Client) (*CreateReadOnlyDBInstanceResponse, error) {
	resp := &CreateReadOnlyDBInstanceResponse{}
	err := client.Request("rds", "CreateReadOnlyDBInstance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateReadOnlyDBInstanceResponse struct {
	RequestId        goaliyun.String
	DBInstanceId     goaliyun.String
	OrderId          goaliyun.String
	ConnectionString goaliyun.String
	Port             goaliyun.String
}

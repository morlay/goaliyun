package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateDBInstanceRequest struct {
	ConnectionMode        string `position:"Query" name:"ConnectionMode"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int64  `position:"Query" name:"DBInstanceStorage"`
	SystemDBCharset       string `position:"Query" name:"SystemDBCharset"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	Engine                string `position:"Query" name:"Engine"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	DBInstanceStorageType string `position:"Query" name:"DBInstanceStorageType"`
	BusinessInfo          string `position:"Query" name:"BusinessInfo"`
	DBInstanceNetType     string `position:"Query" name:"DBInstanceNetType"`
	Period                string `position:"Query" name:"Period"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	UsedTime              string `position:"Query" name:"UsedTime"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	VPCId                 string `position:"Query" name:"VPCId"`
	TunnelId              string `position:"Query" name:"TunnelId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CreateDBInstanceRequest) Invoke(client goaliyun.Client) (*CreateDBInstanceResponse, error) {
	resp := &CreateDBInstanceResponse{}
	err := client.Request("rds", "CreateDBInstance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDBInstanceResponse struct {
	RequestId        goaliyun.String
	DBInstanceId     goaliyun.String
	OrderId          goaliyun.String
	ConnectionString goaliyun.String
	Port             goaliyun.String
}

package rds

import (
	"github.com/morlay/goaliyun"
)

type CloneDBInstanceRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	RestoreTime           string `position:"Query" name:"RestoreTime"`
	Period                string `position:"Query" name:"Period"`
	DBInstanceStorage     int64  `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	BackupId              string `position:"Query" name:"BackupId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	UsedTime              string `position:"Query" name:"UsedTime"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	ResourceGroupId       string `position:"Query" name:"ResourceGroupId"`
	VPCId                 string `position:"Query" name:"VPCId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CloneDBInstanceRequest) Invoke(client goaliyun.Client) (*CloneDBInstanceResponse, error) {
	resp := &CloneDBInstanceResponse{}
	err := client.Request("rds", "CloneDBInstance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CloneDBInstanceResponse struct {
	RequestId        goaliyun.String
	DBInstanceId     goaliyun.String
	OrderId          goaliyun.String
	ConnectionString goaliyun.String
	Port             goaliyun.String
}

package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateDBInstanceReplicaRequest struct {
	ConnectionMode        string `position:"Query" name:"ConnectionMode"`
	DomainMode            string `position:"Query" name:"DomainMode"`
	ReplicaDescription    string `position:"Query" name:"ReplicaDescription"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int64  `position:"Query" name:"DBInstanceStorage"`
	SystemDBCharset       string `position:"Query" name:"SystemDBCharset"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	Engine                string `position:"Query" name:"Engine"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
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
	SourceDBInstanceId    string `position:"Query" name:"SourceDBInstanceId"`
	ReplicaMode           string `position:"Query" name:"ReplicaMode"`
	VPCId                 string `position:"Query" name:"VPCId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CreateDBInstanceReplicaRequest) Invoke(client goaliyun.Client) (*CreateDBInstanceReplicaResponse, error) {
	resp := &CreateDBInstanceReplicaResponse{}
	err := client.Request("rds", "CreateDBInstanceReplica", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDBInstanceReplicaResponse struct {
	RequestId    goaliyun.String
	DBInstanceId goaliyun.String
	OrderId      goaliyun.Integer
	ReplicaId    goaliyun.String
	WorkflowId   goaliyun.String
}

package rds

import (
	"github.com/morlay/goaliyun"
)

type StartDBInstanceDiagnoseRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *StartDBInstanceDiagnoseRequest) Invoke(client goaliyun.Client) (*StartDBInstanceDiagnoseResponse, error) {
	resp := &StartDBInstanceDiagnoseResponse{}
	err := client.Request("rds", "StartDBInstanceDiagnose", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartDBInstanceDiagnoseResponse struct {
	RequestId      goaliyun.String
	DBInstanceName goaliyun.String
	DBInstanceId   goaliyun.String
}

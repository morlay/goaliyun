package rds

import (
	"github.com/morlay/goaliyun"
)

type SwitchDBInstanceNetTypeRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ConnectionStringType   string `position:"Query" name:"ConnectionStringType"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken            string `position:"Query" name:"ClientToken"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *SwitchDBInstanceNetTypeRequest) Invoke(client goaliyun.Client) (*SwitchDBInstanceNetTypeResponse, error) {
	resp := &SwitchDBInstanceNetTypeResponse{}
	err := client.Request("rds", "SwitchDBInstanceNetType", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SwitchDBInstanceNetTypeResponse struct {
	RequestId goaliyun.String
}

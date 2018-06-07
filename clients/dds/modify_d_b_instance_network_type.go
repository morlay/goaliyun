package dds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceNetworkTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	RetainClassic        string `position:"Query" name:"RetainClassic"`
	ClassicExpiredDays   int64  `position:"Query" name:"ClassicExpiredDays"`
	VpcId                string `position:"Query" name:"VpcId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceNetworkTypeRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceNetworkTypeResponse, error) {
	resp := &ModifyDBInstanceNetworkTypeResponse{}
	err := client.Request("dds", "ModifyDBInstanceNetworkType", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceNetworkTypeResponse struct {
	RequestId goaliyun.String
}

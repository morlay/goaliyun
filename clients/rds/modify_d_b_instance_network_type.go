package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceNetworkTypeRequest struct {
	ResourceOwnerId                      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount                 string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                         string `position:"Query" name:"OwnerAccount"`
	OwnerId                              int64  `position:"Query" name:"OwnerId"`
	VSwitchId                            string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress                     string `position:"Query" name:"PrivateIpAddress"`
	RetainClassic                        string `position:"Query" name:"RetainClassic"`
	ClassicExpiredDays                   string `position:"Query" name:"ClassicExpiredDays"`
	VPCId                                string `position:"Query" name:"VPCId"`
	DBInstanceId                         string `position:"Query" name:"DBInstanceId"`
	ReadWriteSplittingPrivateIpAddress   string `position:"Query" name:"ReadWriteSplittingPrivateIpAddress"`
	InstanceNetworkType                  string `position:"Query" name:"InstanceNetworkType"`
	ReadWriteSplittingClassicExpiredDays int64  `position:"Query" name:"ReadWriteSplittingClassicExpiredDays"`
	RegionId                             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceNetworkTypeRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceNetworkTypeResponse, error) {
	resp := &ModifyDBInstanceNetworkTypeResponse{}
	err := client.Request("rds", "ModifyDBInstanceNetworkType", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceNetworkTypeResponse struct {
	RequestId goaliyun.String
	TaskId    goaliyun.String
}

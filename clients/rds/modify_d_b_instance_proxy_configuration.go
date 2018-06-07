package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceProxyConfigurationRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	ProxyConfigurationKey   string `position:"Query" name:"ProxyConfigurationKey"`
	ProxyConfigurationValue string `position:"Query" name:"ProxyConfigurationValue"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceProxyConfigurationRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceProxyConfigurationResponse, error) {
	resp := &ModifyDBInstanceProxyConfigurationResponse{}
	err := client.Request("rds", "ModifyDBInstanceProxyConfiguration", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceProxyConfigurationResponse struct {
	RequestId goaliyun.String
}

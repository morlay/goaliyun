package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifySecurityIpsRequest struct {
	DBInstanceIPArrayName      string `position:"Query" name:"DBInstanceIPArrayName"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	ModifyMode                 string `position:"Query" name:"ModifyMode"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                string `position:"Query" name:"ClientToken"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	SecurityIps                string `position:"Query" name:"SecurityIps"`
	SecurityGroupId            string `position:"Query" name:"SecurityGroupId"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	WhitelistNetworkType       string `position:"Query" name:"WhitelistNetworkType"`
	DBInstanceIPArrayAttribute string `position:"Query" name:"DBInstanceIPArrayAttribute"`
	DBInstanceId               string `position:"Query" name:"DBInstanceId"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityIpsRequest) Invoke(client goaliyun.Client) (*ModifySecurityIpsResponse, error) {
	resp := &ModifySecurityIpsResponse{}
	err := client.Request("rds", "ModifySecurityIps", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityIpsResponse struct {
	RequestId goaliyun.String
	TaskId    goaliyun.String
}

package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifySecurityIpsRequest struct {
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ModifyMode               string `position:"Query" name:"ModifyMode"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	SecurityIps              string `position:"Query" name:"SecurityIps"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	SecurityIpGroupName      string `position:"Query" name:"SecurityIpGroupName"`
	InstanceId               string `position:"Query" name:"InstanceId"`
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	SecurityIpGroupAttribute string `position:"Query" name:"SecurityIpGroupAttribute"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityIpsRequest) Invoke(client goaliyun.Client) (*ModifySecurityIpsResponse, error) {
	resp := &ModifySecurityIpsResponse{}
	err := client.Request("r-kvstore", "ModifySecurityIps", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityIpsResponse struct {
	RequestId goaliyun.String
}

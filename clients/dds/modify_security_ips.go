package dds

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
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	SecurityIpGroupAttribute string `position:"Query" name:"SecurityIpGroupAttribute"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityIpsRequest) Invoke(client goaliyun.Client) (*ModifySecurityIpsResponse, error) {
	resp := &ModifySecurityIpsResponse{}
	err := client.Request("dds", "ModifySecurityIps", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityIpsResponse struct {
	RequestId goaliyun.String
}

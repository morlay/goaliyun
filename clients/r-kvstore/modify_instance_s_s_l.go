package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceSSLRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SSLEnabled           string `position:"Query" name:"SSLEnabled"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceSSLRequest) Invoke(client goaliyun.Client) (*ModifyInstanceSSLResponse, error) {
	resp := &ModifyInstanceSSLResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceSSL", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceSSLResponse struct {
	RequestId  goaliyun.String
	InstanceId goaliyun.String
	TaskId     goaliyun.String
}

package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type RestoreInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RestoreInstanceRequest) Invoke(client goaliyun.Client) (*RestoreInstanceResponse, error) {
	resp := &RestoreInstanceResponse{}
	err := client.Request("r-kvstore", "RestoreInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestoreInstanceResponse struct {
	RequestId goaliyun.String
}

package ossadmin

import (
	"github.com/morlay/goaliyun"
)

type StopOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *StopOssInstanceRequest) Invoke(client goaliyun.Client) (*StopOssInstanceResponse, error) {
	resp := &StopOssInstanceResponse{}
	err := client.Request("ossadmin", "StopOssInstance", "2015-03-02", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopOssInstanceResponse struct {
	RequestId goaliyun.String
}

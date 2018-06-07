package ossadmin

import (
	"github.com/morlay/goaliyun"
)

type ReleaseOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReleaseOssInstanceRequest) Invoke(client goaliyun.Client) (*ReleaseOssInstanceResponse, error) {
	resp := &ReleaseOssInstanceResponse{}
	err := client.Request("ossadmin", "ReleaseOssInstance", "2015-03-02", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseOssInstanceResponse struct {
	RequestId goaliyun.String
}

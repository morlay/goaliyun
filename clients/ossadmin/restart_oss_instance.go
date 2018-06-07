package ossadmin

import (
	"github.com/morlay/goaliyun"
)

type RestartOssInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Region               string `position:"Query" name:"Region"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RestartOssInstanceRequest) Invoke(client goaliyun.Client) (*RestartOssInstanceResponse, error) {
	resp := &RestartOssInstanceResponse{}
	err := client.Request("ossadmin", "RestartOssInstance", "2015-03-02", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestartOssInstanceResponse struct {
	RequestId goaliyun.String
}

package aegis

import (
	"github.com/morlay/goaliyun"
)

type ReleaseInstanceRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ReleaseInstanceRequest) Invoke(client goaliyun.Client) (*ReleaseInstanceResponse, error) {
	resp := &ReleaseInstanceResponse{}
	err := client.Request("aegis", "ReleaseInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseInstanceResponse struct {
	RequestId goaliyun.String
}

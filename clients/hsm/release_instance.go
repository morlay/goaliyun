package hsm

import (
	"github.com/morlay/goaliyun"
)

type ReleaseInstanceRequest struct {
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ReleaseInstanceRequest) Invoke(client goaliyun.Client) (*ReleaseInstanceResponse, error) {
	resp := &ReleaseInstanceResponse{}
	err := client.Request("hsm", "ReleaseInstance", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseInstanceResponse struct {
	RequestId goaliyun.String
}

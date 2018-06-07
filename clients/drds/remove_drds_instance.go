package drds

import (
	"github.com/morlay/goaliyun"
)

type RemoveDrdsInstanceRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *RemoveDrdsInstanceRequest) Invoke(client goaliyun.Client) (*RemoveDrdsInstanceResponse, error) {
	resp := &RemoveDrdsInstanceResponse{}
	err := client.Request("drds", "RemoveDrdsInstance", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveDrdsInstanceResponse struct {
	RequestId goaliyun.String
	Success   bool
}

package drds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDrdsInstanceDescriptionRequest struct {
	Description    string `position:"Query" name:"Description"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ModifyDrdsInstanceDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyDrdsInstanceDescriptionResponse, error) {
	resp := &ModifyDrdsInstanceDescriptionResponse{}
	err := client.Request("drds", "ModifyDrdsInstanceDescription", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDrdsInstanceDescriptionResponse struct {
	RequestId goaliyun.String
	Success   bool
}

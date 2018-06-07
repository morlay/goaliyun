package nas

import (
	"github.com/morlay/goaliyun"
)

type ModifyMountTargetRequest struct {
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	AccessGroupName   string `position:"Query" name:"AccessGroupName"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
	Status            string `position:"Query" name:"Status"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *ModifyMountTargetRequest) Invoke(client goaliyun.Client) (*ModifyMountTargetResponse, error) {
	resp := &ModifyMountTargetResponse{}
	err := client.Request("nas", "ModifyMountTarget", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyMountTargetResponse struct {
	RequestId goaliyun.String
}

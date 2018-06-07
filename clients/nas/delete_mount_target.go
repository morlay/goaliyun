package nas

import (
	"github.com/morlay/goaliyun"
)

type DeleteMountTargetRequest struct {
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *DeleteMountTargetRequest) Invoke(client goaliyun.Client) (*DeleteMountTargetResponse, error) {
	resp := &DeleteMountTargetResponse{}
	err := client.Request("nas", "DeleteMountTarget", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMountTargetResponse struct {
	RequestId goaliyun.String
}

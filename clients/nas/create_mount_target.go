package nas

import (
	"github.com/morlay/goaliyun"
)

type CreateMountTargetRequest struct {
	VSwitchId       string `position:"Query" name:"VSwitchId"`
	VpcId           string `position:"Query" name:"VpcId"`
	NetworkType     string `position:"Query" name:"NetworkType"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	FileSystemId    string `position:"Query" name:"FileSystemId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateMountTargetRequest) Invoke(client goaliyun.Client) (*CreateMountTargetResponse, error) {
	resp := &CreateMountTargetResponse{}
	err := client.Request("nas", "CreateMountTarget", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMountTargetResponse struct {
	RequestId         goaliyun.String
	MountTargetDomain goaliyun.String
}

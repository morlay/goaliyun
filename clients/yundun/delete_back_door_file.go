package yundun

import (
	"github.com/morlay/goaliyun"
)

type DeleteBackDoorFileRequest struct {
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Path       string `position:"Query" name:"Path"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteBackDoorFileRequest) Invoke(client goaliyun.Client) (*DeleteBackDoorFileResponse, error) {
	resp := &DeleteBackDoorFileResponse{}
	err := client.Request("yundun", "DeleteBackDoorFile", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBackDoorFileResponse struct {
	RequestId goaliyun.String
}

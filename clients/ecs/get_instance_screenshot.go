package ecs

import (
	"github.com/morlay/goaliyun"
)

type GetInstanceScreenshotRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Wakeup               string `position:"Query" name:"Wakeup"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetInstanceScreenshotRequest) Invoke(client goaliyun.Client) (*GetInstanceScreenshotResponse, error) {
	resp := &GetInstanceScreenshotResponse{}
	err := client.Request("ecs", "GetInstanceScreenshot", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetInstanceScreenshotResponse struct {
	RequestId  goaliyun.String
	InstanceId goaliyun.String
	Screenshot goaliyun.String
}

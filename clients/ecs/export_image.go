package ecs

import (
	"github.com/morlay/goaliyun"
)

type ExportImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	OSSBucket            string `position:"Query" name:"OSSBucket"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OSSPrefix            string `position:"Query" name:"OSSPrefix"`
	RoleName             string `position:"Query" name:"RoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ImageFormat          string `position:"Query" name:"ImageFormat"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ExportImageRequest) Invoke(client goaliyun.Client) (*ExportImageResponse, error) {
	resp := &ExportImageResponse{}
	err := client.Request("ecs", "ExportImage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ExportImageResponse struct {
	RequestId goaliyun.String
	TaskId    goaliyun.String
	RegionId  goaliyun.String
}

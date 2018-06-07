package vod

import (
	"github.com/morlay/goaliyun"
)

type DeleteStreamRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteStreamRequest) Invoke(client goaliyun.Client) (*DeleteStreamResponse, error) {
	resp := &DeleteStreamResponse{}
	err := client.Request("vod", "DeleteStream", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteStreamResponse struct {
	RequestId goaliyun.String
}

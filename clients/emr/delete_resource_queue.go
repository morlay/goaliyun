package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteResourceQueueRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceQueueId string `position:"Query" name:"ResourceQueueId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteResourceQueueRequest) Invoke(client goaliyun.Client) (*DeleteResourceQueueResponse, error) {
	resp := &DeleteResourceQueueResponse{}
	err := client.Request("emr", "DeleteResourceQueue", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteResourceQueueResponse struct {
	RequestId goaliyun.String
}

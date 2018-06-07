package mts

import (
	"github.com/morlay/goaliyun"
)

type UnbindOutputBucketRequest struct {
	Bucket               string `position:"Query" name:"Bucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UnbindOutputBucketRequest) Invoke(client goaliyun.Client) (*UnbindOutputBucketResponse, error) {
	resp := &UnbindOutputBucketResponse{}
	err := client.Request("mts", "UnbindOutputBucket", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnbindOutputBucketResponse struct {
	RequestId goaliyun.String
}

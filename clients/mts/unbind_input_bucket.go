package mts

import (
	"github.com/morlay/goaliyun"
)

type UnbindInputBucketRequest struct {
	Bucket               string `position:"Query" name:"Bucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RoleArn              string `position:"Query" name:"RoleArn"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UnbindInputBucketRequest) Invoke(client goaliyun.Client) (*UnbindInputBucketResponse, error) {
	resp := &UnbindInputBucketResponse{}
	err := client.Request("mts", "UnbindInputBucket", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnbindInputBucketResponse struct {
	RequestId goaliyun.String
}

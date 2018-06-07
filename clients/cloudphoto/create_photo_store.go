package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type CreatePhotoStoreRequest struct {
	BucketName   string `position:"Query" name:"BucketName"`
	StoreName    string `position:"Query" name:"StoreName"`
	Remark       string `position:"Query" name:"Remark"`
	DefaultQuota int64  `position:"Query" name:"DefaultQuota"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CreatePhotoStoreRequest) Invoke(client goaliyun.Client) (*CreatePhotoStoreResponse, error) {
	resp := &CreatePhotoStoreResponse{}
	err := client.Request("cloudphoto", "CreatePhotoStore", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePhotoStoreResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

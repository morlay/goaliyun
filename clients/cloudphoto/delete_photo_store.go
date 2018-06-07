package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type DeletePhotoStoreRequest struct {
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeletePhotoStoreRequest) Invoke(client goaliyun.Client) (*DeletePhotoStoreResponse, error) {
	resp := &DeletePhotoStoreResponse{}
	err := client.Request("cloudphoto", "DeletePhotoStore", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePhotoStoreResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type RenameFaceRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	FaceName  string `position:"Query" name:"FaceName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *RenameFaceRequest) Invoke(client goaliyun.Client) (*RenameFaceResponse, error) {
	resp := &RenameFaceResponse{}
	err := client.Request("cloudphoto", "RenameFace", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenameFaceResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

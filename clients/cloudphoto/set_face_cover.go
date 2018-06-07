package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type SetFaceCoverRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *SetFaceCoverRequest) Invoke(client goaliyun.Client) (*SetFaceCoverResponse, error) {
	resp := &SetFaceCoverResponse{}
	err := client.Request("cloudphoto", "SetFaceCover", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetFaceCoverResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

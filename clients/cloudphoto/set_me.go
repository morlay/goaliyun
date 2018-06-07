package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type SetMeRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *SetMeRequest) Invoke(client goaliyun.Client) (*SetMeResponse, error) {
	resp := &SetMeResponse{}
	err := client.Request("cloudphoto", "SetMe", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetMeResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type DeleteEventRequest struct {
	EventId   int64  `position:"Query" name:"EventId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteEventRequest) Invoke(client goaliyun.Client) (*DeleteEventResponse, error) {
	resp := &DeleteEventResponse{}
	err := client.Request("cloudphoto", "DeleteEvent", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteEventResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteServiceRequest struct {
	ServiceName string `position:"Query" name:"ServiceName"`
	ServiceId   int64  `position:"Query" name:"ServiceId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DeleteServiceRequest) Invoke(client goaliyun.Client) (*DeleteServiceResponse, error) {
	resp := &DeleteServiceResponse{}
	err := client.Request("csb", "DeleteService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}

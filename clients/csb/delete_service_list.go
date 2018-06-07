package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteServiceListRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteServiceListRequest) Invoke(client goaliyun.Client) (*DeleteServiceListResponse, error) {
	resp := &DeleteServiceListResponse{}
	err := client.Request("csb", "DeleteServiceList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteServiceListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}

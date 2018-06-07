package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteProjectListRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteProjectListRequest) Invoke(client goaliyun.Client) (*DeleteProjectListResponse, error) {
	resp := &DeleteProjectListResponse{}
	err := client.Request("csb", "DeleteProjectList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteProjectListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}

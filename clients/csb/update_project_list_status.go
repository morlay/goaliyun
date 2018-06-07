package csb

import (
	"github.com/morlay/goaliyun"
)

type UpdateProjectListStatusRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateProjectListStatusRequest) Invoke(client goaliyun.Client) (*UpdateProjectListStatusResponse, error) {
	resp := &UpdateProjectListStatusResponse{}
	err := client.Request("csb", "UpdateProjectListStatus", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateProjectListStatusResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}

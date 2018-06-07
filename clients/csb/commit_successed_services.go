package csb

import (
	"github.com/morlay/goaliyun"
)

type CommitSuccessedServicesRequest struct {
	CsbName  string `position:"Query" name:"CsbName"`
	Services string `position:"Body" name:"Services"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CommitSuccessedServicesRequest) Invoke(client goaliyun.Client) (*CommitSuccessedServicesResponse, error) {
	resp := &CommitSuccessedServicesResponse{}
	err := client.Request("csb", "CommitSuccessedServices", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CommitSuccessedServicesResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}

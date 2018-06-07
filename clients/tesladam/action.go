package tesladam

import (
	"github.com/morlay/goaliyun"
)

type ActionRequest struct {
	OrderId  int64  `position:"Query" name:"OrderId"`
	StepCode string `position:"Query" name:"StepCode"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ActionRequest) Invoke(client goaliyun.Client) (*ActionResponse, error) {
	resp := &ActionResponse{}
	err := client.Request("tesladam", "Action", "2018-01-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActionResponse struct {
	Status  bool
	Message goaliyun.String
	Result  goaliyun.String
}

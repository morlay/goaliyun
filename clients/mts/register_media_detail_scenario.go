package mts

import (
	"github.com/morlay/goaliyun"
)

type RegisterMediaDetailScenarioRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Scenario             string `position:"Query" name:"Scenario"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RegisterMediaDetailScenarioRequest) Invoke(client goaliyun.Client) (*RegisterMediaDetailScenarioResponse, error) {
	resp := &RegisterMediaDetailScenarioResponse{}
	err := client.Request("mts", "RegisterMediaDetailScenario", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RegisterMediaDetailScenarioResponse struct {
	RequestId  goaliyun.String
	ScenarioId goaliyun.String
}

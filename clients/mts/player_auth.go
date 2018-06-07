package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type PlayerAuthRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *PlayerAuthRequest) Invoke(client goaliyun.Client) (*PlayerAuthResponse, error) {
	resp := &PlayerAuthResponse{}
	err := client.Request("mts", "PlayerAuth", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PlayerAuthResponse struct {
	RequestId  goaliyun.String
	LogURL     goaliyun.String
	SwitchList PlayerAuth_SwitchList
}

type PlayerAuth_Switch struct {
	State        goaliyun.String
	FunctionId   goaliyun.String
	SwitchId     goaliyun.String
	FunctionName goaliyun.String
}

type PlayerAuth_SwitchList []PlayerAuth_Switch

func (list *PlayerAuth_SwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PlayerAuth_Switch)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

package dyvmsapi

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type IvrCallRequest struct {
	ByeCode              string                 `position:"Query" name:"ByeCode"`
	MenuKeyMaps          *IvrCallMenuKeyMapList `position:"Query" type:"Repeated" name:"MenuKeyMap"`
	ResourceOwnerId      int64                  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                 `position:"Query" name:"ResourceOwnerAccount"`
	StartTtsParams       string                 `position:"Query" name:"StartTtsParams"`
	PlayTimes            int64                  `position:"Query" name:"PlayTimes"`
	OwnerId              int64                  `position:"Query" name:"OwnerId"`
	Timeout              int64                  `position:"Query" name:"Timeout"`
	StartCode            string                 `position:"Query" name:"StartCode"`
	CalledNumber         string                 `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string                 `position:"Query" name:"CalledShowNumber"`
	OutId                string                 `position:"Query" name:"OutId"`
	ByeTtsParams         string                 `position:"Query" name:"ByeTtsParams"`
	RegionId             string                 `position:"Query" name:"RegionId"`
}

func (req *IvrCallRequest) Invoke(client goaliyun.Client) (*IvrCallResponse, error) {
	resp := &IvrCallResponse{}
	err := client.Request("dyvmsapi", "IvrCall", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type IvrCallMenuKeyMap struct {
	Key       string `name:"Key"`
	Code      string `name:"Code"`
	TtsParams string `name:"TtsParams"`
}

type IvrCallResponse struct {
	RequestId goaliyun.String
	CallId    goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}

type IvrCallMenuKeyMapList []IvrCallMenuKeyMap

func (list *IvrCallMenuKeyMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IvrCallMenuKeyMap)
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

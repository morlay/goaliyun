package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveBoxCodeRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Code        string `position:"Query" name:"Code"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *RemoveBoxCodeRequest) Invoke(client goaliyun.Client) (*RemoveBoxCodeResponse, error) {
	resp := &RemoveBoxCodeResponse{}
	err := client.Request("itaas", "RemoveBoxCode", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveBoxCodeResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList RemoveBoxCodeErrorMessageList
}

type RemoveBoxCodeErrorMessage struct {
	ErrorMessage goaliyun.String
}

type RemoveBoxCodeErrorMessageList []RemoveBoxCodeErrorMessage

func (list *RemoveBoxCodeErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveBoxCodeErrorMessage)
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

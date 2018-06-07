package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SetWelcomePageURIRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Welcomeuri  string `position:"Query" name:"Welcomeuri"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SetWelcomePageURIRequest) Invoke(client goaliyun.Client) (*SetWelcomePageURIResponse, error) {
	resp := &SetWelcomePageURIResponse{}
	err := client.Request("itaas", "SetWelcomePageURI", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetWelcomePageURIResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList SetWelcomePageURIErrorMessageList
}

type SetWelcomePageURIErrorMessage struct {
	ErrorMessage goaliyun.String
}

type SetWelcomePageURIErrorMessageList []SetWelcomePageURIErrorMessage

func (list *SetWelcomePageURIErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetWelcomePageURIErrorMessage)
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

package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetWelcomePageURIRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetWelcomePageURIRequest) Invoke(client goaliyun.Client) (*GetWelcomePageURIResponse, error) {
	resp := &GetWelcomePageURIResponse{}
	err := client.Request("itaas", "GetWelcomePageURI", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetWelcomePageURIResponse struct {
	RequestId goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList GetWelcomePageURIErrorMessageList
}

type GetWelcomePageURIErrorMessage struct {
	ErrorMessage goaliyun.String
}

type GetWelcomePageURIErrorMessageList []GetWelcomePageURIErrorMessage

func (list *GetWelcomePageURIErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetWelcomePageURIErrorMessage)
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

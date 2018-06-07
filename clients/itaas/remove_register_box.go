package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveRegisterBoxRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Drsessionid string `position:"Query" name:"Drsessionid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *RemoveRegisterBoxRequest) Invoke(client goaliyun.Client) (*RemoveRegisterBoxResponse, error) {
	resp := &RemoveRegisterBoxResponse{}
	err := client.Request("itaas", "RemoveRegisterBox", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveRegisterBoxResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList RemoveRegisterBoxErrorMessageList
}

type RemoveRegisterBoxErrorMessage struct {
	ErrorMessage goaliyun.String
}

type RemoveRegisterBoxErrorMessageList []RemoveRegisterBoxErrorMessage

func (list *RemoveRegisterBoxErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveRegisterBoxErrorMessage)
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

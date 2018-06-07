package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateBoxCodeRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CreateBoxCodeRequest) Invoke(client goaliyun.Client) (*CreateBoxCodeResponse, error) {
	resp := &CreateBoxCodeResponse{}
	err := client.Request("itaas", "CreateBoxCode", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBoxCodeResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList CreateBoxCodeErrorMessageList
	Data      CreateBoxCodeData
}

type CreateBoxCodeErrorMessage struct {
	ErrorMessage goaliyun.String
}

type CreateBoxCodeData struct {
	ClientAppid       goaliyun.String
	Code              goaliyun.String
	CurrentTimeMillis goaliyun.Integer
	ExpiresIn         goaliyun.Integer
	ExpiresInUnit     goaliyun.String
}

type CreateBoxCodeErrorMessageList []CreateBoxCodeErrorMessage

func (list *CreateBoxCodeErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateBoxCodeErrorMessage)
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

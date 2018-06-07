package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetRegisterBoxNumberRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetRegisterBoxNumberRequest) Invoke(client goaliyun.Client) (*GetRegisterBoxNumberResponse, error) {
	resp := &GetRegisterBoxNumberResponse{}
	err := client.Request("itaas", "GetRegisterBoxNumber", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetRegisterBoxNumberResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList GetRegisterBoxNumberErrorMessageList
	Data      GetRegisterBoxNumberData
}

type GetRegisterBoxNumberErrorMessage struct {
	ErrorMessage goaliyun.String
}

type GetRegisterBoxNumberData struct {
	ActivedNumber goaliyun.Integer
	BuyNumber     goaliyun.Integer
	BoxesList     GetRegisterBoxNumberBoxInfoList
}

type GetRegisterBoxNumberBoxInfo struct {
	CurVersion      goaliyun.String
	DrName          goaliyun.String
	DrSessionId     goaliyun.String
	DrStatus        goaliyun.String
	DrStatusTxt     goaliyun.String
	Ipaddress       goaliyun.String
	LastReportTimeL goaliyun.Integer
	OnlineTimeL     goaliyun.Integer
	Screencode      goaliyun.String
	SysVersion      goaliyun.String
}

type GetRegisterBoxNumberErrorMessageList []GetRegisterBoxNumberErrorMessage

func (list *GetRegisterBoxNumberErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxNumberErrorMessage)
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

type GetRegisterBoxNumberBoxInfoList []GetRegisterBoxNumberBoxInfo

func (list *GetRegisterBoxNumberBoxInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxNumberBoxInfo)
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

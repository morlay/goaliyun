package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetRegisterBoxListRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetRegisterBoxListRequest) Invoke(client goaliyun.Client) (*GetRegisterBoxListResponse, error) {
	resp := &GetRegisterBoxListResponse{}
	err := client.Request("itaas", "GetRegisterBoxList", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetRegisterBoxListResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList GetRegisterBoxListErrorMessageList
	Data      GetRegisterBoxListData
}

type GetRegisterBoxListErrorMessage struct {
	ErrorMessage goaliyun.String
}

type GetRegisterBoxListData struct {
	ActivedNumber goaliyun.Integer
	BuyNumber     goaliyun.Integer
	BoxesList     GetRegisterBoxListBoxInfoList
}

type GetRegisterBoxListBoxInfo struct {
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

type GetRegisterBoxListErrorMessageList []GetRegisterBoxListErrorMessage

func (list *GetRegisterBoxListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxListErrorMessage)
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

type GetRegisterBoxListBoxInfoList []GetRegisterBoxListBoxInfo

func (list *GetRegisterBoxListBoxInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxListBoxInfo)
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

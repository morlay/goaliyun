package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetRegisterHistoryListRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetRegisterHistoryListRequest) Invoke(client goaliyun.Client) (*GetRegisterHistoryListResponse, error) {
	resp := &GetRegisterHistoryListResponse{}
	err := client.Request("itaas", "GetRegisterHistoryList", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetRegisterHistoryListResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	Data      GetRegisterHistoryListRegisterHistoryInfoList
	ErrorList GetRegisterHistoryListErrorMessageList
}

type GetRegisterHistoryListRegisterHistoryInfo struct {
	CreateTimeL  goaliyun.Integer
	DrIp         goaliyun.String
	DrMac        goaliyun.String
	DrName       goaliyun.String
	Eventinfo    goaliyun.String
	Eventtype    goaliyun.Integer
	EventtypeTxt goaliyun.String
	Memo         goaliyun.String
	Screencode   goaliyun.String
}

type GetRegisterHistoryListErrorMessage struct {
	ErrorMessage goaliyun.String
}

type GetRegisterHistoryListRegisterHistoryInfoList []GetRegisterHistoryListRegisterHistoryInfo

func (list *GetRegisterHistoryListRegisterHistoryInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterHistoryListRegisterHistoryInfo)
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

type GetRegisterHistoryListErrorMessageList []GetRegisterHistoryListErrorMessage

func (list *GetRegisterHistoryListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterHistoryListErrorMessage)
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

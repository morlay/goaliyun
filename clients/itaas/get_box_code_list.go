package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetBoxCodeListRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetBoxCodeListRequest) Invoke(client goaliyun.Client) (*GetBoxCodeListResponse, error) {
	resp := &GetBoxCodeListResponse{}
	err := client.Request("itaas", "GetBoxCodeList", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetBoxCodeListResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	Data      GetBoxCodeListBoxCodeInfoList
	ErrorList GetBoxCodeListErrorMessageList
}

type GetBoxCodeListBoxCodeInfo struct {
	BeginTime  goaliyun.Integer
	BoxInfo    goaliyun.String
	Code       goaliyun.String
	EndTime    goaliyun.Integer
	ModifyTime goaliyun.Integer
	Operator   goaliyun.String
	Screencode goaliyun.String
	Status     goaliyun.Integer
	StatusTxt  goaliyun.String
}

type GetBoxCodeListErrorMessage struct {
	ErrorMessage goaliyun.String
}

type GetBoxCodeListBoxCodeInfoList []GetBoxCodeListBoxCodeInfo

func (list *GetBoxCodeListBoxCodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBoxCodeListBoxCodeInfo)
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

type GetBoxCodeListErrorMessageList []GetBoxCodeListErrorMessage

func (list *GetBoxCodeListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBoxCodeListErrorMessage)
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

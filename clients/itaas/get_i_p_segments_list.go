package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetIPSegmentsListRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetIPSegmentsListRequest) Invoke(client goaliyun.Client) (*GetIPSegmentsListResponse, error) {
	resp := &GetIPSegmentsListResponse{}
	err := client.Request("itaas", "GetIPSegmentsList", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetIPSegmentsListResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	Data      GetIPSegmentsListIpsegmentInfoList
	ErrorList GetIPSegmentsListErrorMessageList
}

type GetIPSegmentsListIpsegmentInfo struct {
	CreateTimeL goaliyun.Integer
	Ipsegment   goaliyun.String
	Memo        goaliyun.String
	ModifyTimeL goaliyun.Integer
	Uuid        goaliyun.String
}

type GetIPSegmentsListErrorMessage struct {
	ErrorMessage goaliyun.String
}

type GetIPSegmentsListIpsegmentInfoList []GetIPSegmentsListIpsegmentInfo

func (list *GetIPSegmentsListIpsegmentInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetIPSegmentsListIpsegmentInfo)
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

type GetIPSegmentsListErrorMessageList []GetIPSegmentsListErrorMessage

func (list *GetIPSegmentsListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetIPSegmentsListErrorMessage)
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

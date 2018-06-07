package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindApprovalOrderListRequest struct {
	ProjectName string `position:"Query" name:"ProjectName"`
	Alias       string `position:"Query" name:"Alias"`
	ServiceName string `position:"Query" name:"ServiceName"`
	ServiceId   int64  `position:"Query" name:"ServiceId"`
	PageNum     int64  `position:"Query" name:"PageNum"`
	OnlyPending string `position:"Query" name:"OnlyPending"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *FindApprovalOrderListRequest) Invoke(client goaliyun.Client) (*FindApprovalOrderListResponse, error) {
	resp := &FindApprovalOrderListResponse{}
	err := client.Request("csb", "FindApprovalOrderList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindApprovalOrderListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindApprovalOrderListData
}

type FindApprovalOrderListData struct {
	CurrentPage goaliyun.Integer
	PageNumber  goaliyun.Integer
	OrderList   FindApprovalOrderListOrderList
}

type FindApprovalOrderListOrder struct {
	Alias               goaliyun.String
	CredentialGroupId   goaliyun.Integer
	CsbId               goaliyun.Integer
	GmtCreate           goaliyun.Integer
	GmtModified         goaliyun.Integer
	GroupName           goaliyun.String
	Id                  goaliyun.Integer
	ProjectName         goaliyun.String
	ServiceId           goaliyun.Integer
	ServiceName         goaliyun.String
	ServiceStatus       goaliyun.Integer
	ServiceVersion      goaliyun.String
	StatisticName       goaliyun.String
	Status              goaliyun.Integer
	StrictWhiteListJson goaliyun.String
	UserId              goaliyun.String
	UserName            goaliyun.String
	SlaInfo             FindApprovalOrderListSlaInfo
	Total               FindApprovalOrderListTotal
}

type FindApprovalOrderListSlaInfo struct {
	Qph goaliyun.Integer
	Qps goaliyun.Integer
}

type FindApprovalOrderListTotal struct {
	ErrorNum goaliyun.Integer
	Total    goaliyun.Integer
}

type FindApprovalOrderListOrderList []FindApprovalOrderListOrder

func (list *FindApprovalOrderListOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindApprovalOrderListOrder)
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

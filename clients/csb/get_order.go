package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetOrderRequest struct {
	OrderId     int64  `position:"Query" name:"OrderId"`
	ServiceName string `position:"Query" name:"ServiceName"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetOrderRequest) Invoke(client goaliyun.Client) (*GetOrderResponse, error) {
	resp := &GetOrderResponse{}
	err := client.Request("csb", "GetOrder", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOrderResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetOrderData
}

type GetOrderData struct {
	Order GetOrderOrder
}

type GetOrderOrder struct {
	Alias                 goaliyun.String
	CredentialGroupId     goaliyun.Integer
	CsbId                 goaliyun.Integer
	DauthGroupName        goaliyun.String
	GmtCreate             goaliyun.Integer
	GmtModified           goaliyun.Integer
	GroupName             goaliyun.String
	Id                    goaliyun.Integer
	ProjectName           goaliyun.String
	ServiceId             goaliyun.Integer
	ServiceName           goaliyun.String
	ServiceStatus         goaliyun.Integer
	ServiceVersion        goaliyun.String
	StatisticName         goaliyun.String
	Status                goaliyun.Integer
	StrictWhiteListJson   goaliyun.String
	UserId                goaliyun.String
	ErrorTypeCatagoryList GetOrderErrorTypeCatagoryList
	StrictWhiteList       GetOrderStrictWhiteListList
	Service               GetOrderService
	SlaInfo               GetOrderSlaInfo
	Total                 GetOrderTotal
}

type GetOrderErrorTypeCatagory struct {
	Total    goaliyun.Integer
	ErrorNum goaliyun.Integer
	Name     goaliyun.String
}

type GetOrderService struct {
	AccessParamsJSON    goaliyun.String
	Active              bool
	Alias               goaliyun.String
	AllVisiable         bool
	ConsumeTypesJSON    goaliyun.String
	CreateTime          goaliyun.Integer
	CsbId               goaliyun.Integer
	ErrDefJSON          goaliyun.String
	Id                  goaliyun.Integer
	InterfaceName       goaliyun.String
	OldVersion          goaliyun.String
	OttFlag             bool
	OwnerId             goaliyun.String
	PrincipalName       goaliyun.String
	ProjectId           goaliyun.String
	ProjectName         goaliyun.String
	ProvideType         goaliyun.String
	SSL                 bool
	Scope               goaliyun.String
	ServiceName         goaliyun.String
	ServiceProviderType goaliyun.String
	ServiceVersion      goaliyun.String
	SkipAuth            bool
	StatisticName       goaliyun.String
	Status              goaliyun.Integer
	UserId              goaliyun.Integer
	ValidConsumeTypes   bool
	ValidProvideType    bool
}

type GetOrderSlaInfo struct {
	Qph goaliyun.String
	Qps goaliyun.String
}

type GetOrderTotal struct {
	ErrorNum goaliyun.Integer
	Total    goaliyun.Integer
}

type GetOrderErrorTypeCatagoryList []GetOrderErrorTypeCatagory

func (list *GetOrderErrorTypeCatagoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetOrderErrorTypeCatagory)
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

type GetOrderStrictWhiteListList []goaliyun.String

func (list *GetOrderStrictWhiteListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

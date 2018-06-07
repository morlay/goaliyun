package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDomainListRequest struct {
	EndExpirationDate     int64  `position:"Query" name:"EndExpirationDate"`
	ProductDomainType     string `position:"Query" name:"ProductDomainType"`
	OrderKeyType          string `position:"Query" name:"OrderKeyType"`
	DomainName            string `position:"Query" name:"DomainName"`
	StartExpirationDate   int64  `position:"Query" name:"StartExpirationDate"`
	PageNum               int64  `position:"Query" name:"PageNum"`
	OrderByType           string `position:"Query" name:"OrderByType"`
	DomainGroupId         string `position:"Query" name:"DomainGroupId"`
	EndRegistrationDate   int64  `position:"Query" name:"EndRegistrationDate"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	PageSize              int64  `position:"Query" name:"PageSize"`
	Lang                  string `position:"Query" name:"Lang"`
	QueryType             string `position:"Query" name:"QueryType"`
	StartRegistrationDate int64  `position:"Query" name:"StartRegistrationDate"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *QueryDomainListRequest) Invoke(client goaliyun.Client) (*QueryDomainListResponse, error) {
	resp := &QueryDomainListResponse{}
	err := client.Request("domain", "QueryDomainList", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDomainListResponse struct {
	RequestId      goaliyun.String
	TotalItemNum   goaliyun.Integer
	CurrentPageNum goaliyun.Integer
	TotalPageNum   goaliyun.Integer
	PageSize       goaliyun.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryDomainListDomainList
}

type QueryDomainListDomain struct {
	DomainName             goaliyun.String
	InstanceId             goaliyun.String
	ExpirationDate         goaliyun.String
	RegistrationDate       goaliyun.String
	DomainType             goaliyun.String
	DomainStatus           goaliyun.String
	ProductId              goaliyun.String
	ExpirationDateLong     goaliyun.Integer
	RegistrationDateLong   goaliyun.Integer
	Premium                bool
	DomainAuditStatus      goaliyun.String
	ExpirationDateStatus   goaliyun.String
	RegistrantType         goaliyun.String
	DomainGroupId          goaliyun.String
	Remark                 goaliyun.String
	DomainGroupName        goaliyun.String
	ExpirationCurrDateDiff goaliyun.Integer
}

type QueryDomainListDomainList []QueryDomainListDomain

func (list *QueryDomainListDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDomainListDomain)
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

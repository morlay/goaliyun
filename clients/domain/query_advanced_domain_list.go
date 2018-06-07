package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAdvancedDomainListRequest struct {
	ProductDomainType     string `position:"Query" name:"ProductDomainType"`
	PageNum               int64  `position:"Query" name:"PageNum"`
	Excluded              string `position:"Query" name:"Excluded"`
	StartLength           int64  `position:"Query" name:"StartLength"`
	ExcludedSuffix        string `position:"Query" name:"ExcludedSuffix"`
	PageSize              int64  `position:"Query" name:"PageSize"`
	Lang                  string `position:"Query" name:"Lang"`
	ExcludedPrefix        string `position:"Query" name:"ExcludedPrefix"`
	KeyWord               string `position:"Query" name:"KeyWord"`
	ProductDomainTypeSort string `position:"Query" name:"ProductDomainTypeSort"`
	EndExpirationDate     int64  `position:"Query" name:"EndExpirationDate"`
	Suffixs               string `position:"Query" name:"Suffixs"`
	DomainNameSort        string `position:"Query" name:"DomainNameSort"`
	ExpirationDateSort    string `position:"Query" name:"ExpirationDateSort"`
	StartExpirationDate   int64  `position:"Query" name:"StartExpirationDate"`
	DomainStatus          int64  `position:"Query" name:"DomainStatus"`
	DomainGroupId         int64  `position:"Query" name:"DomainGroupId"`
	KeyWordSuffix         string `position:"Query" name:"KeyWordSuffix"`
	KeyWordPrefix         string `position:"Query" name:"KeyWordPrefix"`
	TradeType             int64  `position:"Query" name:"TradeType"`
	EndRegistrationDate   int64  `position:"Query" name:"EndRegistrationDate"`
	Form                  int64  `position:"Query" name:"Form"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	RegistrationDateSort  string `position:"Query" name:"RegistrationDateSort"`
	StartRegistrationDate int64  `position:"Query" name:"StartRegistrationDate"`
	EndLength             int64  `position:"Query" name:"EndLength"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *QueryAdvancedDomainListRequest) Invoke(client goaliyun.Client) (*QueryAdvancedDomainListResponse, error) {
	resp := &QueryAdvancedDomainListResponse{}
	err := client.Request("domain", "QueryAdvancedDomainList", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAdvancedDomainListResponse struct {
	RequestId      goaliyun.String
	TotalItemNum   goaliyun.Integer
	CurrentPageNum goaliyun.Integer
	TotalPageNum   goaliyun.Integer
	PageSize       goaliyun.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryAdvancedDomainListDomainList
}

type QueryAdvancedDomainListDomain struct {
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

type QueryAdvancedDomainListDomainList []QueryAdvancedDomainListDomain

func (list *QueryAdvancedDomainListDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAdvancedDomainListDomain)
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

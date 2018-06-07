package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTransferInListRequest struct {
	SubmissionStartDate    int64  `position:"Query" name:"SubmissionStartDate"`
	UserClientIp           string `position:"Query" name:"UserClientIp"`
	SubmissionEndDate      int64  `position:"Query" name:"SubmissionEndDate"`
	DomainName             string `position:"Query" name:"DomainName"`
	SimpleTransferInStatus string `position:"Query" name:"SimpleTransferInStatus"`
	PageSize               int64  `position:"Query" name:"PageSize"`
	Lang                   string `position:"Query" name:"Lang"`
	PageNum                int64  `position:"Query" name:"PageNum"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *QueryTransferInListRequest) Invoke(client goaliyun.Client) (*QueryTransferInListResponse, error) {
	resp := &QueryTransferInListResponse{}
	err := client.Request("domain-intl", "QueryTransferInList", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTransferInListResponse struct {
	RequestId      goaliyun.String
	TotalItemNum   goaliyun.Integer
	CurrentPageNum goaliyun.Integer
	TotalPageNum   goaliyun.Integer
	PageSize       goaliyun.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryTransferInListTransferInInfoList
}

type QueryTransferInListTransferInInfo struct {
	SubmissionDate                              goaliyun.String
	ModificationDate                            goaliyun.String
	UserId                                      goaliyun.String
	InstanceId                                  goaliyun.String
	DomainName                                  goaliyun.String
	Status                                      goaliyun.Integer
	SimpleTransferInStatus                      goaliyun.String
	ResultCode                                  goaliyun.String
	ResultDate                                  goaliyun.String
	ResultMsg                                   goaliyun.String
	TransferAuthorizationCodeSubmissionDate     goaliyun.String
	NeedMailCheck                               bool
	Email                                       goaliyun.String
	WhoisMailStatus                             bool
	ExpirationDate                              goaliyun.String
	ProgressBarType                             goaliyun.Integer
	SubmissionDateLong                          goaliyun.Integer
	ModificationDateLong                        goaliyun.Integer
	ResultDateLong                              goaliyun.Integer
	ExpirationDateLong                          goaliyun.Integer
	TransferAuthorizationCodeSubmissionDateLong goaliyun.Integer
}

type QueryTransferInListTransferInInfoList []QueryTransferInListTransferInInfo

func (list *QueryTransferInListTransferInInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTransferInListTransferInInfo)
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

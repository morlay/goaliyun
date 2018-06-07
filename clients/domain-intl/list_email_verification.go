package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListEmailVerificationRequest struct {
	BeginCreateTime    int64  `position:"Query" name:"BeginCreateTime"`
	EndCreateTime      int64  `position:"Query" name:"EndCreateTime"`
	PageSize           int64  `position:"Query" name:"PageSize"`
	Lang               string `position:"Query" name:"Lang"`
	PageNum            int64  `position:"Query" name:"PageNum"`
	Email              string `position:"Query" name:"Email"`
	VerificationStatus int64  `position:"Query" name:"VerificationStatus"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *ListEmailVerificationRequest) Invoke(client goaliyun.Client) (*ListEmailVerificationResponse, error) {
	resp := &ListEmailVerificationResponse{}
	err := client.Request("domain-intl", "ListEmailVerification", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListEmailVerificationResponse struct {
	RequestId      goaliyun.String
	TotalItemNum   goaliyun.Integer
	CurrentPageNum goaliyun.Integer
	TotalPageNum   goaliyun.Integer
	PageSize       goaliyun.Integer
	PrePage        bool
	NextPage       bool
	Data           ListEmailVerificationEmailVerificationList
}

type ListEmailVerificationEmailVerification struct {
	GmtCreate           goaliyun.String
	GmtModified         goaliyun.String
	Email               goaliyun.String
	UserId              goaliyun.String
	EmailVerificationNo goaliyun.String
	TokenSendTime       goaliyun.String
	VerificationStatus  goaliyun.Integer
	VerificationTime    goaliyun.String
	SendIp              goaliyun.String
	ConfirmIp           goaliyun.String
}

type ListEmailVerificationEmailVerificationList []ListEmailVerificationEmailVerification

func (list *ListEmailVerificationEmailVerificationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEmailVerificationEmailVerification)
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

package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindCredentialsListRequest struct {
	CsbId     int64  `position:"Query" name:"CsbId"`
	PageNum   int64  `position:"Query" name:"PageNum"`
	GroupName string `position:"Query" name:"GroupName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *FindCredentialsListRequest) Invoke(client goaliyun.Client) (*FindCredentialsListResponse, error) {
	resp := &FindCredentialsListResponse{}
	err := client.Request("csb", "FindCredentialsList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindCredentialsListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindCredentialsListData
}

type FindCredentialsListData struct {
	CurrentPage    goaliyun.Integer
	PageNumber     goaliyun.Integer
	CredentialList FindCredentialsListCredentialList
}

type FindCredentialsListCredential struct {
	GmtCreate         goaliyun.Integer
	Id                goaliyun.Integer
	Name              goaliyun.String
	OwnerAttr         goaliyun.String
	UserId            goaliyun.String
	CurrentCredential FindCredentialsListCurrentCredential
	NewCredential     FindCredentialsListNewCredential
}

type FindCredentialsListCurrentCredential struct {
	AccessKey goaliyun.String
	SecretKey goaliyun.String
}

type FindCredentialsListNewCredential struct {
	AccessKey goaliyun.String
	SecretKey goaliyun.String
}

type FindCredentialsListCredentialList []FindCredentialsListCredential

func (list *FindCredentialsListCredentialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindCredentialsListCredential)
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

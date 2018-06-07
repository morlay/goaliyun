package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type QueryTransferInByInstanceIdRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *QueryTransferInByInstanceIdRequest) Invoke(client goaliyun.Client) (*QueryTransferInByInstanceIdResponse, error) {
	resp := &QueryTransferInByInstanceIdResponse{}
	err := client.Request("domain-intl", "QueryTransferInByInstanceId", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTransferInByInstanceIdResponse struct {
	RequestId                                   goaliyun.String
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

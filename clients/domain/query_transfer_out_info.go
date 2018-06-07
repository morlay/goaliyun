package domain

import (
	"github.com/morlay/goaliyun"
)

type QueryTransferOutInfoRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *QueryTransferOutInfoRequest) Invoke(client goaliyun.Client) (*QueryTransferOutInfoResponse, error) {
	resp := &QueryTransferOutInfoResponse{}
	err := client.Request("domain", "QueryTransferOutInfo", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTransferOutInfoResponse struct {
	RequestId                         goaliyun.String
	Status                            goaliyun.Integer
	Email                             goaliyun.String
	TransferAuthorizationCodeSendDate goaliyun.String
	ExpirationDate                    goaliyun.String
	PendingRequestDate                goaliyun.String
	ResultCode                        goaliyun.String
	ResultMsg                         goaliyun.String
}

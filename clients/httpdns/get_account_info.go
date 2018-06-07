package httpdns

import (
	"github.com/morlay/goaliyun"
)

type GetAccountInfoRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetAccountInfoRequest) Invoke(client goaliyun.Client) (*GetAccountInfoResponse, error) {
	resp := &GetAccountInfoResponse{}
	err := client.Request("httpdns", "GetAccountInfo", "2016-02-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAccountInfoResponse struct {
	RequestId   goaliyun.String
	AccountInfo GetAccountInfoAccountInfo
}

type GetAccountInfoAccountInfo struct {
	AccountId              goaliyun.String
	MonthFreeCount         goaliyun.Integer
	MonthHttpsResolveCount goaliyun.Integer
	MonthResolveCount      goaliyun.Integer
	PackageCount           goaliyun.Integer
	UserStatus             goaliyun.Integer
	SignSecret             goaliyun.String
	UnsignedEnabled        bool
	SignedCount            goaliyun.Integer
	UnsignedCount          goaliyun.Integer
}

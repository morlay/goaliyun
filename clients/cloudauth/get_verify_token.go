package cloudauth

import (
	"github.com/morlay/goaliyun"
)

type GetVerifyTokenRequest struct {
	UserData        string `position:"Query" name:"UserData"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Binding         string `position:"Query" name:"Binding"`
	TicketId        string `position:"Query" name:"TicketId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetVerifyTokenRequest) Invoke(client goaliyun.Client) (*GetVerifyTokenResponse, error) {
	resp := &GetVerifyTokenResponse{}
	err := client.Request("cloudauth", "GetVerifyToken", "2018-05-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetVerifyTokenResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetVerifyTokenData
}

type GetVerifyTokenData struct {
	VerifyToken GetVerifyTokenVerifyToken
	StsToken    GetVerifyTokenStsToken
}

type GetVerifyTokenVerifyToken struct {
	Token           goaliyun.String
	DurationSeconds goaliyun.Integer
}

type GetVerifyTokenStsToken struct {
	AccessKeyId     goaliyun.String
	AccessKeySecret goaliyun.String
	Expiration      goaliyun.String
	EndPoint        goaliyun.String
	BucketName      goaliyun.String
	Path            goaliyun.String
	Token           goaliyun.String
}

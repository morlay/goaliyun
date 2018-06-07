package afs

import (
	"github.com/morlay/goaliyun"
)

type DescribeCaptchaDayRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
	Type            string `position:"Query" name:"Type"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeCaptchaDayRequest) Invoke(client goaliyun.Client) (*DescribeCaptchaDayResponse, error) {
	resp := &DescribeCaptchaDayResponse{}
	err := client.Request("afs", "DescribeCaptchaDay", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCaptchaDayResponse struct {
	RequestId  goaliyun.String
	BizCode    goaliyun.String
	HasData    bool
	CaptchaDay DescribeCaptchaDayCaptchaDay
}

type DescribeCaptchaDayCaptchaDay struct {
	Init                        goaliyun.Integer
	AskForVerify                goaliyun.Integer
	DirecetStrategyInterception goaliyun.Integer
	TwiceVerify                 goaliyun.Integer
	Pass                        goaliyun.Integer
	CheckTested                 goaliyun.Integer
	UncheckTested               goaliyun.Integer
	LegalSign                   goaliyun.Integer
	MaliciousFlow               goaliyun.Integer
}

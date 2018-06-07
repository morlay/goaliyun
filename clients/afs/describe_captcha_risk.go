package afs

import (
	"github.com/morlay/goaliyun"
)

type DescribeCaptchaRiskRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeCaptchaRiskRequest) Invoke(client goaliyun.Client) (*DescribeCaptchaRiskResponse, error) {
	resp := &DescribeCaptchaRiskResponse{}
	err := client.Request("afs", "DescribeCaptchaRisk", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCaptchaRiskResponse struct {
	RequestId      goaliyun.String
	BizCode        goaliyun.String
	NumOfThisMonth goaliyun.Integer
	NumOfLastMonth goaliyun.Integer
	RiskLevel      goaliyun.String
}

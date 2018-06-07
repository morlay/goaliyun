package afs

import (
	"github.com/morlay/goaliyun"
)

type AnalyzeNvcRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Data            string `position:"Query" name:"Data"`
	ScoreJsonStr    string `position:"Query" name:"ScoreJsonStr"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AnalyzeNvcRequest) Invoke(client goaliyun.Client) (*AnalyzeNvcResponse, error) {
	resp := &AnalyzeNvcResponse{}
	err := client.Request("afs", "AnalyzeNvc", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AnalyzeNvcResponse struct {
	RequestId goaliyun.String
	BizCode   goaliyun.String
}

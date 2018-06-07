package lubanruler

import (
	"github.com/morlay/goaliyun"
)

type GetResultByLubanrulerRequest struct {
	AonePipelineId int64  `position:"Query" name:"AonePipelineId"`
	Token          string `position:"Query" name:"Token"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *GetResultByLubanrulerRequest) Invoke(client goaliyun.Client) (*GetResultByLubanrulerResponse, error) {
	resp := &GetResultByLubanrulerResponse{}
	err := client.Request("lubanruler", "GetResultByLubanruler", "2017-12-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetResultByLubanrulerResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RunStatus goaliyun.String
	RequestId goaliyun.String
	Data      GetResultByLubanrulerData
}

type GetResultByLubanrulerData struct {
	AonePipelineId         goaliyun.Integer
	AppName                goaliyun.String
	ScmUrl                 goaliyun.String
	ScmBranch              goaliyun.String
	TaskStatus             goaliyun.String
	BlockerCnt             goaliyun.Integer
	CriticalCnt            goaliyun.Integer
	InfoCnt                goaliyun.Integer
	MajorCnt               goaliyun.Integer
	MinorCnt               goaliyun.Integer
	Complexity             goaliyun.String
	DuplicatedLinesDensity goaliyun.String
	ReliabilityRating      goaliyun.String
	SecurityRating         goaliyun.String
	SqaleRating            goaliyun.String
	LineOfCode             goaliyun.String
	Functions              goaliyun.String
	ResultUrl              goaliyun.String
	CommentLinesDensity    goaliyun.String
}

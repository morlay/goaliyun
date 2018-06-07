package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type WafLogRequest struct {
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *WafLogRequest) Invoke(client goaliyun.Client) (*WafLogResponse, error) {
	resp := &WafLogResponse{}
	err := client.Request("yundun", "WafLog", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type WafLogResponse struct {
	RequestId   goaliyun.String
	WebAttack   goaliyun.Integer
	NewWafUser  bool
	WafOpened   bool
	InWhiteList bool
	DomainCount goaliyun.Integer
	StartTime   goaliyun.String
	EndTime     goaliyun.String
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalCount  goaliyun.Integer
	LogList     WafLogWafLogList
}

type WafLogWafLog struct {
	SourceIp goaliyun.String
	Time     goaliyun.String
	Url      goaliyun.String
	Type     goaliyun.String
	Status   goaliyun.Integer
}

type WafLogWafLogList []WafLogWafLog

func (list *WafLogWafLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WafLogWafLog)
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

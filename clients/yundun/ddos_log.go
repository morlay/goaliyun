package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DdosLogRequest struct {
	JstOwnerId   int64  `position:"Query" name:"JstOwnerId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DdosLogRequest) Invoke(client goaliyun.Client) (*DdosLogResponse, error) {
	resp := &DdosLogResponse{}
	err := client.Request("yundun", "DdosLog", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DdosLogResponse struct {
	RequestId    goaliyun.String
	AttackStatus goaliyun.Integer
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	TotalCount   goaliyun.Integer
	LogList      DdosLogDdosLogList
}

type DdosLogDdosLog struct {
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	Reason       goaliyun.String
	Status       goaliyun.Integer
	Bps          goaliyun.Integer
	Pps          goaliyun.Integer
	Qps          goaliyun.Integer
	AttackType   goaliyun.String
	AttackIpList goaliyun.String
	Type         goaliyun.Integer
}

type DdosLogDdosLogList []DdosLogDdosLog

func (list *DdosLogDdosLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosLogDdosLog)
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

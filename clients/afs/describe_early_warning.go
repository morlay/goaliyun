package afs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeEarlyWarningRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeEarlyWarningRequest) Invoke(client goaliyun.Client) (*DescribeEarlyWarningResponse, error) {
	resp := &DescribeEarlyWarningResponse{}
	err := client.Request("afs", "DescribeEarlyWarning", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeEarlyWarningResponse struct {
	RequestId     goaliyun.String
	HasWarning    bool
	BizCode       goaliyun.String
	EarlyWarnings DescribeEarlyWarningEarlyWarningList
}

type DescribeEarlyWarningEarlyWarning struct {
	WarnOpen  bool
	Title     goaliyun.String
	Content   goaliyun.String
	Frequency goaliyun.String
	TimeOpen  bool
	TimeBegin goaliyun.String
	TimeEnd   goaliyun.String
	Channel   goaliyun.String
}

type DescribeEarlyWarningEarlyWarningList []DescribeEarlyWarningEarlyWarning

func (list *DescribeEarlyWarningEarlyWarningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEarlyWarningEarlyWarning)
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

package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeWarningRequest struct {
	TypeNames       string `position:"Query" name:"TypeNames"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RiskName        string `position:"Query" name:"RiskName"`
	StatusList      string `position:"Query" name:"StatusList"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RiskLevels      string `position:"Query" name:"RiskLevels"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	CurrentPage     int64  `position:"Query" name:"CurrentPage"`
	Dealed          string `position:"Query" name:"Dealed"`
	SubTypeNames    string `position:"Query" name:"SubTypeNames"`
	Uuids           string `position:"Query" name:"Uuids"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeWarningRequest) Invoke(client goaliyun.Client) (*DescribeWarningResponse, error) {
	resp := &DescribeWarningResponse{}
	err := client.Request("aegis", "DescribeWarning", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeWarningResponse struct {
	RequestId   goaliyun.String
	Count       goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalCount  goaliyun.Integer
	CurrentPage goaliyun.Integer
	Warnings    DescribeWarningWarningList
}

type DescribeWarningWarningList []goaliyun.String

func (list *DescribeWarningWarningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

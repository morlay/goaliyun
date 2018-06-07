package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeServiceHealthRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeServiceHealthRequest) Invoke(client goaliyun.Client) (*DescribeServiceHealthResponse, error) {
	resp := &DescribeServiceHealthResponse{}
	err := client.Request("emr", "DescribeServiceHealth", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeServiceHealthResponse struct {
	RequestId                 goaliyun.String
	Name                      goaliyun.String
	ComponentHealthResultList DescribeServiceHealthComponentHealthResultList
	HealthResult              DescribeServiceHealthHealthResult
}

type DescribeServiceHealthComponentHealthResult struct {
	Key           goaliyun.String
	PassNumber    goaliyun.Integer
	ErrorNumber   goaliyun.Integer
	WarningNumber goaliyun.Integer
	UnKnownNumber goaliyun.Integer
}

type DescribeServiceHealthHealthResult struct {
	Key           goaliyun.String
	PassNumber    goaliyun.Integer
	ErrorNumber   goaliyun.Integer
	WarningNumber goaliyun.Integer
	UnKnownNumber goaliyun.Integer
}

type DescribeServiceHealthComponentHealthResultList []DescribeServiceHealthComponentHealthResult

func (list *DescribeServiceHealthComponentHealthResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServiceHealthComponentHealthResult)
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

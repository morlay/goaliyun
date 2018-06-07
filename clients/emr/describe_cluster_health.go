package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterHealthRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterHealthRequest) Invoke(client goaliyun.Client) (*DescribeClusterHealthResponse, error) {
	resp := &DescribeClusterHealthResponse{}
	err := client.Request("emr", "DescribeClusterHealth", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterHealthResponse struct {
	RequestId             goaliyun.String
	ClusterId             goaliyun.Integer
	ServiceHealthInfoList DescribeClusterHealthServiceHealthInfoList
	HealthResult          DescribeClusterHealthHealthResult
}

type DescribeClusterHealthServiceHealthInfo struct {
	Key           goaliyun.String
	PassNumber    goaliyun.Integer
	ErrorNumber   goaliyun.Integer
	WarningNumber goaliyun.Integer
	UnKnownNumber goaliyun.Integer
}

type DescribeClusterHealthHealthResult struct {
	Key           goaliyun.String
	PassNumber    goaliyun.Integer
	ErrorNumber   goaliyun.Integer
	WarningNumber goaliyun.Integer
	UnKnownNumber goaliyun.Integer
}

type DescribeClusterHealthServiceHealthInfoList []DescribeClusterHealthServiceHealthInfo

func (list *DescribeClusterHealthServiceHealthInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterHealthServiceHealthInfo)
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

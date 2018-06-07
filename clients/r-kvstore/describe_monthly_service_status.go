package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMonthlyServiceStatusRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Month                string `position:"Query" name:"Month"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMonthlyServiceStatusRequest) Invoke(client goaliyun.Client) (*DescribeMonthlyServiceStatusResponse, error) {
	resp := &DescribeMonthlyServiceStatusResponse{}
	err := client.Request("r-kvstore", "DescribeMonthlyServiceStatus", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMonthlyServiceStatusResponse struct {
	RequestId        goaliyun.String
	TotalCount       goaliyun.Integer
	InstanceSLAInfos DescribeMonthlyServiceStatusInstanceSLAInfoList
}

type DescribeMonthlyServiceStatusInstanceSLAInfo struct {
	InstanceId goaliyun.String
	UptimePct  goaliyun.Float
}

type DescribeMonthlyServiceStatusInstanceSLAInfoList []DescribeMonthlyServiceStatusInstanceSLAInfo

func (list *DescribeMonthlyServiceStatusInstanceSLAInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonthlyServiceStatusInstanceSLAInfo)
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

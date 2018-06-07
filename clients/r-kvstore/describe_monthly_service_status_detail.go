package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMonthlyServiceStatusDetailRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	Month                string `position:"Query" name:"Month"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMonthlyServiceStatusDetailRequest) Invoke(client goaliyun.Client) (*DescribeMonthlyServiceStatusDetailResponse, error) {
	resp := &DescribeMonthlyServiceStatusDetailResponse{}
	err := client.Request("r-kvstore", "DescribeMonthlyServiceStatusDetail", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMonthlyServiceStatusDetailResponse struct {
	RequestId     goaliyun.String
	InstanceId    goaliyun.String
	UptimePct     goaliyun.Float
	AffectedInfos DescribeMonthlyServiceStatusDetailAffectedInfoList
}

type DescribeMonthlyServiceStatusDetailAffectedInfo struct {
	StartTime   goaliyun.String
	EndTime     goaliyun.String
	Description goaliyun.String
}

type DescribeMonthlyServiceStatusDetailAffectedInfoList []DescribeMonthlyServiceStatusDetailAffectedInfo

func (list *DescribeMonthlyServiceStatusDetailAffectedInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonthlyServiceStatusDetailAffectedInfo)
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

package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeEipMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeEipMonitorDataRequest) Invoke(client goaliyun.Client) (*DescribeEipMonitorDataResponse, error) {
	resp := &DescribeEipMonitorDataResponse{}
	err := client.Request("vpc", "DescribeEipMonitorData", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeEipMonitorDataResponse struct {
	RequestId       goaliyun.String
	EipMonitorDatas DescribeEipMonitorDataEipMonitorDataList
}

type DescribeEipMonitorDataEipMonitorData struct {
	EipRX        goaliyun.Integer
	EipTX        goaliyun.Integer
	EipFlow      goaliyun.Integer
	EipBandwidth goaliyun.Integer
	EipPackets   goaliyun.Integer
	TimeStamp    goaliyun.String
}

type DescribeEipMonitorDataEipMonitorDataList []DescribeEipMonitorDataEipMonitorData

func (list *DescribeEipMonitorDataEipMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipMonitorDataEipMonitorData)
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

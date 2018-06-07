package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBandwidthPackagePublicIpMonitorDataRequest struct {
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

func (req *DescribeBandwidthPackagePublicIpMonitorDataRequest) Invoke(client goaliyun.Client) (*DescribeBandwidthPackagePublicIpMonitorDataResponse, error) {
	resp := &DescribeBandwidthPackagePublicIpMonitorDataResponse{}
	err := client.Request("vpc", "DescribeBandwidthPackagePublicIpMonitorData", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBandwidthPackagePublicIpMonitorDataResponse struct {
	RequestId    goaliyun.String
	MonitorDatas DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList
}

type DescribeBandwidthPackagePublicIpMonitorDataMonitorData struct {
	RX                   goaliyun.Integer
	TX                   goaliyun.Integer
	ReceivedBandwidth    goaliyun.Integer
	TransportedBandwidth goaliyun.Integer
	Flow                 goaliyun.Integer
	Bandwidth            goaliyun.Integer
	Packets              goaliyun.Integer
	TimeStamp            goaliyun.String
}

type DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList []DescribeBandwidthPackagePublicIpMonitorDataMonitorData

func (list *DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthPackagePublicIpMonitorDataMonitorData)
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

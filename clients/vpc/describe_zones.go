package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeZonesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeZonesRequest) Invoke(client goaliyun.Client) (*DescribeZonesResponse, error) {
	resp := &DescribeZonesResponse{}
	err := client.Request("vpc", "DescribeZones", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeZonesResponse struct {
	RequestId goaliyun.String
	Zones     DescribeZonesZoneList
}

type DescribeZonesZone struct {
	ZoneId    goaliyun.String
	LocalName goaliyun.String
}

type DescribeZonesZoneList []DescribeZonesZone

func (list *DescribeZonesZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesZone)
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

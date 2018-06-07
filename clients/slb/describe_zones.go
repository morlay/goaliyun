package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeZonesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeZonesRequest) Invoke(client goaliyun.Client) (*DescribeZonesResponse, error) {
	resp := &DescribeZonesResponse{}
	err := client.Request("slb", "DescribeZones", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeZonesResponse struct {
	RequestId goaliyun.String
	Zones     DescribeZonesZoneList
}

type DescribeZonesZone struct {
	ZoneId     goaliyun.String
	LocalName  goaliyun.String
	SlaveZones DescribeZonesSlaveZoneList
}

type DescribeZonesSlaveZone struct {
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

type DescribeZonesSlaveZoneList []DescribeZonesSlaveZone

func (list *DescribeZonesSlaveZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesSlaveZone)
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

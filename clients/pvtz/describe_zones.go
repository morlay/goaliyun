package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeZonesRequest struct {
	PageSize     int64  `position:"Query" name:"PageSize"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Keyword      string `position:"Query" name:"Keyword"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeZonesRequest) Invoke(client goaliyun.Client) (*DescribeZonesResponse, error) {
	resp := &DescribeZonesResponse{}
	err := client.Request("pvtz", "DescribeZones", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeZonesResponse struct {
	RequestId  goaliyun.String
	TotalItems goaliyun.Integer
	TotalPages goaliyun.Integer
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	Zones      DescribeZonesZoneList
}

type DescribeZonesZone struct {
	ZoneId          goaliyun.String
	ZoneName        goaliyun.String
	Remark          goaliyun.String
	RecordCount     goaliyun.Integer
	CreateTime      goaliyun.String
	CreateTimestamp goaliyun.Integer
	UpdateTime      goaliyun.String
	UpdateTimestamp goaliyun.Integer
	IsPtr           bool
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

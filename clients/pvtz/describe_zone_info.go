package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeZoneInfoRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeZoneInfoRequest) Invoke(client goaliyun.Client) (*DescribeZoneInfoResponse, error) {
	resp := &DescribeZoneInfoResponse{}
	err := client.Request("pvtz", "DescribeZoneInfo", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeZoneInfoResponse struct {
	RequestId       goaliyun.String
	ZoneId          goaliyun.String
	ZoneName        goaliyun.String
	Remark          goaliyun.String
	RecordCount     goaliyun.Integer
	CreateTime      goaliyun.String
	CreateTimestamp goaliyun.Integer
	UpdateTime      goaliyun.String
	UpdateTimestamp goaliyun.Integer
	IsPtr           bool
	BindVpcs        DescribeZoneInfoVpcList
}

type DescribeZoneInfoVpc struct {
	ReionId    goaliyun.String
	VpcId      goaliyun.String
	VpcName    goaliyun.String
	RegionName goaliyun.String
}

type DescribeZoneInfoVpcList []DescribeZoneInfoVpc

func (list *DescribeZoneInfoVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneInfoVpc)
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

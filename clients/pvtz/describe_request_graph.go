package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRequestGraphRequest struct {
	VpcId          string `position:"Query" name:"VpcId"`
	UserClientIp   string `position:"Query" name:"UserClientIp"`
	ZoneId         string `position:"Query" name:"ZoneId"`
	Lang           string `position:"Query" name:"Lang"`
	StartTimestamp int64  `position:"Query" name:"StartTimestamp"`
	EndTimestamp   int64  `position:"Query" name:"EndTimestamp"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeRequestGraphRequest) Invoke(client goaliyun.Client) (*DescribeRequestGraphResponse, error) {
	resp := &DescribeRequestGraphResponse{}
	err := client.Request("pvtz", "DescribeRequestGraph", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRequestGraphResponse struct {
	RequestId      goaliyun.String
	RequestDetails DescribeRequestGraphZoneRequestTopList
}

type DescribeRequestGraphZoneRequestTop struct {
	Time         goaliyun.String
	Timestamp    goaliyun.Integer
	RequestCount goaliyun.Integer
}

type DescribeRequestGraphZoneRequestTopList []DescribeRequestGraphZoneRequestTop

func (list *DescribeRequestGraphZoneRequestTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRequestGraphZoneRequestTop)
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

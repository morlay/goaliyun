package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetOSSStatisRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	StartStatisTime      string `position:"Query" name:"StartStatisTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Level                string `position:"Query" name:"Level"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	EndStatisTime        string `position:"Query" name:"EndStatisTime"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetOSSStatisRequest) Invoke(client goaliyun.Client) (*GetOSSStatisResponse, error) {
	resp := &GetOSSStatisResponse{}
	err := client.Request("vod", "GetOSSStatis", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOSSStatisResponse struct {
	RequestId             goaliyun.String
	MaxStorageUtilization goaliyun.Integer
	OssStatisList         GetOSSStatisOSSMetricList
}

type GetOSSStatisOSSMetric struct {
	StatTime           goaliyun.String
	StorageUtilization goaliyun.Integer
}

type GetOSSStatisOSSMetricList []GetOSSStatisOSSMetric

func (list *GetOSSStatisOSSMetricList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetOSSStatisOSSMetric)
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

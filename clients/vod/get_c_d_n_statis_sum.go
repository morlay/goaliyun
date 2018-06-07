package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetCDNStatisSumRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	StartStatisTime      string `position:"Query" name:"StartStatisTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Level                string `position:"Query" name:"Level"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	EndStatisTime        string `position:"Query" name:"EndStatisTime"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetCDNStatisSumRequest) Invoke(client goaliyun.Client) (*GetCDNStatisSumResponse, error) {
	resp := &GetCDNStatisSumResponse{}
	err := client.Request("vod", "GetCDNStatisSum", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetCDNStatisSumResponse struct {
	RequestId        goaliyun.String
	SumFlowDataValue goaliyun.Integer
	MaxBpsDataValue  goaliyun.Integer
	CDNStatisList    GetCDNStatisSumCDNMetricList
}

type GetCDNStatisSumCDNMetric struct {
	StatTime              goaliyun.String
	FlowDataDomesticValue goaliyun.Integer
	FlowDataOverseasValue goaliyun.Integer
	FlowDataValue         goaliyun.Integer
	BpsDataDomesticValue  goaliyun.Integer
	BpsDataOverseasValue  goaliyun.Integer
	BpsDataValue          goaliyun.Integer
}

type GetCDNStatisSumCDNMetricList []GetCDNStatisSumCDNMetric

func (list *GetCDNStatisSumCDNMetricList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCDNStatisSumCDNMetric)
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

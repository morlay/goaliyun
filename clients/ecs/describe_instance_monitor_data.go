package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceMonitorDataRequest) Invoke(client goaliyun.Client) (*DescribeInstanceMonitorDataResponse, error) {
	resp := &DescribeInstanceMonitorDataResponse{}
	err := client.Request("ecs", "DescribeInstanceMonitorData", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceMonitorDataResponse struct {
	RequestId   goaliyun.String
	MonitorData DescribeInstanceMonitorDataInstanceMonitorDataList
}

type DescribeInstanceMonitorDataInstanceMonitorData struct {
	InstanceId                   goaliyun.String
	CPU                          goaliyun.Integer
	IntranetRX                   goaliyun.Integer
	IntranetTX                   goaliyun.Integer
	IntranetBandwidth            goaliyun.Integer
	InternetRX                   goaliyun.Integer
	InternetTX                   goaliyun.Integer
	InternetBandwidth            goaliyun.Integer
	IOPSRead                     goaliyun.Integer
	IOPSWrite                    goaliyun.Integer
	BPSRead                      goaliyun.Integer
	BPSWrite                     goaliyun.Integer
	CPUCreditUsage               goaliyun.Float
	CPUCreditBalance             goaliyun.Float
	CPUAdvanceCreditBalance      goaliyun.Float
	CPUNotpaidSurplusCreditUsage goaliyun.Float
	TimeStamp                    goaliyun.String
}

type DescribeInstanceMonitorDataInstanceMonitorDataList []DescribeInstanceMonitorDataInstanceMonitorData

func (list *DescribeInstanceMonitorDataInstanceMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceMonitorDataInstanceMonitorData)
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

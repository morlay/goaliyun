package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCenVbrHealthCheckRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	VbrInstanceOwnerId   int64  `position:"Query" name:"VbrInstanceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrInstanceId        string `position:"Query" name:"VbrInstanceId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VbrInstanceRegionId  string `position:"Query" name:"VbrInstanceRegionId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCenVbrHealthCheckRequest) Invoke(client goaliyun.Client) (*DescribeCenVbrHealthCheckResponse, error) {
	resp := &DescribeCenVbrHealthCheckResponse{}
	err := client.Request("cbn", "DescribeCenVbrHealthCheck", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCenVbrHealthCheckResponse struct {
	RequestId       goaliyun.String
	TotalCount      goaliyun.Integer
	PageNumber      goaliyun.Integer
	PageSize        goaliyun.Integer
	VbrHealthChecks DescribeCenVbrHealthCheckVbrHealthCheckList
}

type DescribeCenVbrHealthCheckVbrHealthCheck struct {
	CenId               goaliyun.String
	VbrInstanceId       goaliyun.String
	LinkStatus          goaliyun.String
	PacketLoss          goaliyun.Integer
	HealthCheckSourceIp goaliyun.String
	HealthCheckTargetIp goaliyun.String
	Delay               goaliyun.Integer
}

type DescribeCenVbrHealthCheckVbrHealthCheckList []DescribeCenVbrHealthCheckVbrHealthCheck

func (list *DescribeCenVbrHealthCheckVbrHealthCheckList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenVbrHealthCheckVbrHealthCheck)
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

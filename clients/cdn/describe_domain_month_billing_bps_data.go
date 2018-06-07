package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeDomainMonthBillingBpsDataRequest struct {
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	DomainName         string `position:"Query" name:"DomainName"`
	EndTime            string `position:"Query" name:"EndTime"`
	StartTime          string `position:"Query" name:"StartTime"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainMonthBillingBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainMonthBillingBpsDataResponse, error) {
	resp := &DescribeDomainMonthBillingBpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainMonthBillingBpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainMonthBillingBpsDataResponse struct {
	RequestId              goaliyun.String
	DomainName             goaliyun.String
	StartTime              goaliyun.String
	EndTime                goaliyun.String
	Bandwidth95Bps         goaliyun.Float
	DomesticBandwidth95Bps goaliyun.Float
	OverseasBandwidth95Bps goaliyun.Float
	MonthavgBps            goaliyun.Float
	DomesticMonthavgBps    goaliyun.Float
	OverseasMonthavgBps    goaliyun.Float
	Month4thBps            goaliyun.Float
	DomesticMonth4thBps    goaliyun.Float
	OverseasMonth4thBps    goaliyun.Float
}

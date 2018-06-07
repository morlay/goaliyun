package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeDomainMax95BpsDataRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainMax95BpsDataRequest) Invoke(client goaliyun.Client) (*DescribeDomainMax95BpsDataResponse, error) {
	resp := &DescribeDomainMax95BpsDataResponse{}
	err := client.Request("cdn", "DescribeDomainMax95BpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainMax95BpsDataResponse struct {
	RequestId        goaliyun.String
	DomainName       goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	Max95Bps         goaliyun.String
	DomesticMax95Bps goaliyun.String
	OverseasMax95Bps goaliyun.String
}

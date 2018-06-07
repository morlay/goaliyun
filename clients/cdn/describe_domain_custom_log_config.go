package cdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeDomainCustomLogConfigRequest struct {
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainCustomLogConfigRequest) Invoke(client goaliyun.Client) (*DescribeDomainCustomLogConfigResponse, error) {
	resp := &DescribeDomainCustomLogConfigResponse{}
	err := client.Request("cdn", "DescribeDomainCustomLogConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainCustomLogConfigResponse struct {
	RequestId goaliyun.String
	ConfigId  goaliyun.String
	Remark    goaliyun.String
	Sample    goaliyun.String
	Tag       goaliyun.String
}

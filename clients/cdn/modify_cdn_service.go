package cdn

import (
	"github.com/morlay/goaliyun"
)

type ModifyCdnServiceRequest struct {
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *ModifyCdnServiceRequest) Invoke(client goaliyun.Client) (*ModifyCdnServiceResponse, error) {
	resp := &ModifyCdnServiceResponse{}
	err := client.Request("cdn", "ModifyCdnService", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCdnServiceResponse struct {
	RequestId goaliyun.String
}

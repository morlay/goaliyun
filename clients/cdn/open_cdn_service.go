package cdn

import (
	"github.com/morlay/goaliyun"
)

type OpenCdnServiceRequest struct {
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *OpenCdnServiceRequest) Invoke(client goaliyun.Client) (*OpenCdnServiceResponse, error) {
	resp := &OpenCdnServiceResponse{}
	err := client.Request("cdn", "OpenCdnService", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OpenCdnServiceResponse struct {
	RequestId goaliyun.String
}

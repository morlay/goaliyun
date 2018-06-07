package market

import (
	"github.com/morlay/goaliyun"
)

type ActivateLicenseRequest struct {
	Identification string `position:"Query" name:"Identification"`
	LicenseCode    string `position:"Query" name:"LicenseCode"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ActivateLicenseRequest) Invoke(client goaliyun.Client) (*ActivateLicenseResponse, error) {
	resp := &ActivateLicenseResponse{}
	err := client.Request("market", "ActivateLicense", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActivateLicenseResponse struct {
	RequestId goaliyun.String
	Success   bool
}

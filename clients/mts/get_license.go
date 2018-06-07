package mts

import (
	"github.com/morlay/goaliyun"
)

type GetLicenseRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	LicenseUrl           string `position:"Query" name:"LicenseUrl"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetLicenseRequest) Invoke(client goaliyun.Client) (*GetLicenseResponse, error) {
	resp := &GetLicenseResponse{}
	err := client.Request("mts", "GetLicense", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetLicenseResponse struct {
	RequestId goaliyun.String
	License   goaliyun.String
}

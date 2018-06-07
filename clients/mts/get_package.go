package mts

import (
	"github.com/morlay/goaliyun"
)

type GetPackageRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetPackageRequest) Invoke(client goaliyun.Client) (*GetPackageResponse, error) {
	resp := &GetPackageResponse{}
	err := client.Request("mts", "GetPackage", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPackageResponse struct {
	RequestId   goaliyun.String
	CertPackage goaliyun.String
}

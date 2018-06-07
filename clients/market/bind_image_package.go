package market

import (
	"github.com/morlay/goaliyun"
)

type BindImagePackageRequest struct {
	EcsInstanceId          string `position:"Query" name:"EcsInstanceId"`
	ImagePackageInstanceId string `position:"Query" name:"ImagePackageInstanceId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *BindImagePackageRequest) Invoke(client goaliyun.Client) (*BindImagePackageResponse, error) {
	resp := &BindImagePackageResponse{}
	err := client.Request("market", "BindImagePackage", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BindImagePackageResponse struct {
	RequestId goaliyun.String
	Success   bool
}

package aegis

import (
	"github.com/morlay/goaliyun"
)

type UpgradeInstanceRequest struct {
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	VmNumber    int64  `position:"Query" name:"VmNumber"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	VersionCode int64  `position:"Query" name:"VersionCode"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *UpgradeInstanceRequest) Invoke(client goaliyun.Client) (*UpgradeInstanceResponse, error) {
	resp := &UpgradeInstanceResponse{}
	err := client.Request("aegis", "UpgradeInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeInstanceResponse struct {
	OrderId   goaliyun.String
	RequestId goaliyun.String
}

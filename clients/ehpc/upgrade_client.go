package ehpc

import (
	"github.com/morlay/goaliyun"
)

type UpgradeClientRequest struct {
	ClientVersion string `position:"Query" name:"ClientVersion"`
	ClusterId     string `position:"Query" name:"ClusterId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *UpgradeClientRequest) Invoke(client goaliyun.Client) (*UpgradeClientResponse, error) {
	resp := &UpgradeClientResponse{}
	err := client.Request("ehpc", "UpgradeClient", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeClientResponse struct {
	RequestId goaliyun.String
}

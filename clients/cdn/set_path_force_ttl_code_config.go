package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetPathForceTtlCodeConfigRequest struct {
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	CodeString    string `position:"Query" name:"CodeString"`
	Path          string `position:"Query" name:"Path"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetPathForceTtlCodeConfigRequest) Invoke(client goaliyun.Client) (*SetPathForceTtlCodeConfigResponse, error) {
	resp := &SetPathForceTtlCodeConfigResponse{}
	err := client.Request("cdn", "SetPathForceTtlCodeConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetPathForceTtlCodeConfigResponse struct {
	RequestId goaliyun.String
}

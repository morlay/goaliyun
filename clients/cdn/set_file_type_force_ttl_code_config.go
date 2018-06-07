package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetFileTypeForceTtlCodeConfigRequest struct {
	FileType      string `position:"Query" name:"FileType"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	CodeString    string `position:"Query" name:"CodeString"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetFileTypeForceTtlCodeConfigRequest) Invoke(client goaliyun.Client) (*SetFileTypeForceTtlCodeConfigResponse, error) {
	resp := &SetFileTypeForceTtlCodeConfigResponse{}
	err := client.Request("cdn", "SetFileTypeForceTtlCodeConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetFileTypeForceTtlCodeConfigResponse struct {
	RequestId goaliyun.String
}

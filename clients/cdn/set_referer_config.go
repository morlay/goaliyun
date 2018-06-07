package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetRefererConfigRequest struct {
	ReferList     string `position:"Query" name:"ReferList"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	ReferType     string `position:"Query" name:"ReferType"`
	DisableAst    string `position:"Query" name:"DisableAst"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	AllowEmpty    string `position:"Query" name:"AllowEmpty"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetRefererConfigRequest) Invoke(client goaliyun.Client) (*SetRefererConfigResponse, error) {
	resp := &SetRefererConfigResponse{}
	err := client.Request("cdn", "SetRefererConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetRefererConfigResponse struct {
	RequestId goaliyun.String
}

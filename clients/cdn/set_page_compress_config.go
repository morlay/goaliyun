package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetPageCompressConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetPageCompressConfigRequest) Invoke(client goaliyun.Client) (*SetPageCompressConfigResponse, error) {
	resp := &SetPageCompressConfigResponse{}
	err := client.Request("cdn", "SetPageCompressConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetPageCompressConfigResponse struct {
	RequestId goaliyun.String
}

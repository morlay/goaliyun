package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetRemoteReqAuthConfigRequest struct {
	AuthPath      string `position:"Query" name:"AuthPath"`
	DomainName    string `position:"Query" name:"DomainName"`
	AuthEnable    string `position:"Query" name:"AuthEnable"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TimeOut       string `position:"Query" name:"TimeOut"`
	AuthType      string `position:"Query" name:"AuthType"`
	AuthProvider  string `position:"Query" name:"AuthProvider"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	AuthCrash     string `position:"Query" name:"AuthCrash"`
	AuthAddr      string `position:"Query" name:"AuthAddr"`
	AuthFileType  string `position:"Query" name:"AuthFileType"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetRemoteReqAuthConfigRequest) Invoke(client goaliyun.Client) (*SetRemoteReqAuthConfigResponse, error) {
	resp := &SetRemoteReqAuthConfigResponse{}
	err := client.Request("cdn", "SetRemoteReqAuthConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetRemoteReqAuthConfigResponse struct {
	RequestId goaliyun.String
}

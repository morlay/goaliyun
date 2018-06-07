package live

import (
	"github.com/morlay/goaliyun"
)

type AddLiveMixNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddLiveMixNotifyConfigRequest) Invoke(client goaliyun.Client) (*AddLiveMixNotifyConfigResponse, error) {
	resp := &AddLiveMixNotifyConfigResponse{}
	err := client.Request("live", "AddLiveMixNotifyConfig", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveMixNotifyConfigResponse struct {
	RequestId goaliyun.String
}

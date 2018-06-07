package cdn

import (
	"github.com/morlay/goaliyun"
)

type AddLiveAppRecordConfigRequest struct {
	OssBucket       string `position:"Query" name:"OssBucket"`
	AppName         string `position:"Query" name:"AppName"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OssEndpoint     string `position:"Query" name:"OssEndpoint"`
	OssObjectPrefix string `position:"Query" name:"OssObjectPrefix"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AddLiveAppRecordConfigRequest) Invoke(client goaliyun.Client) (*AddLiveAppRecordConfigResponse, error) {
	resp := &AddLiveAppRecordConfigResponse{}
	err := client.Request("cdn", "AddLiveAppRecordConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddLiveAppRecordConfigResponse struct {
	RequestId goaliyun.String
}

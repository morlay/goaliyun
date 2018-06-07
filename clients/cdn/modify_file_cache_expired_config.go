package cdn

import (
	"github.com/morlay/goaliyun"
)

type ModifyFileCacheExpiredConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	Weight        string `position:"Query" name:"Weight"`
	CacheContent  string `position:"Query" name:"CacheContent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TTL           string `position:"Query" name:"TTL"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyFileCacheExpiredConfigRequest) Invoke(client goaliyun.Client) (*ModifyFileCacheExpiredConfigResponse, error) {
	resp := &ModifyFileCacheExpiredConfigResponse{}
	err := client.Request("cdn", "ModifyFileCacheExpiredConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFileCacheExpiredConfigResponse struct {
	RequestId goaliyun.String
}

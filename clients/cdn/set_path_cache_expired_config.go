package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetPathCacheExpiredConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Weight        string `position:"Query" name:"Weight"`
	CacheContent  string `position:"Query" name:"CacheContent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TTL           string `position:"Query" name:"TTL"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetPathCacheExpiredConfigRequest) Invoke(client goaliyun.Client) (*SetPathCacheExpiredConfigResponse, error) {
	resp := &SetPathCacheExpiredConfigResponse{}
	err := client.Request("cdn", "SetPathCacheExpiredConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetPathCacheExpiredConfigResponse struct {
	RequestId goaliyun.String
}

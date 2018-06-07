package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteCacheExpiredConfigRequest struct {
	CacheType     string `position:"Query" name:"CacheType"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DeleteCacheExpiredConfigRequest) Invoke(client goaliyun.Client) (*DeleteCacheExpiredConfigResponse, error) {
	resp := &DeleteCacheExpiredConfigResponse{}
	err := client.Request("cdn", "DeleteCacheExpiredConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCacheExpiredConfigResponse struct {
	RequestId goaliyun.String
}

package cdn

import (
	"github.com/morlay/goaliyun"
)

type PushObjectCacheRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *PushObjectCacheRequest) Invoke(client goaliyun.Client) (*PushObjectCacheResponse, error) {
	resp := &PushObjectCacheResponse{}
	err := client.Request("cdn", "PushObjectCache", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushObjectCacheResponse struct {
	RequestId  goaliyun.String
	PushTaskId goaliyun.String
}

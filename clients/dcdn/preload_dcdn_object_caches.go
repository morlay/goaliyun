package dcdn

import (
	"github.com/morlay/goaliyun"
)

type PreloadDcdnObjectCachesRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *PreloadDcdnObjectCachesRequest) Invoke(client goaliyun.Client) (*PreloadDcdnObjectCachesResponse, error) {
	resp := &PreloadDcdnObjectCachesResponse{}
	err := client.Request("dcdn", "PreloadDcdnObjectCaches", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PreloadDcdnObjectCachesResponse struct {
	RequestId     goaliyun.String
	PreloadTaskId goaliyun.String
}

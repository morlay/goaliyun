package dcdn

import (
	"github.com/morlay/goaliyun"
)

type RefreshDcdnObjectCachesRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ObjectType    string `position:"Query" name:"ObjectType"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *RefreshDcdnObjectCachesRequest) Invoke(client goaliyun.Client) (*RefreshDcdnObjectCachesResponse, error) {
	resp := &RefreshDcdnObjectCachesResponse{}
	err := client.Request("dcdn", "RefreshDcdnObjectCaches", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefreshDcdnObjectCachesResponse struct {
	RequestId     goaliyun.String
	RefreshTaskId goaliyun.String
}

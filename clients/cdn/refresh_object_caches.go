package cdn

import (
	"github.com/morlay/goaliyun"
)

type RefreshObjectCachesRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ObjectPath    string `position:"Query" name:"ObjectPath"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	ObjectType    string `position:"Query" name:"ObjectType"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *RefreshObjectCachesRequest) Invoke(client goaliyun.Client) (*RefreshObjectCachesResponse, error) {
	resp := &RefreshObjectCachesResponse{}
	err := client.Request("cdn", "RefreshObjectCaches", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefreshObjectCachesResponse struct {
	RequestId     goaliyun.String
	RefreshTaskId goaliyun.String
}

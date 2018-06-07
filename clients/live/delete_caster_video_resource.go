package live

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasterVideoResourceRequest struct {
	ResourceId string `position:"Query" name:"ResourceId"`
	CasterId   string `position:"Query" name:"CasterId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasterVideoResourceRequest) Invoke(client goaliyun.Client) (*DeleteCasterVideoResourceResponse, error) {
	resp := &DeleteCasterVideoResourceResponse{}
	err := client.Request("live", "DeleteCasterVideoResource", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasterVideoResourceResponse struct {
	RequestId goaliyun.String
}

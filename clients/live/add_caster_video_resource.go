package live

import (
	"github.com/morlay/goaliyun"
)

type AddCasterVideoResourceRequest struct {
	BeginOffset   int64  `position:"Query" name:"BeginOffset"`
	VodUrl        string `position:"Query" name:"VodUrl"`
	LiveStreamUrl string `position:"Query" name:"LiveStreamUrl"`
	LocationId    string `position:"Query" name:"LocationId"`
	CasterId      string `position:"Query" name:"CasterId"`
	EndOffset     int64  `position:"Query" name:"EndOffset"`
	ResourceName  string `position:"Query" name:"ResourceName"`
	RepeatNum     int64  `position:"Query" name:"RepeatNum"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MaterialId    string `position:"Query" name:"MaterialId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddCasterVideoResourceRequest) Invoke(client goaliyun.Client) (*AddCasterVideoResourceResponse, error) {
	resp := &AddCasterVideoResourceResponse{}
	err := client.Request("live", "AddCasterVideoResource", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCasterVideoResourceResponse struct {
	RequestId  goaliyun.String
	ResourceId goaliyun.String
}

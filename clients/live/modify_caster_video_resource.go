package live

import (
	"github.com/morlay/goaliyun"
)

type ModifyCasterVideoResourceRequest struct {
	ResourceId    string `position:"Query" name:"ResourceId"`
	BeginOffset   int64  `position:"Query" name:"BeginOffset"`
	VodUrl        string `position:"Query" name:"VodUrl"`
	LiveStreamUrl string `position:"Query" name:"LiveStreamUrl"`
	CasterId      string `position:"Query" name:"CasterId"`
	EndOffset     int64  `position:"Query" name:"EndOffset"`
	ResourceName  string `position:"Query" name:"ResourceName"`
	RepeatNum     int64  `position:"Query" name:"RepeatNum"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MaterialId    string `position:"Query" name:"MaterialId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyCasterVideoResourceRequest) Invoke(client goaliyun.Client) (*ModifyCasterVideoResourceResponse, error) {
	resp := &ModifyCasterVideoResourceResponse{}
	err := client.Request("live", "ModifyCasterVideoResource", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCasterVideoResourceResponse struct {
	RequestId  goaliyun.String
	CasterId   goaliyun.String
	ResourceId goaliyun.String
}

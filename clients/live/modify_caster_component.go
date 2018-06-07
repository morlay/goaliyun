package live

import (
	"github.com/morlay/goaliyun"
)

type ModifyCasterComponentRequest struct {
	ComponentId         string `position:"Query" name:"ComponentId"`
	ComponentType       string `position:"Query" name:"ComponentType"`
	ImageLayerContent   string `position:"Query" name:"ImageLayerContent"`
	CasterId            string `position:"Query" name:"CasterId"`
	Effect              string `position:"Query" name:"Effect"`
	ComponentLayer      string `position:"Query" name:"ComponentLayer"`
	CaptionLayerContent string `position:"Query" name:"CaptionLayerContent"`
	ComponentName       string `position:"Query" name:"ComponentName"`
	OwnerId             int64  `position:"Query" name:"OwnerId"`
	TextLayerContent    string `position:"Query" name:"TextLayerContent"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *ModifyCasterComponentRequest) Invoke(client goaliyun.Client) (*ModifyCasterComponentResponse, error) {
	resp := &ModifyCasterComponentResponse{}
	err := client.Request("live", "ModifyCasterComponent", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCasterComponentResponse struct {
	RequestId   goaliyun.String
	ComponentId goaliyun.String
}

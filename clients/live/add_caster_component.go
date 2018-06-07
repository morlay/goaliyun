package live

import (
	"github.com/morlay/goaliyun"
)

type AddCasterComponentRequest struct {
	ComponentType       string `position:"Query" name:"ComponentType"`
	LocationId          string `position:"Query" name:"LocationId"`
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

func (req *AddCasterComponentRequest) Invoke(client goaliyun.Client) (*AddCasterComponentResponse, error) {
	resp := &AddCasterComponentResponse{}
	err := client.Request("live", "AddCasterComponent", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCasterComponentResponse struct {
	RequestId   goaliyun.String
	ComponentId goaliyun.String
}

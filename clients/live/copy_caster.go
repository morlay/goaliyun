package live

import (
	"github.com/morlay/goaliyun"
)

type CopyCasterRequest struct {
	SrcCasterId string `position:"Query" name:"SrcCasterId"`
	CasterName  string `position:"Query" name:"CasterName"`
	ClientToken string `position:"Query" name:"ClientToken"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CopyCasterRequest) Invoke(client goaliyun.Client) (*CopyCasterResponse, error) {
	resp := &CopyCasterResponse{}
	err := client.Request("live", "CopyCaster", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CopyCasterResponse struct {
	RequestId goaliyun.String
	CasterId  goaliyun.String
}

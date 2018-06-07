package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApRadioConfigRequest struct {
	RequireMode        string `position:"Query" name:"RequireMode"`
	Htmode             string `position:"Query" name:"Htmode"`
	Frag               int64  `position:"Query" name:"Frag"`
	Minrate            int64  `position:"Query" name:"Minrate"`
	McastRate          int64  `position:"Query" name:"McastRate"`
	Probereq           int64  `position:"Query" name:"Probereq"`
	Channel            int64  `position:"Query" name:"Channel"`
	Shortgi            int64  `position:"Query" name:"Shortgi"`
	Hwmode             string `position:"Query" name:"Hwmode"`
	Uapsd              int64  `position:"Query" name:"Uapsd"`
	BeaconInt          int64  `position:"Query" name:"BeaconInt"`
	Mac                string `position:"Query" name:"Mac"`
	Rts                int64  `position:"Query" name:"Rts"`
	Txpower            int64  `position:"Query" name:"Txpower"`
	Noscan             int64  `position:"Query" name:"Noscan"`
	BcastRate          int64  `position:"Query" name:"BcastRate"`
	Disabled           int64  `position:"Query" name:"Disabled"`
	InstantlyEffective int64  `position:"Query" name:"InstantlyEffective"`
	Id                 int64  `position:"Query" name:"Id"`
	RadioIndex         int64  `position:"Query" name:"RadioIndex"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *SaveApRadioConfigRequest) Invoke(client goaliyun.Client) (*SaveApRadioConfigResponse, error) {
	resp := &SaveApRadioConfigResponse{}
	err := client.Request("cloudwf", "SaveApRadioConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApRadioConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}

package pvtz

import (
	"github.com/morlay/goaliyun"
)

type AddZoneRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	ZoneName     string `position:"Query" name:"ZoneName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *AddZoneRequest) Invoke(client goaliyun.Client) (*AddZoneResponse, error) {
	resp := &AddZoneResponse{}
	err := client.Request("pvtz", "AddZone", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddZoneResponse struct {
	RequestId goaliyun.String
	Success   bool
	ZoneId    goaliyun.String
	ZoneName  goaliyun.String
}

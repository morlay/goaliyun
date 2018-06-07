package pvtz

import (
	"github.com/morlay/goaliyun"
)

type CheckZoneNameRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	ZoneName     string `position:"Query" name:"ZoneName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CheckZoneNameRequest) Invoke(client goaliyun.Client) (*CheckZoneNameResponse, error) {
	resp := &CheckZoneNameResponse{}
	err := client.Request("pvtz", "CheckZoneName", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckZoneNameResponse struct {
	RequestId goaliyun.String
	Success   bool
	Check     bool
}

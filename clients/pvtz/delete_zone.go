package pvtz

import (
	"github.com/morlay/goaliyun"
)

type DeleteZoneRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteZoneRequest) Invoke(client goaliyun.Client) (*DeleteZoneResponse, error) {
	resp := &DeleteZoneResponse{}
	err := client.Request("pvtz", "DeleteZone", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteZoneResponse struct {
	RequestId goaliyun.String
	ZoneId    goaliyun.String
}

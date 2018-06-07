package pvtz

import (
	"github.com/morlay/goaliyun"
)

type UpdateZoneRemarkRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Remark       string `position:"Query" name:"Remark"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *UpdateZoneRemarkRequest) Invoke(client goaliyun.Client) (*UpdateZoneRemarkResponse, error) {
	resp := &UpdateZoneRemarkResponse{}
	err := client.Request("pvtz", "UpdateZoneRemark", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateZoneRemarkResponse struct {
	RequestId goaliyun.String
	ZoneId    goaliyun.String
}

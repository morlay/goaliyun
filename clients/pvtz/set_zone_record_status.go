package pvtz

import (
	"github.com/morlay/goaliyun"
)

type SetZoneRecordStatusRequest struct {
	RecordId     int64  `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SetZoneRecordStatusRequest) Invoke(client goaliyun.Client) (*SetZoneRecordStatusResponse, error) {
	resp := &SetZoneRecordStatusResponse{}
	err := client.Request("pvtz", "SetZoneRecordStatus", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetZoneRecordStatusResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.Integer
	Status    goaliyun.String
}

package pvtz

import (
	"github.com/morlay/goaliyun"
)

type DeleteZoneRecordRequest struct {
	RecordId     int64  `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteZoneRecordRequest) Invoke(client goaliyun.Client) (*DeleteZoneRecordResponse, error) {
	resp := &DeleteZoneRecordResponse{}
	err := client.Request("pvtz", "DeleteZoneRecord", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteZoneRecordResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.Integer
}

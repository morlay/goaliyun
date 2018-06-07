package pvtz

import (
	"github.com/morlay/goaliyun"
)

type UpdateZoneRecordRequest struct {
	Rr           string `position:"Query" name:"Rr"`
	RecordId     int64  `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Type         string `position:"Query" name:"Type"`
	Priority     int64  `position:"Query" name:"Priority"`
	Ttl          int64  `position:"Query" name:"Ttl"`
	Value        string `position:"Query" name:"Value"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *UpdateZoneRecordRequest) Invoke(client goaliyun.Client) (*UpdateZoneRecordResponse, error) {
	resp := &UpdateZoneRecordResponse{}
	err := client.Request("pvtz", "UpdateZoneRecord", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateZoneRecordResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.Integer
}

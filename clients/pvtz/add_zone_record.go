package pvtz

import (
	"github.com/morlay/goaliyun"
)

type AddZoneRecordRequest struct {
	Rr           string `position:"Query" name:"Rr"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Lang         string `position:"Query" name:"Lang"`
	Type         string `position:"Query" name:"Type"`
	Priority     int64  `position:"Query" name:"Priority"`
	Ttl          int64  `position:"Query" name:"Ttl"`
	Value        string `position:"Query" name:"Value"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *AddZoneRecordRequest) Invoke(client goaliyun.Client) (*AddZoneRecordResponse, error) {
	resp := &AddZoneRecordResponse{}
	err := client.Request("pvtz", "AddZoneRecord", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddZoneRecordResponse struct {
	RequestId goaliyun.String
	Success   bool
	RecordId  goaliyun.Integer
}

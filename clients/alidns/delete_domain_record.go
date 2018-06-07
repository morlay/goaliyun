package alidns

import (
	"github.com/morlay/goaliyun"
)

type DeleteDomainRecordRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteDomainRecordRequest) Invoke(client goaliyun.Client) (*DeleteDomainRecordResponse, error) {
	resp := &DeleteDomainRecordResponse{}
	err := client.Request("alidns", "DeleteDomainRecord", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDomainRecordResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.String
}

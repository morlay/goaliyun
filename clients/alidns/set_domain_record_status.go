package alidns

import (
	"github.com/morlay/goaliyun"
)

type SetDomainRecordStatusRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	Status       string `position:"Query" name:"Status"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SetDomainRecordStatusRequest) Invoke(client goaliyun.Client) (*SetDomainRecordStatusResponse, error) {
	resp := &SetDomainRecordStatusResponse{}
	err := client.Request("alidns", "SetDomainRecordStatus", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDomainRecordStatusResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.String
	Status    goaliyun.String
}

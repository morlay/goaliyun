package alidns

import (
	"github.com/morlay/goaliyun"
)

type DescribeDomainRecordInfoRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDomainRecordInfoRequest) Invoke(client goaliyun.Client) (*DescribeDomainRecordInfoResponse, error) {
	resp := &DescribeDomainRecordInfoResponse{}
	err := client.Request("alidns", "DescribeDomainRecordInfo", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDomainRecordInfoResponse struct {
	RequestId  goaliyun.String
	DomainId   goaliyun.String
	DomainName goaliyun.String
	PunyCode   goaliyun.String
	GroupId    goaliyun.String
	GroupName  goaliyun.String
	RecordId   goaliyun.String
	RR         goaliyun.String
	Type       goaliyun.String
	Value      goaliyun.String
	TTL        goaliyun.Integer
	Priority   goaliyun.Integer
	Line       goaliyun.String
	Status     goaliyun.String
	Locked     bool
}

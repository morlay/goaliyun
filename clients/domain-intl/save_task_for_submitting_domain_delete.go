package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveTaskForSubmittingDomainDeleteRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveTaskForSubmittingDomainDeleteRequest) Invoke(client goaliyun.Client) (*SaveTaskForSubmittingDomainDeleteResponse, error) {
	resp := &SaveTaskForSubmittingDomainDeleteResponse{}
	err := client.Request("domain-intl", "SaveTaskForSubmittingDomainDelete", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveTaskForSubmittingDomainDeleteResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

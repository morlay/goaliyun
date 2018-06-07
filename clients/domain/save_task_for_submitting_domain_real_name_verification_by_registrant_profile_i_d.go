package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest struct {
	InstanceId          string `position:"Query" name:"InstanceId"`
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	DomainName          string `position:"Query" name:"DomainName"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                string `position:"Query" name:"Lang"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) Invoke(client goaliyun.Client) (*SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse, error) {
	resp := &SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse{}
	err := client.Request("domain", "SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileID", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

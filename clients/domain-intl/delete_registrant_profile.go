package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type DeleteRegistrantProfileRequest struct {
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                string `position:"Query" name:"Lang"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *DeleteRegistrantProfileRequest) Invoke(client goaliyun.Client) (*DeleteRegistrantProfileResponse, error) {
	resp := &DeleteRegistrantProfileResponse{}
	err := client.Request("domain-intl", "DeleteRegistrantProfile", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteRegistrantProfileResponse struct {
	RequestId goaliyun.String
}

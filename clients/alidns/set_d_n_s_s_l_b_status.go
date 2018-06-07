package alidns

import (
	"github.com/morlay/goaliyun"
)

type SetDNSSLBStatusRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	SubDomain    string `position:"Query" name:"SubDomain"`
	Open         string `position:"Query" name:"Open"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SetDNSSLBStatusRequest) Invoke(client goaliyun.Client) (*SetDNSSLBStatusResponse, error) {
	resp := &SetDNSSLBStatusResponse{}
	err := client.Request("alidns", "SetDNSSLBStatus", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDNSSLBStatusResponse struct {
	RequestId   goaliyun.String
	RecordCount goaliyun.Integer
	Open        bool
}

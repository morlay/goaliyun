package yundun

import (
	"github.com/morlay/goaliyun"
)

type SummaryRequest struct {
	JstOwnerId  int64  `position:"Query" name:"JstOwnerId"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SummaryRequest) Invoke(client goaliyun.Client) (*SummaryResponse, error) {
	resp := &SummaryResponse{}
	err := client.Request("yundun", "Summary", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SummaryResponse struct {
	RequestId         goaliyun.String
	Status            goaliyun.Integer
	AbnormalHostCount goaliyun.Integer
	Ddos              SummaryDdos
	BruteForce        SummaryBruteForce
	Webshell          SummaryWebshell
	RemoteLogin       SummaryRemoteLogin
	WebAttack         SummaryWebAttack
	WebLeak           SummaryWebLeak
}

type SummaryDdos struct {
	Count     goaliyun.Integer
	HostCount goaliyun.Integer
}

type SummaryBruteForce struct {
	Count     goaliyun.Integer
	HostCount goaliyun.Integer
}

type SummaryWebshell struct {
	Count     goaliyun.Integer
	HostCount goaliyun.Integer
}

type SummaryRemoteLogin struct {
	Count     goaliyun.Integer
	HostCount goaliyun.Integer
}

type SummaryWebAttack struct {
	Count     goaliyun.Integer
	HostCount goaliyun.Integer
}

type SummaryWebLeak struct {
	Count     goaliyun.Integer
	HostCount goaliyun.Integer
}

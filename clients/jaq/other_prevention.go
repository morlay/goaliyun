package jaq

import (
	"github.com/morlay/goaliyun"
)

type OtherPreventionRequest struct {
	CallerName      string `position:"Query" name:"CallerName"`
	Ip              string `position:"Query" name:"Ip"`
	ProtocolVersion string `position:"Query" name:"ProtocolVersion"`
	Source          int64  `position:"Query" name:"Source"`
	PhoneNumber     string `position:"Query" name:"PhoneNumber"`
	Email           string `position:"Query" name:"Email"`
	UserId          string `position:"Query" name:"UserId"`
	IdType          int64  `position:"Query" name:"IdType"`
	CurrentUrl      string `position:"Query" name:"CurrentUrl"`
	Agent           string `position:"Query" name:"Agent"`
	Cookie          string `position:"Query" name:"Cookie"`
	SessionId       string `position:"Query" name:"SessionId"`
	MacAddress      string `position:"Query" name:"MacAddress"`
	Referer         string `position:"Query" name:"Referer"`
	UserName        string `position:"Query" name:"UserName"`
	CompanyName     string `position:"Query" name:"CompanyName"`
	Address         string `position:"Query" name:"Address"`
	IDNumber        string `position:"Query" name:"IDNumber"`
	BankCardNumber  string `position:"Query" name:"BankCardNumber"`
	RegisterIp      string `position:"Query" name:"RegisterIp"`
	RegisterDate    int64  `position:"Query" name:"RegisterDate"`
	LoginIp         string `position:"Query" name:"LoginIp"`
	LoginDate       int64  `position:"Query" name:"LoginDate"`
	ExtendData      string `position:"Query" name:"ExtendData"`
	PasswordHash    string `position:"Query" name:"PasswordHash"`
	JsToken         string `position:"Query" name:"JsToken"`
	SDKToken        string `position:"Query" name:"SDKToken"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *OtherPreventionRequest) Invoke(client goaliyun.Client) (*OtherPreventionResponse, error) {
	resp := &OtherPreventionResponse{}
	err := client.Request("jaq", "OtherPrevention", "2016-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OtherPreventionResponse struct {
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Data      OtherPreventionData
}

type OtherPreventionData struct {
	FnalDecision     goaliyun.Integer
	EventId          goaliyun.String
	UserId           goaliyun.String
	FinalScore       goaliyun.Integer
	FinalDesc        goaliyun.String
	Detail           goaliyun.String
	CaptchaCheckData goaliyun.String
}

package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApRadioSsidConfigRequest struct {
	Nasid               string `position:"Query" name:"Nasid"`
	AuthPort            int64  `position:"Query" name:"AuthPort"`
	Hidden              int64  `position:"Query" name:"Hidden"`
	DynamicVlan         int64  `position:"Query" name:"DynamicVlan"`
	AuthServer          string `position:"Query" name:"AuthServer"`
	SecondaryAcctServer string `position:"Query" name:"SecondaryAcctServer"`
	Ssid                string `position:"Query" name:"Ssid"`
	Cir                 int64  `position:"Query" name:"Cir"`
	Mac                 string `position:"Query" name:"Mac"`
	SecondaryAcctSecret string `position:"Query" name:"SecondaryAcctSecret"`
	Ieee80211w          int64  `position:"Query" name:"Ieee.80211.w"`
	Network             int64  `position:"Query" name:"Network"`
	Isolate             int64  `position:"Query" name:"Isolate"`
	ApAssetId           int64  `position:"Query" name:"ApAssetId"`
	EncKey              string `position:"Query" name:"EncKey"`
	MulticastForward    int64  `position:"Query" name:"MulticastForward"`
	Encryption          string `position:"Query" name:"Encryption"`
	Wmm                 int64  `position:"Query" name:"Wmm"`
	AuthCache           int64  `position:"Query" name:"AuthCache"`
	Disabled            int64  `position:"Query" name:"Disabled"`
	Id                  int64  `position:"Query" name:"Id"`
	RadioIndex          int64  `position:"Query" name:"RadioIndex"`
	IgnoreWeakProbe     int64  `position:"Query" name:"IgnoreWeakProbe"`
	Maxassoc            int64  `position:"Query" name:"Maxassoc"`
	AcctServer          string `position:"Query" name:"AcctServer"`
	SecondaryAuthServer string `position:"Query" name:"SecondaryAuthServer"`
	DaeClient           string `position:"Query" name:"DaeClient"`
	DaeSecret           string `position:"Query" name:"DaeSecret"`
	DisassocLowAck      int64  `position:"Query" name:"DisassocLowAck"`
	SecondaryAuthPort   int64  `position:"Query" name:"SecondaryAuthPort"`
	AcctSecret          string `position:"Query" name:"AcctSecret"`
	DisassocWeakRssi    int64  `position:"Query" name:"DisassocWeakRssi"`
	SecondaryAcctPort   int64  `position:"Query" name:"SecondaryAcctPort"`
	DaePort             int64  `position:"Query" name:"DaePort"`
	SsidLb              int64  `position:"Query" name:"SsidLb"`
	AcctPort            int64  `position:"Query" name:"AcctPort"`
	MaxInactivity       int64  `position:"Query" name:"MaxInactivity"`
	VlanDhcp            int64  `position:"Query" name:"VlanDhcp"`
	InstantlyEffective  int64  `position:"Query" name:"InstantlyEffective"`
	ShortPreamble       int64  `position:"Query" name:"ShortPreamble"`
	AuthSecret          string `position:"Query" name:"AuthSecret"`
	SecondaryAuthSecret string `position:"Query" name:"SecondaryAuthSecret"`
	Ownip               string `position:"Query" name:"Ownip"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *SaveApRadioSsidConfigRequest) Invoke(client goaliyun.Client) (*SaveApRadioSsidConfigResponse, error) {
	resp := &SaveApRadioSsidConfigResponse{}
	err := client.Request("cloudwf", "SaveApRadioSsidConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApRadioSsidConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}

package ubsms

import (
	"github.com/morlay/goaliyun"
)

type NotifyUserBusinessCommandRequest struct {
	Uid         string `position:"Query" name:"Uid"`
	ServiceCode string `position:"Query" name:"ServiceCode"`
	Cmd         string `position:"Query" name:"Cmd"`
	Region      string `position:"Query" name:"Region"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	Password    string `position:"Query" name:"Password"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *NotifyUserBusinessCommandRequest) Invoke(client goaliyun.Client) (*NotifyUserBusinessCommandResponse, error) {
	resp := &NotifyUserBusinessCommandResponse{}
	err := client.Request("ubsms", "NotifyUserBusinessCommand", "2015-06-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NotifyUserBusinessCommandResponse struct {
	RequestId goaliyun.String
	Success   bool
}

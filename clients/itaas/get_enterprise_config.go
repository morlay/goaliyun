package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetEnterpriseConfigRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetEnterpriseConfigRequest) Invoke(client goaliyun.Client) (*GetEnterpriseConfigResponse, error) {
	resp := &GetEnterpriseConfigResponse{}
	err := client.Request("itaas", "GetEnterpriseConfig", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetEnterpriseConfigResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList GetEnterpriseConfigErrorMessageList
	Data      GetEnterpriseConfigData
}

type GetEnterpriseConfigErrorMessage struct {
	ErrorMessage goaliyun.String
}

type GetEnterpriseConfigData struct {
	AuthorizationNeedAccessToken bool
	DrMeetingQrUrl               goaliyun.String
	DrWelcomeUrl                 goaliyun.String
	ShareMboxNubmer              goaliyun.Integer
	ShareNeedInternet            bool
	ShareServiceFlag             bool
}

type GetEnterpriseConfigErrorMessageList []GetEnterpriseConfigErrorMessage

func (list *GetEnterpriseConfigErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEnterpriseConfigErrorMessage)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

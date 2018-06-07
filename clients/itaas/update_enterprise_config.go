package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpdateEnterpriseConfigRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	ConfigKey   string `position:"Query" name:"ConfigKey"`
	ConfigValue string `position:"Query" name:"ConfigValue"`
	Memo        string `position:"Query" name:"Memo"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *UpdateEnterpriseConfigRequest) Invoke(client goaliyun.Client) (*UpdateEnterpriseConfigResponse, error) {
	resp := &UpdateEnterpriseConfigResponse{}
	err := client.Request("itaas", "UpdateEnterpriseConfig", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateEnterpriseConfigResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList UpdateEnterpriseConfigErrorMessageList
}

type UpdateEnterpriseConfigErrorMessage struct {
	ErrorMessage goaliyun.String
}

type UpdateEnterpriseConfigErrorMessageList []UpdateEnterpriseConfigErrorMessage

func (list *UpdateEnterpriseConfigErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateEnterpriseConfigErrorMessage)
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

package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateEnterpriseRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	BoxNumber   int64  `position:"Query" name:"BoxNumber"`
	ServiceFlag string `position:"Query" name:"ServiceFlag"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CreateEnterpriseRequest) Invoke(client goaliyun.Client) (*CreateEnterpriseResponse, error) {
	resp := &CreateEnterpriseResponse{}
	err := client.Request("itaas", "CreateEnterprise", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateEnterpriseResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList CreateEnterpriseErrorMessageList
}

type CreateEnterpriseErrorMessage struct {
	ErrorMessage goaliyun.String
}

type CreateEnterpriseErrorMessageList []CreateEnterpriseErrorMessage

func (list *CreateEnterpriseErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateEnterpriseErrorMessage)
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

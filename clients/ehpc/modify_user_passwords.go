package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyUserPasswordsRequest struct {
	ClusterId string                       `position:"Query" name:"ClusterId"`
	Users     *ModifyUserPasswordsUserList `position:"Query" type:"Repeated" name:"User"`
	RegionId  string                       `position:"Query" name:"RegionId"`
}

func (req *ModifyUserPasswordsRequest) Invoke(client goaliyun.Client) (*ModifyUserPasswordsResponse, error) {
	resp := &ModifyUserPasswordsResponse{}
	err := client.Request("ehpc", "ModifyUserPasswords", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyUserPasswordsUser struct {
	Name     string `name:"Name"`
	Password string `name:"Password"`
}

type ModifyUserPasswordsResponse struct {
	RequestId goaliyun.String
}

type ModifyUserPasswordsUserList []ModifyUserPasswordsUser

func (list *ModifyUserPasswordsUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyUserPasswordsUser)
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

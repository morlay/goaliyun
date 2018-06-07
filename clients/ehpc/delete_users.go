package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteUsersRequest struct {
	ClusterId string               `position:"Query" name:"ClusterId"`
	Users     *DeleteUsersUserList `position:"Query" type:"Repeated" name:"User"`
	RegionId  string               `position:"Query" name:"RegionId"`
}

func (req *DeleteUsersRequest) Invoke(client goaliyun.Client) (*DeleteUsersResponse, error) {
	resp := &DeleteUsersResponse{}
	err := client.Request("ehpc", "DeleteUsers", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteUsersUser struct {
	Name string `name:"Name"`
}

type DeleteUsersResponse struct {
	RequestId goaliyun.String
}

type DeleteUsersUserList []DeleteUsersUser

func (list *DeleteUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteUsersUser)
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

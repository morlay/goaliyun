package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAliasesRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryAliasesRequest) Invoke(client goaliyun.Client) (*QueryAliasesResponse, error) {
	resp := &QueryAliasesResponse{}
	err := client.Request("push", "QueryAliases", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAliasesResponse struct {
	RequestId  goaliyun.String
	AliasInfos QueryAliasesAliasInfoList
}

type QueryAliasesAliasInfo struct {
	AliasName goaliyun.String
}

type QueryAliasesAliasInfoList []QueryAliasesAliasInfo

func (list *QueryAliasesAliasInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAliasesAliasInfo)
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

package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MetastoreListTablesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreListTablesRequest) Invoke(client goaliyun.Client) (*MetastoreListTablesResponse, error) {
	resp := &MetastoreListTablesResponse{}
	err := client.Request("emr", "MetastoreListTables", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreListTablesResponse struct {
	RequestId  goaliyun.String
	TableNames MetastoreListTablesTableNameList
}

type MetastoreListTablesTableNameList []goaliyun.String

func (list *MetastoreListTablesTableNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

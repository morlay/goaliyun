package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MetastoreSearchTablesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	DbName          string `position:"Query" name:"DbName"`
	TableName       string `position:"Query" name:"TableName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *MetastoreSearchTablesRequest) Invoke(client goaliyun.Client) (*MetastoreSearchTablesResponse, error) {
	resp := &MetastoreSearchTablesResponse{}
	err := client.Request("emr", "MetastoreSearchTables", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreSearchTablesResponse struct {
	RequestId  goaliyun.String
	TableNames MetastoreSearchTablesTableNameList
}

type MetastoreSearchTablesTableNameList []goaliyun.String

func (list *MetastoreSearchTablesTableNameList) UnmarshalJSON(data []byte) error {
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

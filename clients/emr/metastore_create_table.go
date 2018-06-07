package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MetastoreCreateTableRequest struct {
	ResourceOwnerId int64                           `position:"Query" name:"ResourceOwnerId"`
	FieldDelimiter  string                          `position:"Query" name:"FieldDelimiter"`
	DbName          string                          `position:"Query" name:"DbName"`
	Columns         *MetastoreCreateTableColumnList `position:"Query" type:"Repeated" name:"Column"`
	LocationUri     string                          `position:"Query" name:"LocationUri"`
	TableName       string                          `position:"Query" name:"TableName"`
	RegionId        string                          `position:"Query" name:"RegionId"`
}

func (req *MetastoreCreateTableRequest) Invoke(client goaliyun.Client) (*MetastoreCreateTableResponse, error) {
	resp := &MetastoreCreateTableResponse{}
	err := client.Request("emr", "MetastoreCreateTable", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MetastoreCreateTableColumn struct {
	Name    string `name:"Name"`
	Type    string `name:"Type"`
	Comment string `name:"Comment"`
}

type MetastoreCreateTableResponse struct {
	RequestId goaliyun.String
}

type MetastoreCreateTableColumnList []MetastoreCreateTableColumn

func (list *MetastoreCreateTableColumnList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MetastoreCreateTableColumn)
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

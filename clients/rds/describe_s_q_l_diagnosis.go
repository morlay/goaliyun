package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLDiagnosisRequest struct {
	SQLDiagId    string `position:"Query" name:"SQLDiagId"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeSQLDiagnosisRequest) Invoke(client goaliyun.Client) (*DescribeSQLDiagnosisResponse, error) {
	resp := &DescribeSQLDiagnosisResponse{}
	err := client.Request("rds", "DescribeSQLDiagnosis", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLDiagnosisResponse struct {
	RequestId goaliyun.String
	SQLList   DescribeSQLDiagnosisSQLListList
}

type DescribeSQLDiagnosisSQLListList []goaliyun.String

func (list *DescribeSQLDiagnosisSQLListList) UnmarshalJSON(data []byte) error {
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

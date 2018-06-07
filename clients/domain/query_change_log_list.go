package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryChangeLogListRequest struct {
	EndDate      int64  `position:"Query" name:"EndDate"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	Lang         string `position:"Query" name:"Lang"`
	PageNum      int64  `position:"Query" name:"PageNum"`
	StartDate    int64  `position:"Query" name:"StartDate"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *QueryChangeLogListRequest) Invoke(client goaliyun.Client) (*QueryChangeLogListResponse, error) {
	resp := &QueryChangeLogListResponse{}
	err := client.Request("domain", "QueryChangeLogList", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryChangeLogListResponse struct {
	RequestId      goaliyun.String
	TotalItemNum   goaliyun.Integer
	CurrentPageNum goaliyun.Integer
	TotalPageNum   goaliyun.Integer
	PageSize       goaliyun.Integer
	PrePage        bool
	NextPage       bool
	ResultLimit    bool
	Data           QueryChangeLogListChangeLogList
}

type QueryChangeLogListChangeLog struct {
	DomainName         goaliyun.String
	Result             goaliyun.String
	Operation          goaliyun.String
	OperationIPAddress goaliyun.String
	Details            goaliyun.String
	Time               goaliyun.String
}

type QueryChangeLogListChangeLogList []QueryChangeLogListChangeLog

func (list *QueryChangeLogListChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryChangeLogListChangeLog)
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

package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVpnConnectionLogsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MinutePeriod         int64  `position:"Query" name:"MinutePeriod"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	From                 int64  `position:"Query" name:"From"`
	To                   int64  `position:"Query" name:"To"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpnConnectionLogsRequest) Invoke(client goaliyun.Client) (*DescribeVpnConnectionLogsResponse, error) {
	resp := &DescribeVpnConnectionLogsResponse{}
	err := client.Request("vpc", "DescribeVpnConnectionLogs", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpnConnectionLogsResponse struct {
	RequestId   goaliyun.String
	Count       goaliyun.Integer
	IsCompleted bool
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	Data        DescribeVpnConnectionLogsDatumList
}

type DescribeVpnConnectionLogsDatumList []goaliyun.String

func (list *DescribeVpnConnectionLogsDatumList) UnmarshalJSON(data []byte) error {
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

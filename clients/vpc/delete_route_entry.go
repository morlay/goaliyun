package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteRouteEntryRequest struct {
	ResourceOwnerId      int64                            `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	DestinationCidrBlock string                           `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string                           `position:"Query" name:"OwnerAccount"`
	NextHopId            string                           `position:"Query" name:"NextHopId"`
	OwnerId              int64                            `position:"Query" name:"OwnerId"`
	NextHopLists         *DeleteRouteEntryNextHopListList `position:"Query" type:"Repeated" name:"NextHopList"`
	RouteTableId         string                           `position:"Query" name:"RouteTableId"`
	RegionId             string                           `position:"Query" name:"RegionId"`
}

func (req *DeleteRouteEntryRequest) Invoke(client goaliyun.Client) (*DeleteRouteEntryResponse, error) {
	resp := &DeleteRouteEntryResponse{}
	err := client.Request("vpc", "DeleteRouteEntry", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteRouteEntryNextHopList struct {
	NextHopType string `name:"NextHopType"`
	NextHopId   string `name:"NextHopId"`
}

type DeleteRouteEntryResponse struct {
	RequestId goaliyun.String
}

type DeleteRouteEntryNextHopListList []DeleteRouteEntryNextHopList

func (list *DeleteRouteEntryNextHopListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteRouteEntryNextHopList)
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

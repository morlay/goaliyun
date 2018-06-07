package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateRouteEntryRequest struct {
	ResourceOwnerId      int64                            `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                           `position:"Query" name:"ClientToken"`
	DestinationCidrBlock string                           `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string                           `position:"Query" name:"OwnerAccount"`
	NextHopId            string                           `position:"Query" name:"NextHopId"`
	OwnerId              int64                            `position:"Query" name:"OwnerId"`
	NextHopType          string                           `position:"Query" name:"NextHopType"`
	NextHopLists         *CreateRouteEntryNextHopListList `position:"Query" type:"Repeated" name:"NextHopList"`
	RouteTableId         string                           `position:"Query" name:"RouteTableId"`
	RegionId             string                           `position:"Query" name:"RegionId"`
}

func (req *CreateRouteEntryRequest) Invoke(client goaliyun.Client) (*CreateRouteEntryResponse, error) {
	resp := &CreateRouteEntryResponse{}
	err := client.Request("vpc", "CreateRouteEntry", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateRouteEntryNextHopList struct {
	NextHopType string `name:"NextHopType"`
	NextHopId   string `name:"NextHopId"`
	Weight      int64  `name:"Weight"`
}

type CreateRouteEntryResponse struct {
	RequestId goaliyun.String
}

type CreateRouteEntryNextHopListList []CreateRouteEntryNextHopList

func (list *CreateRouteEntryNextHopListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateRouteEntryNextHopList)
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

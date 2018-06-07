package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddCasterEpisodeGroupRequest struct {
	SideOutputUrl string                         `position:"Query" name:"SideOutputUrl"`
	Items         *AddCasterEpisodeGroupItemList `position:"Query" type:"Repeated" name:"Item"`
	ClientToken   string                         `position:"Query" name:"ClientToken"`
	DomainName    string                         `position:"Query" name:"DomainName"`
	StartTime     string                         `position:"Query" name:"StartTime"`
	RepeatNum     int64                          `position:"Query" name:"RepeatNum"`
	CallbackUrl   string                         `position:"Query" name:"CallbackUrl"`
	OwnerId       int64                          `position:"Query" name:"OwnerId"`
	RegionId      string                         `position:"Query" name:"RegionId"`
}

func (req *AddCasterEpisodeGroupRequest) Invoke(client goaliyun.Client) (*AddCasterEpisodeGroupResponse, error) {
	resp := &AddCasterEpisodeGroupResponse{}
	err := client.Request("live", "AddCasterEpisodeGroup", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCasterEpisodeGroupItem struct {
	ItemName string `name:"ItemName"`
	VodUrl   string `name:"VodUrl"`
}

type AddCasterEpisodeGroupResponse struct {
	RequestId goaliyun.String
	ProgramId goaliyun.String
	ItemIds   AddCasterEpisodeGroupItemIdList
}

type AddCasterEpisodeGroupItemList []AddCasterEpisodeGroupItem

func (list *AddCasterEpisodeGroupItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterEpisodeGroupItem)
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

type AddCasterEpisodeGroupItemIdList []goaliyun.String

func (list *AddCasterEpisodeGroupItemIdList) UnmarshalJSON(data []byte) error {
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

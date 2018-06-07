package ram

import (
	"github.com/morlay/goaliyun"
)

type SetSecurityPreferenceRequest struct {
	AllowUserToManageAccessKeys string `position:"Query" name:"AllowUserToManageAccessKeys"`
	AllowUserToManageMFADevices string `position:"Query" name:"AllowUserToManageMFADevices"`
	AllowUserToManagePublicKeys string `position:"Query" name:"AllowUserToManagePublicKeys"`
	EnableSaveMFATicket         string `position:"Query" name:"EnableSaveMFATicket"`
	AllowUserToChangePassword   string `position:"Query" name:"AllowUserToChangePassword"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *SetSecurityPreferenceRequest) Invoke(client goaliyun.Client) (*SetSecurityPreferenceResponse, error) {
	resp := &SetSecurityPreferenceResponse{}
	err := client.Request("ram", "SetSecurityPreference", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetSecurityPreferenceResponse struct {
	RequestId          goaliyun.String
	SecurityPreference SetSecurityPreferenceSecurityPreference
}

type SetSecurityPreferenceSecurityPreference struct {
	LoginProfilePreference SetSecurityPreferenceLoginProfilePreference
	AccessKeyPreference    SetSecurityPreferenceAccessKeyPreference
	PublicKeyPreference    SetSecurityPreferencePublicKeyPreference
	MFAPreference          SetSecurityPreferenceMFAPreference
}

type SetSecurityPreferenceLoginProfilePreference struct {
	EnableSaveMFATicket       bool
	AllowUserToChangePassword bool
}

type SetSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys bool
}

type SetSecurityPreferencePublicKeyPreference struct {
	AllowUserToManagePublicKeys bool
}

type SetSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices bool
}

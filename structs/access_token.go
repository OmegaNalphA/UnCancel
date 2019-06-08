package structs

//PersonalAccessToken is the Facebook Access Token Struct
type PersonalAccessToken struct {
	AccessToken string `json:"at"`
	FBID        string `json:"fbID"`
}

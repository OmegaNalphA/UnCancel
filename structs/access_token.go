package structs

//FBAccessToken is the Facebook Access Token Struct
type FBAccessToken struct {
	AccessToken string `json:"at"`
	FBID        string `json:"fbID"`
}

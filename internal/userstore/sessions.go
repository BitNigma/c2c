package userstore

type Session struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Time  string `json:"time"`
}

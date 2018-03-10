package iksm

type stage struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Id    string `json:"id"`
}

type stages struct {
	Stage []stage `json:"stages"`
}

package config

type Market struct {
	Market   string
	Region   string
	Language string
	Country  string
}


type Attribute struct {
	Key string `json:"key"`
	Val string `json:"val"`
}


type RMB struct {
	Market Attribute
	Region Attribute
	Language Attribute
	Country Attribute
}

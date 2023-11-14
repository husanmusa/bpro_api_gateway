package models

type Region struct {
	RegionId    int32  `json:"regionId"`
	Name        string `json:"name"`
	NameUzCyrl  string `json:"nameUzCyrl"`
	NameUzLatn  string `json:"nameUzLatn"`
	NameRu      string `json:"nameRu"`
	Active      int32  `json:"active"`
	DistrictsId int32  `json:"districtsId"`
	Code        int32  `json:"code"`
}

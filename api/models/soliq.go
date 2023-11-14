package models

type ReferenceStatus struct {
	Active     bool   `json:"active"`
	IsSafe     bool   `json:"isSafe"`
	Suspended  bool   `json:"suspended"`
	VatPayer   bool   `json:"vatPayer"`
	VatRegCode string `json:"vatRegCode"`
}

type TaxpayerType struct {
	Code   int16  `json:"code"`
	Name   string `json:"name"`
	NameRu string `json:"nameRu"`
	NameUz string `json:"nameUz"`
}
type TaxGap struct {
	Date       string   `json:"date"`
	TaxGap     *float32 `json:"taxGap"`
	TinOrPinfl string   `json:"tinOrPinfl"`
}
type DataStatus struct {
	Data    ReferenceStatus `json:"data"`
	Reason  interface{}     `json:"reason"`
	Success bool            `json:"success"`
}

type DataTaxpayer struct {
	Data    TaxpayerType `json:"data"`
	Reason  interface{}  `json:"reason"`
	Success bool         `json:"success"`
}

type FacturaByTin struct {
	Account       string          `json:"account"`
	Accountant    string          `json:"accountant"`
	Address       string          `json:"address"`
	Director      string          `json:"director"`
	DirectorPinfl int64           `json:"directorPinfl"`
	DirectorTin   string          `json:"directorTin"`
	IsBudget      int             `json:"isBudget"`
	Mfo           string          `json:"mfo"`
	Na1Code       int             `json:"na1Code"`
	Na1Name       string          `json:"na1Name"`
	Name          string          `json:"name"`
	Ns10Code      int             `json:"ns10Code"`
	Ns11Code      int             `json:"ns11Code"`
	Oked          string          `json:"oked"`
	ShortName     string          `json:"shortName"`
	StatusCode    int             `json:"statusCode"`
	StatusName    string          `json:"statusName"`
	Tin           string          `json:"tin"`
	Status        ReferenceStatus `json:"status"`
	Taxpayer      TaxpayerType    `json:"taxpayer"`
	TaxGap        TaxGap          `json:"taxGap"`
}

type Provider struct {
	Enabled      bool   `json:"enabled"`
	ProviderName string `json:"providerName"`
	ProviderTin  string `json:"providerTin"`
	SystemName   string `json:"systemName"`
}

type GetProviders struct {
	Providers []Provider `json:"providers"`
}

type OtherInn struct {
	OtherInn   string `json:"otherInn"`
	BranchCode string `json:"branchCode"`
}

type PhisFactura struct {
	Address       *string `json:"address"`
	FullName      *string `json:"fullName"`
	IsItd         *bool   `json:"isItd"`
	Ns10Code      *int32  `json:"ns10Code"`
	Ns11Code      *int32  `json:"ns11Code"`
	PassIssueDate *string `json:"passIssueDate"`
	PassNumber    *string `json:"passNumber"`
	PassOrg       *string `json:"passOrg"`
	PassSeries    *string `json:"passSeries"`
	PersonalNum   *string `json:"personalNum"`
	Tin           *string `json:"tin"`
	Name          *string `json:"name"`
}

type ByInn struct {
	Inn string `json:"inn"`
}

type BalanceBilling struct {
	Tin            string  `json:"tin"`
	IsActive       int     `json:"isActive"`
	CurrentBalance float32 `json:"currentBalance"`
	DocPrice       float32 `json:"docPrice"`
}

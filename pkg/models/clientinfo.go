package models

type IPInfo struct {
	IP                 string  `json:"ip" gorm:"type:varchar(50);unique;"`
	Network            string  `json:"network" gorm:"type:varchar(50)"`
	Version            string  `json:"version" gorm:"type:varchar(50)"`
	City               string  `json:"city" gorm:"type:varchar(50)"`
	Region             string  `json:"region" gorm:"type:varchar(50)"`
	RegionCode         string  `json:"region_code" gorm:"type:varchar(50)"`
	Country            string  `json:"country" gorm:"type:varchar(50)"`
	CountryName        string  `json:"country_name" gorm:"type:varchar(50)"`
	CountryCode        string  `json:"country_code" gorm:"type:varchar(50)"`
	CountryCodeISO3    string  `json:"country_code_iso3" gorm:"type:varchar(50)"`
	CountryCapital     string  `json:"country_capital" gorm:"type:varchar(50)"`
	CountryTLD         string  `json:"country_tld" gorm:"type:varchar(50)"`
	ContinentCode      string  `json:"continent_code" gorm:"type:varchar(50)"`
	InEU               bool    `json:"in_eu"`
	Postal             *string `json:"postal" gorm:"type:varchar(50)"` // Use pointer for nullable field
	Latitude           float64 `json:"latitude" gorm:"type:varchar(50)"`
	Longitude          float64 `json:"longitude" gorm:"type:varchar(50)"`
	TimeZone           string  `json:"timezone" gorm:"type:varchar(50)"`
	UTCOffset          string  `json:"utc_offset" gorm:"type:varchar(50)"`
	CountryCallingCode string  `json:"country_calling_code" gorm:"type:varchar(50)"`
	Currency           string  `json:"currency" gorm:"type:varchar(50)"`
	CurrencyName       string  `json:"currency_name" gorm:"type:varchar(50)"`
	Languages          string  `json:"languages" gorm:"type:varchar(50)"`
	CountryArea        float64 `json:"country_area" gorm:"type:varchar(50)"`
	CountryPopulation  int     `json:"country_population" gorm:"type:varchar(50)"`
	ASN                string  `json:"asn" gorm:"type:varchar(50)"`
	Org                string  `json:"org" gorm:"type:varchar(50)"`
}

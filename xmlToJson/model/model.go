package model

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

const (
	SHORT_DATE_FORM = "2006-01-02"
)

type Housing struct {
	Member *MemberAttr `xml:"member" json:"member"`
}

type MemberAttr struct {
	EstateAgentId string     `xml:"estateAgentID" json:"estateAgentID"`
	Items         *ItemsAttr `xml:"items" json:"items"`
}

type ItemsAttr struct {
	Item []*ItemAttr `xml:"item" json:"item"`
}

type CustomTime struct {
	time.Time
}

type CustomInteger int

type ItemAttr struct {
	UniqueObjectId           string        `xml:"uniqueobjectid" json:"uniqueobjectid"`
	Street                   string        `xml:"street" json:"street"`
	HouseNumber              CustomInteger `xml:"houseNumber" json:"houseNumber"`                         // houseNumber : string - The source(xml) value is CDATA with not number letter
	HouseNumberAddtion       string        `xml:"houseNumberAddtion" json:"houseNumberAddtion,omitempty"` // string - The source value is empty.
	PostalCode               string        `xml:"postalCode" json:"postalCode"`
	City                     string        `xml:"City" json:"City"`
	SubArea                  string        `xml:"SubArea" json:"SubArea"`
	HouseType                int           `xml:"houseType" json:"houseType"`
	Nieuwbouw                int           `xml:"nieuwbouw" json:"Nieuwbouw"`
	Furnished                int           `xml:"furnished" json:"furnished"`
	MinPrice                 CustomInteger `xml:"minprice" json:"minprice"` // string - The source value is CDATA.
	NroOfRooms               int           `xml:"NroOfRooms" json:"NroOfRooms"`
	NroOfLivingRooms         int           `xml:"NroOfLivingRooms" json:"NroOfLivingRooms"`
	Projectnaam              string        `xml:"Projectnaam" json:"Projectnaam,omitempty"`                 // string - The source value is empty.
	WoningTypeInProject      string        `xml:"WoningtypeInProject" json:"WoningtypeInProject,omitempty"` // string - The source value is empty.
	Available                CustomTime    `xml:"Available" json:"Available"`
	InsertDate               CustomTime    `xml:"insertDate" json:"insertDate"`
	Description_nl           string        `xml:"description_nl" json:"description_nl,omitempty"`
	Description_en           string        `xml:"description_en" json:"description_en,omitempty"`
	Description_fr           string        `xml:"description_fr" json:"description_fr,omitempty"`
	Description_de           string        `xml:"description_de" json:"description_de,omitempty"`
	Description_es           string        `xml:"description_es" json:"description_es,omitempty"`
	Description_it           string        `xml:"description_it" json:"description_it,omitempty"`
	Photos                   *PhotoAttr    `xml:"photos" json:"photos"`
	UpdatePhotos             int           `xml:"updatePhotos" json:"updatePhotos"`
	Pdf                      string        `xml:"pdf,omitempty" json:"pdf,omitempty"` // string - The source value is empty.
	ImageSlider              int           `xml:"imageslider" json:"imageslider"`
	Size_M2                  int           `xml:"size_m2" json:"size_m2"`
	NumberOfBathrooms        int           `xml:"numberOfBathrooms" json:"numberOfBathrooms"`
	ContractLength_Months    CustomInteger `xml:"contractLentgh_months" json:"contractLentgh_months,omitempty"`       // string - The source value is empty. And the source has a typo. (Lentgh -> Length)
	MinContractLength_Months CustomInteger `xml:"minContractLentgh_months" json:"minContractLentgh_months,omitempty"` // string - The source value is empty. And the source has a typo. (Lentgh -> Length)
	BuildYear                CustomInteger `xml:"buildYear" json:"buildYear,omitempty"`                               // string - The source value is empty.
	Parking                  int           `xml:"Parking" json:"Parking"`
	Bath                     int           `xml:"bath" json:"bath"`
	SeparateShower           int           `xml:"separateShower" json:"separateShower"`
	Lift                     int           `xml:"lift" json:"lift"`
	Garden                   int           `xml:"garden" json:"garden"`
	GardenLigging            string        `xml:"gardenLigging" json:"gardenLigging,omitempty"` // string - The source value is empty.
	GardenSizeM2             int           `xml:"gardenSizeM2" json:"gardenSizeM2"`
	RoofTerrass              int           `xml:"roofTerrass" json:"roofTerrass"`
	RoofTerrassLigging       string        `xml:"roofTerrassLigging" json:"roofTerrassLigging,omitempty"` // string - The source value is empty.
	RoofTerrassSizeM2        int           `xml:"roofTerrassSizeM2" json:"roofTerrassSizeM2"`
	Balcony                  int           `xml:"balcony" json:"balcony"`
	BalconyLigging           string        `xml:"BalconyLigging" json:"BalconyLigging,omitempty"` // string - The source value is empty.
	BalconySizeM2            int           `xml:"BalconySizeM2" json:"BalconySizeM2"`
	SwimmingPool             int           `xml:"swimmingPool" json:"swimmingPool"`
	AirConditioning          int           `xml:"airConditioning" json:"airConditioning"`
	FirePlace                int           `xml:"firePlace" json:"firePlace"`
	Garage                   int           `xml:"garage" json:"garage"`
	Cellar                   int           `xml:"cellar" json:"cellar"`
	PublicTransportQualityId int           `xml:"publicTransportQualityID" json:"publicTransportQualityID"`
	ShowhouseNumber          int           `xml:"showhouseNumber" json:"showhouseNumber"`
	GroundFloor              int           `xml:"groundFloor" json:"groundFloor"`
	FloorQuality             string        `xml:"floorQuality" json:"floorQuality"`
	RentIncluded             string        `xml:"rentIncluded" json:"rentIncluded"`
	Huislijn                 int           `xml:"huislijn" json:"huislijn"`
	Houselink                int           `xml:"houselink" json:"houselink"`
	Properazzi               int           `xml:"properazzi" json:"properazzi"`
	Enormo                   int           `xml:"enormo" json:"enormo"`
	Woningennet              int           `xml:"woningennet" json:"woningennet"`
	Kamerlink                int           `xml:"kamerlink" json:"kamerlink"`
	Expatica                 int           `xml:"expatica" json:"expatica"`
	Brixter                  int           `xml:"brixter" json:"brixter"`
	Jaap                     int           `xml:"jaap" json:"jaap"`
	Huisvoormij              int           `xml:"huisvoormij" json:"huisvoormij"`
	Huizenzoeker             int           `xml:"huizenzoeker" json:"huizenzoeker"`
	Kamernl                  int           `xml:"kamernl" json:"kamernl"`
	Speurders                int           `xml:"speurders" json:"speurders"`
	Hierismijnhuis           int           `xml:"hierismijnhuis" json:"hierismijnhuis"`
	Hierismijnhuiskamer      int           `xml:"hierismijnhuiskamer" json:"hierismijnhuiskamer"`
	Mansion                  int           `xml:"mansion" json:"mansion"`
	Mitula                   int           `xml:"mitula" json:"mitula"`
	Huurprofielen            int           `xml:"huurprofielen" json:"huurprofielen"`
	Huurstunt                int           `xml:"huurstunt" json:"huurstunt"`
	Iamexpat                 int           `xml:"iamexpat" json:"iamexpat"`
}

type PhotoAttr struct {
	Photo []string `xml:"photo"`
}

// parse date format "yyyy-mm-dd"
func (c *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(SHORT_DATE_FORM, v)
	if err != nil {
		return err
	}
	*c = CustomTime{parse}
	return nil
}

// parse int from cdata or empty value
func (i *CustomInteger) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var convert int
	var err error
	var v string
	d.DecodeElement(&v, &start)
	if len(v) > 0 {
		numStr := strings.TrimSpace(strings.Trim(v, ">")) // house number field has ">" letter
		convert, err = strconv.Atoi(numStr)
		if err != nil {
			return err
		}
	}

	*i = CustomInteger(convert)
	return nil
}

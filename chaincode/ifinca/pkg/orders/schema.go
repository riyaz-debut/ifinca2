// Package orders Related functions
package orders

import (
	"github.com/chaincode/ifinca/pkg/core/utils"
)

// Order Structure
type Order struct {
	DataInputs           DataInputs  `json:"data_inputs"`
	OrderNo              string      `json:"order_no"`
	MainQuantity         float64     `json:"main_quantity"`
	MainBaseUnit         string      `json:"main_base_unit"`
	Quantity             float64     `json:"quantity"`
	BaseUnit             string      `json:"base_unit"`
	PriceUnit            string      `json:"price_unit"`
	QrCode               string      `json:"qr_code"`
	AcceptedQuantity     float64     `json:"accepted_quantity"`
	Elevation            string      `json:"elevation"`
	DeliveryDate         int64       `json:"delivery_date"`
	ScreenSize           string      `json:"screen_size"`
	Price                int64       `json:"price"`
	MajorDefects         string      `json:"major_defects"`
	IfincaBonus          int64       `json:"ifinca_bonus"`
	SecondaryDefects     string      `json:"secondary_defects"`
	CupScore             int64       `json:"cup_score"`
	Country              string      `json:"country"`
	Moisture             string      `json:"moisture"`
	Farm                 string      `json:"farm"`
	SampleRequest        string      `json:"sample_request"`
	ImporterDeliveryDate int64       `json:"importer_delivery_date"`
	ExporterDeliveryDate int64       `json:"exporter_delivery_date"`
	RoasterDeliveryDate  int64       `json:"roaster_delivery_date"`
	CafeDeliveryDate     int64       `json:"cafe_delivery_date"`
	Region               string      `json:"region"`
	Process              string      `json:"process"`
	Variety              string      `json:"variety"`
	Certificates         string      `json:"certificates"`
	Status               int64       `json:"status"`
	CafeStores           []CafeStore `json:"cafe_stores"`
	Roasters             []CafeStore `json:"roasters"`
	Importers            []CafeStore `json:"importers"`
	Coops                []CafeStore `json:"coops"`
	DocType              string      `json:"doc_type"`
	utils.MetaData
}

// CafeStore Struct
type CafeStore struct {
	Address     Address `json:"address"`
	Name        string  `json:"name"`
	CountryCode string  `json:"country_code"`
	Phone       string  `json:"phone"`
	ContactName string  `json:"contact_name"`
	Status      int64   `json:"status"`
	ID          string  `json:"_id"`
}

// DataInputs Struct
type DataInputs struct {
	MillCost      int64 `json:"mill_cost"`
	ExporterCost  int64 `json:"exporter_cost"`
	ImporterCost  int64 `json:"importer_cost"`
	RoasterCost   int64 `json:"roaster_cost"`
	CafeCost      int64 `json:"cafe_cost"`
	TrueCostPerLB int64 `json:"true_cost_per_lb"`
	Status        int64 `json:"status"`
}

// SubOrders Structure
type SubOrders struct {
	SubOrders []SubOrder `json:"sub_orders"`
	utils.MetaData
}

// SubOrder Structure
type SubOrder struct {
	OrderNo                 string     `json:"order_no"`
	Supplier                Vendor     `json:"supplier"`
	DataPoints              DataPoints `json:"data_points"`
	Quantity                float64    `json:"quantity"`
	AcceptedQuantity        float64    `json:"accepted_quantity"`
	FilledQuantity          float64    `json:"filled_quantity"`
	DeclinedDatapointsCount int64      `json:"declined_datapoints_count"`
	CupScore                int64      `json:"cup_score"`
	MillProcess             string     `json:"mill_process"`
	AdditionalNotes         string     `json:"additional_notes"`
	Status                  int64      `json:"status"`
	OrderID                 string     `json:"order_id"`
	ActionDate              string     `json:"action_date"`
	Vendors                 []Vendor   `json:"vendors"`
	DeliveryDate            int64      `json:"delivery_date"`
	DocType                 string     `json:"doc_type"`
	utils.MetaData
}

// DataPoints Struct
type DataPoints struct {
	RawWeight       int64  `json:"raw_weight"`
	WeightFactor    int64  `json:"weight_factor"`
	PricePaid       int64  `json:"price_paid"`
	Factor          string `json:"factor"`
	MoistureContent string `json:"moisture_content"`
	HarvestMonth    string `json:"harvest_month"`
	Reason          string `json:"reason"`
}

// Vendor Struct
type Vendor struct {
	Type        int64   `json:"type"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	ContactName string  `json:"contact_name"`
	CountryCode string  `json:"country_code"`
	Phone       string  `json:"phone"`
	ProfilePic  string  `json:"profile_pic"`
	ID          string  `json:"_id"`
	Address     Address `json:"address"`
}

// Address Struct
type Address struct {
	Line    string `json:"line"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Pincode string `json:"pincode"`
}

// Key Struct
type Key struct {
	Key string `json:"key"`
}

// OrderNo Struct
type OrderNo struct {
	OrderNo string `json:"order_no"`
}

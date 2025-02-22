package models

import (
	"encoding/json"
	"strconv"
	"strings"
)

// type EC2Price struct {
// 	Products map[string]EC2Product `json:"products"`
// }

type EC2Product struct {
	Sku           string         `json:"sku"`
	ProductFamily string         `json:"productFamily"`
	Attributes    AwsEc2Inctance `json:"attributes"`
}

type RDSProduct struct {
	Sku           string         `json:"sku"`
	ProductFamily string         `json:"productFamily"`
	Attributes    AwsRdsInstance `json:"attributes"`
}

type PriceDimensions map[string]struct {
	Sku             string                      `json:"sku"`
	PriceDimensions map[string]DimensionDetails `json:"priceDimensions"`
	TermAttributes  TermAttributes              `json:"termAttributes"`
}

type TermAttributes struct {
	LeaseContractLength string `json:"LeaseContractLength"`
	PurchaseOption      string `json:"PurchaseOption"`
	OfferingClass       string `json:"OfferingClass"`
}

// DimensionDetails represents the details of each price dimension
type DimensionDetails struct {
	RateCode     string            `json:"rateCode"`
	Description  string            `json:"description"`
	BeginRange   string            `json:"beginRange"`
	EndRange     string            `json:"endRange"`
	Unit         string            `json:"unit"`
	PricePerUnit map[string]string `json:"pricePerUnit"`
	AppliesTo    []interface{}     `json:"appliesTo"`
}

type StringInt int
type StringFloat float64

// StringIntのカスタムUnmarshalJSONメソッド
func (si *StringInt) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		*si = 0
	} else {
		*si = StringInt(i)
	}
	return nil
}

// StringFloatのカスタムUnmarshalJSONメソッド
func (sf *StringFloat) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	str = strings.TrimSuffix(str, " GiB")
	str = strings.TrimSuffix(str, " GB")
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		*sf = 0
	} else {
		*sf = StringFloat(f)
	}
	return nil
}

func (si *AwsRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(si.String())
}

func (si *AwsRegion) UnmarshalJSON(data []byte) error {
	return si.UnmarshalText(data[1 : len(data)-1])
}

// func (de *DatabaseEngine) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(de.String())
// }

func (de *DatabaseEngine) UnmarshalJSON(data []byte) error {
	var err error
	switch string(data[1 : len(data)-1]) {
	case "Aurora MySQL":
		err = de.UnmarshalText([]byte("aurora-mysql"))
	case "Aurora PostgreSQL":
		err = de.UnmarshalText([]byte("aurora-postgresql"))
	case "MySQL":
		err = de.UnmarshalText([]byte("mysql"))
	case "PostgreSQL":
		*de = DatabaseEnginePostgresql
	default:
		err = ErrInvalidDatabaseEngine(string(data))
		return nil
	}
	return err
}

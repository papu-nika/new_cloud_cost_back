package lambda

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/papu-nika/new_cloud_cost_back/db"
	"github.com/papu-nika/new_cloud_cost_back/db/models"
)

type AwsLambdPriceFile struct {
	Products map[string]AwsLambdaProduct `json:"products"`
	Terms    struct {
		OnDemand map[string]map[string]AwsLambdaTerm `json:"OnDemand"`
	}
}

type AwsLambdaTerm struct {
	Sku             string `json:"sku"`
	PriceDimensions map[string]AwsLambdaPriceDimensions
}

type AwsLambdaPriceDimensions struct {
	PricePerUnit map[string]string `json:"pricePerUnit"`
	Unit         string            `json:"unit"`
	BeginRange   string            `json:"beginRange"`
	EndRange     string            `json:"endRange"`
}

type AwsLambdaProduct struct {
	Sku           string              `json:"sku"`
	ProductFamily string              `json:"productFamily"`
	Attributes    AwsLambdaAttributes `json:"attributes"`
}

type AwsLambdaAttributes struct {
	Servicecode      string `json:"servicecode"`
	Location         string `json:"location"`
	LocationType     string `json:"locationType"`
	Group            string `json:"group"`
	GroupDescription string `json:"groupDescription"`
	Usagetype        string `json:"usagetype"`
	Operation        string `json:"operation"`
	RegionCode       string `json:"regionCode"`
}

func Import() {
	db.DB.Where("1=1").Delete(&models.AwsLambda{}).Commit()

	lambdaPriceFile := &AwsLambdPriceFile{}
	f, err := os.ReadFile("price_file/lambda.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(f, lambdaPriceFile)

	var lambdaModels []models.AwsLambda
	for _, p := range lambdaPriceFile.Products {
		if p.Attributes.RegionCode == "" {
			continue
		}
		switch p.Attributes.Group {
		case "AWS-Lambda-Requests":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "x86",
				Type:         models.LambdaTypeRequests,
			})
		case "AWS-Lambda-Requests-ARM":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "arm",
				Type:         models.LambdaTypeRequests,
			})
		case "AWS-Lambda-Duration":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "x86",
				Type:         models.LambdaTypeDuration,
			})
		case "AWS-Lambda-Duration-ARM":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "arm",
				Type:         models.LambdaTypeDuration,
			})

		case "AWS-Lambda-Edge-Duration":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "x86",
				Type:         models.LambdaTypeEdgeDuration,
			})
		case "AWS-Lambda-Edge-Requests":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "x86",
				Type:         models.LambdaTypeEdgeRequest,
			})
		case "AWS-Lambda-Provisioned-Concurrency":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "x86",
				Type:         models.LambdaTypeProvisionedConcurrency,
			})
		case "AWS-Lambda-Provisioned-Concurrency-ARM":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "arm",
				Type:         models.LambdaTypeProvisionedConcurrency,
			})
		case "AWS-Lambda-Duration-Provisioned":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "x86",
				Type:         models.LambdaTypeProvisionedDuration,
			})
		case "AWS-Lambda-Duration-Provisioned-ARM":
			lambdaModels = append(lambdaModels, models.AwsLambda{
				ID:           p.Sku,
				Regioncode:   p.Attributes.RegionCode,
				Architecture: "arm",
				Type:         models.LambdaTypeProvisionedDuration,
			})
		}
	}

	db.DB.Create(&lambdaModels)

	for sku, term := range lambdaPriceFile.Terms.OnDemand {
		for _, t := range term {
			for _, pd := range t.PriceDimensions {
				var price float64
				if pd.BeginRange != "0" {
					continue
				}
				for _, p := range pd.PricePerUnit {
					price, _ = strconv.ParseFloat(p, 64)
				}
				db.DB.Where(
					"id = ?",
					sku,
				).Updates(
					&models.AwsLambda{
						Unit:  pd.Unit,
						Price: sql.NullFloat64{Float64: price, Valid: true},
					})
			}
		}

	}
}

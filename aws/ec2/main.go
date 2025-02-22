package ec2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/papu-nika/new_cloud_cost_back/aws/region"
	"github.com/papu-nika/new_cloud_cost_back/db"
	"github.com/papu-nika/new_cloud_cost_back/db/models"
	"gorm.io/gorm"
)

func Import() {
	f, err := os.Open("price_file/instance.json")
	// f, err := os.Open("price_file/test.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(bufio.NewReader(f))
	start := time.Now()
	db.DB.Where("1=1").Delete(&models.AwsEc2Inctance{}).Commit()

	log.Printf("Start: %s", start)
	log.Printf("ImportEC2Product: %s", time.Since(start))
	ImportEC2Product(decoder, db.DB)

	log.Printf("ImportEC2Terms: %s", time.Since(start))
	ImportEC2Terms(decoder, db.DB)
}

func ImportEC2Product(decoder *json.Decoder, db *gorm.DB) {
	var err error
	// まず、全体のJSONオブジェクトの開始部分を読み取る
	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err, "1")
	}

	// "products"フィールドまで読み込む
	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			log.Fatal(err, "2")
		}

		// "products"フィールドに到達したら、その内容をストリーム処理する
		if t == "products" {
			t, err = decoder.Token() // '{' トークンを消費
			if err != nil {
				log.Fatal(err, "3")
			}

			// "products"フィールドの内容をストリーム処理
			for decoder.More() {
				instances, isDone, err := getSliceAwsEc2Inctance(decoder, 2000)
				if err != nil {
					log.Fatal(err, "5")
				}

				db.Create(&instances)
				fmt.Println("product saved: ", len(instances))
				if isDone {
					return
				}
			}
		}
	}
}

func getSliceAwsEc2Inctance(decoder *json.Decoder, len int) (instances []models.AwsEc2Inctance, done bool, err error) {
	var awsInstances []models.AwsEc2Inctance

	for i := 0; i < len; {
		var product models.EC2Product
		// キーを読み込む
		if !decoder.More() {
			break
		}

		_, err = decoder.Token()
		if err != nil {
			log.Fatal(err, "4")
		}

		// 値を読み込む
		err = decoder.Decode(&product)
		if err != nil {
			// log.Println(err, "3aaaa")　// 対象外のリージョンでエラーが発生するためコメントアウト
			continue
		}

		if product.Attributes.Capacitystatus != "Used" ||
			product.ProductFamily != "Compute Instance" ||
			product.Attributes.Tenancy != "Shared" ||
			product.Attributes.Preinstalledsw != "NA" ||
			product.Attributes.Licensemodel != "No License required" {
			continue
		}
		if !slices.Contains(region.ENABLED_REGIONS, product.Attributes.Regioncode.String()) {
			continue
		}

		product.Attributes.ID = product.Sku
		if os := exchangeOS(string(product.Attributes.Operatingsystem.String())); os == 0 {
			i++
			continue
		} else {
			product.Attributes.Operatingsystem = os
			awsInstances = append(awsInstances, product.Attributes)
			i++
		}
	}
	return awsInstances, false, nil
}

func ImportEC2Terms(decoder *json.Decoder, db *gorm.DB) {
	var err error
	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err, "1")
	}

	// "products"フィールドまで読み込む
	for {
		t, err := decoder.Token()
		if err != nil {
			log.Fatal(err, "2")
		}

		// "terms"
		if t == "terms" {
			t, err = decoder.Token() // '{' トークンを消費
			t, err = decoder.Token() // 'OnDemand' トークンを消費
			t, err = decoder.Token() // 'OnDemand' トークンを消費

			var wg sync.WaitGroup
			priceChan := make(chan []models.PriceDimensions, 50)
			for i := 0; i < 50; i++ {
				wg.Add(1)
				go proccessPrices(priceChan, db, &wg)
			}

			for decoder.More() {
				if !decoder.More() {
					break
				}
				prices, err := getPriceSlice(decoder, 5000)
				if err != nil {
					log.Fatal(err, "6")
				}
				priceChan <- prices
			}
			close(priceChan)
			wg.Wait()
		}
	}
}

func proccessPrices(priceChan <-chan []models.PriceDimensions, db *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	for prices := range priceChan {
		tx := db.Begin()
		if err := tx.Error; err != nil {
			panic(err)
		}
		for _, price := range prices {
			for sku, p := range price {
				for _, detail := range p.PriceDimensions {
					// Pass Partial Upfront
					if p.TermAttributes.PurchaseOption == "Partial Upfront" || p.TermAttributes.PurchaseOption == "No Upfront" {
						continue
						// 1Year and and Standard and AllUpfront
					} else if (p.TermAttributes.LeaseContractLength == "1yr" || p.TermAttributes.LeaseContractLength == "1 yr") &&
						(p.TermAttributes.PurchaseOption == "AllUpfront" || p.TermAttributes.PurchaseOption == "All Upfront") &&
						(p.TermAttributes.OfferingClass == "standard") {
						if oneYearPrice, err := strconv.ParseFloat(detail.PricePerUnit["USD"], 64); err != nil {
							log.Fatal(err, "8")
						} else if oneYearPrice == 0 {
							continue
						} else {
							if err := tx.Where("id = ?", p.Sku).
								Updates(models.AwsEc2Inctance{
									OneYearReservedStandardPrice: models.StringFloat(oneYearPrice),
								}).Error; err != nil {
								panic(err)
							}
						}
						// 3Year and and Sta	ndard and AllUpfront
					} else if (p.TermAttributes.LeaseContractLength == "3yr" || p.TermAttributes.LeaseContractLength == "3 yr") &&
						(p.TermAttributes.PurchaseOption == "AllUpfront" || p.TermAttributes.PurchaseOption == "All Upfront") &&
						(p.TermAttributes.OfferingClass == "standard") {
						if threeYearPrice, err := strconv.ParseFloat(detail.PricePerUnit["USD"], 64); err != nil {
							log.Fatal(err, "8")
						} else if threeYearPrice == 0 {
							continue
						} else {
							if err := tx.Where("id = ?", p.Sku).
								Updates(models.AwsEc2Inctance{
									ThreeYearReservedStandardPrice: models.StringFloat(threeYearPrice),
								}).Error; err != nil {
								panic(err)
							}
						}
						// 1Year and and Convertible and AllUpfront
					} else if (p.TermAttributes.LeaseContractLength == "1yr" || p.TermAttributes.LeaseContractLength == "1 yr") &&
						(p.TermAttributes.PurchaseOption == "AllUpfront" || p.TermAttributes.PurchaseOption == "All Upfront") &&
						(p.TermAttributes.OfferingClass == "convertible") {
						if oneYearConvertiblePrice, err := strconv.ParseFloat(detail.PricePerUnit["USD"], 64); err != nil {
							log.Fatal(err, "8")
						} else if oneYearConvertiblePrice == 0 {
							continue
						} else {
							if err := tx.Where("id = ?", p.Sku).
								Updates(models.AwsEc2Inctance{
									OneYearReservedConvertiblePrice: models.StringFloat(oneYearConvertiblePrice),
								}).Error; err != nil {
								panic(err)
							}
						}
					} else if (p.TermAttributes.LeaseContractLength == "3yr" || p.TermAttributes.LeaseContractLength == "3 yr") &&
						(p.TermAttributes.PurchaseOption == "AllUpfront" || p.TermAttributes.PurchaseOption == "All Upfront") &&
						(p.TermAttributes.OfferingClass == "convertible") {
						if threeYearConvertiblePrice, err := strconv.ParseFloat(detail.PricePerUnit["USD"], 64); err != nil {
							log.Fatal(err, "8")
						} else if threeYearConvertiblePrice == 0 {
							continue
						} else {
							if err := tx.Where("id = ?", p.Sku).
								Updates(models.AwsEc2Inctance{
									ThreeYearReservedConvertiblePrice: models.StringFloat(threeYearConvertiblePrice),
								}).Error; err != nil {
								panic(err)
							}
						}
					} else {
						ondemandPrice, err := strconv.ParseFloat(detail.PricePerUnit["USD"], 64)
						if err != nil {
							log.Fatal(err, "8")
						}
						if err := tx.Where("id = ?", sku[0:strings.Index(sku, ".")]).
							Updates(models.AwsEc2Inctance{
								Ondemandprice: models.StringFloat(ondemandPrice),
							}).Error; err != nil {
							panic(err)
						}
					}

				}
			}
		}
		tx.Commit()
		fmt.Println("price saved: ", len(prices))
	}
}

var isPassdReServe bool

func getPriceSlice(decoder *json.Decoder, length int) ([]models.PriceDimensions, error) {
	// var err error
	var isCheckPoint bool
	var prices []models.PriceDimensions
	for i := 0; i < length; i++ {
		// キーを読み込む
		if !decoder.More() {
			break
		}

		t, err := decoder.Token()
		if err != nil {
			log.Fatal(err, "4")
		}
		if fmt.Sprintf("%s", t) == "H99AJXVC6HMMX7HD" {
			fmt.Println("found H99AJXVC6HMMX7HD")
			isCheckPoint = true
		}

		// 値を読み込む
		var price models.PriceDimensions
		err = decoder.Decode(&price)
		if err != nil {
			log.Fatal(err, "7")
		}

		prices = append(prices, price)
		// Reservedの処理
		if isCheckPoint && !isPassdReServe {
			t, err = decoder.Token() // '{' トークンを消費
			fmt.Println(t)
			t, err = decoder.Token() // 'Reserved' トークンを消費
			fmt.Println(t)
			t, err = decoder.Token() // 'OnDemand' トークンを消費
			fmt.Println(t)
			isPassdReServe = true
			isCheckPoint = false
			return prices, nil
		}
	}
	return prices, nil
}

func exchangeOS(os string) models.Os {
	switch os {
	case "Linux":
		return models.OsLinux
	// case "RHEL":
	// 	return "rhel"
	case "Windows":
		return models.OsWindows
	// case "Ubuntu Pro":
	// 	return "ubuntu-pro"
	default:
		return 0
	}
}

package rds

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/papu-nika/new_cloud_cost_back/db"
	"github.com/papu-nika/new_cloud_cost_back/db/models"
	"gorm.io/gorm"
)

var (
	needDBEngine = []string{
		"Aurora MySQL",
		"Aurora PostgreSQL",
		"MariaDB",
		"MySQL",
		"Oracle",
		"PostgreSQL",
		"SQL Server",
	}
)

func Import() {
	f, err := os.Open("price_file/rds.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(bufio.NewReader(f))
	start := time.Now()
	db.DB.Where("1=1").Delete(&models.AwsRdsInstance{}).Commit()
	db.DB.Where("1=1").Delete(&models.AwsAuroraServerless{}).Commit()

	log.Printf("Start import RDS: %s", start)
	log.Printf("ImportRDSProduct: %s", time.Since(start))
	ImportRDSProduct(decoder, db.DB)

	log.Printf("ImportRDSTerms: %s", time.Since(start))
	ImportRDSTerms(decoder, db.DB)
}

func ImportRDSProduct(decoder *json.Decoder, db *gorm.DB) {
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
				instances, isDone, err := getSliceAwsRDSInctance(db, decoder, 2000)
				if err != nil {
					log.Fatal(err, "5")
				}

				if err := db.Create(&instances).Error; err != nil {
					panic(err)
				}
				fmt.Println("product saved: ", len(instances))
				if isDone {
					return
				}
			}
		}
	}
}

func getSliceAwsRDSInctance(db *gorm.DB, decoder *json.Decoder, len int) (instances []models.AwsRdsInstance, done bool, err error) {
	var awsInstances []models.AwsRdsInstance

	for i := 0; i < len; {
		var product models.RDSProduct
		// キーを読み込む
		if !decoder.More() {
			break
		}

		_, err := decoder.Token()
		if err != nil {
			log.Fatal(err, "4")
		}

		// 値を読み込む
		err = decoder.Decode(&product)
		if err != nil {
			log.Println(err, "3")
		}

		if product.ProductFamily != "Database Instance" {
			if product.ProductFamily == "ServerlessV2" {
				var serverlessV2 models.AwsAuroraServerless
				serverlessV2.ID = product.Sku
				serverlessV2.Regioncode = product.Attributes.Regioncode
				serverlessV2.Databaseengine = product.Attributes.Databaseengine
				if product.Attributes.Storage == "" {
					serverlessV2.Isauroraiooptimizationmode = false
				} else {
					serverlessV2.Isauroraiooptimizationmode = true
				}
				dbEngine := exchangeDBEngine(product.Attributes.Databaseengine)
				if dbEngine == "" {
					continue
				}
				serverlessV2.Databaseengine = dbEngine

				db.Create(&serverlessV2)
			}
			continue
		}
		dbEngine := exchangeDBEngine(product.Attributes.Databaseengine)
		if dbEngine == "" {
			continue
		}
		product.Attributes.Databaseengine = dbEngine
		product.Attributes.ID = product.Sku
		awsInstances = append(awsInstances, product.Attributes)
		i++
	}
	return awsInstances, false, nil
}

func ImportRDSTerms(decoder *json.Decoder, db *gorm.DB) {
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

			// "terms"フィールドの内容をストリーム処理
			var tx *gorm.DB
			for decoder.More() {
				// キーを読み込む
				if !decoder.More() {
					break
				}

				var prices []models.PriceDimensions
				prices, err = getPriceSlice(decoder, 5000)
				if err != nil {
					log.Fatal(err, "5")
				}
				tx = db.Begin()
				if tx.Error != nil {
					panic(tx.Error)
				}
				for _, price := range prices {
					for sku, p := range price {
						for _, detail := range p.PriceDimensions {
							ondemandPrice, err := strconv.ParseFloat(detail.PricePerUnit["USD"], 64)
							if err != nil {
								log.Fatal(err, "8")
							}
							ondemandPriceNull := sql.NullFloat64{Float64: ondemandPrice, Valid: true}

							if strings.Contains(detail.Description, "Serverless v2") {
								if err := tx.Where("id = ?", sku[0:strings.Index(sku, ".")]).
									Updates(models.AwsAuroraServerless{
										Ondemandprice: ondemandPriceNull,
									}).Error; err != nil {
									panic(err)
								}
							} else {
								if err := tx.Where("id = ?", sku[0:strings.Index(sku, ".")]).
									Updates(models.AwsRdsInstance{
										Ondemandprice: ondemandPriceNull,
									}).Error; err != nil {
									panic(err)
								}
							}
						}
					}
				}

				if err := tx.Commit().Error; err != nil {
					panic(err)
				}
				fmt.Println("updated: ", len(prices))

			}
			tx.Commit()
		}
	}
}

func getPriceSlice(decoder *json.Decoder, len int) ([]models.PriceDimensions, error) {
	var err error
	var prices []models.PriceDimensions
	for i := 0; i < len; i++ {
		// キーを読み込む
		if !decoder.More() {
			break
		}

		_, err = decoder.Token()
		if err != nil {
			log.Fatal(err, "4")
		}

		// 値を読み込む
		var price models.PriceDimensions
		err = decoder.Decode(&price)
		if err != nil {
			log.Fatal(err, "7")
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func exchangeDBEngine(os string) string {
	switch os {

	case "Aurora MySQL":
		return "aurora-mysql"
	case "Aurora PostgreSQL":
		return "aurora-postgresql"
	case "MariaDB":
		return "mariadb"
	case "MySQL":
		return "mysql"
	case "Oracle":
		return "oracle"
	case "PostgreSQL":
		return "postgresql"
	case "SQL Server":
		return "sqlserver"
	default:
		return ""
	}
}

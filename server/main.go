package server

import (
	"fmt"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/papu-nika/new_cloud_cost_back/api"
	"github.com/papu-nika/new_cloud_cost_back/db"
	"github.com/papu-nika/new_cloud_cost_back/db/models"
)

// Server構造体の定義
type Server struct{}

func (s *Server) GetAwsEc2Instances(c *gin.Context, params api.GetAwsEc2InstancesParams) {
	if err := s.Validate(c, params); err != nil {
		var resp api.GetAwsEc2Instances400Response
		if err := resp.VisitGetAwsEc2InstancesResponse(c.Writer); err != nil {
			c.JSON(500, gin.H{"error": "failed to get response"})
		} else {
			c.JSON(400, resp)
		}
		slog.Debug(err.Error())
		return
	}

	var instaneModel models.AwsEc2Inctance
	if params.Os != nil {
		var os models.Os
		os.UnmarshalText([]byte(params.Os.Os))
		instaneModel.Operatingsystem = os
	} else {
		// Default linux
		instaneModel.Operatingsystem = models.OsLinux
	}

	var awsRegion models.AwsRegion
	awsRegion.UnmarshalText([]byte(params.Region.Region))
	if awsRegion != 0 {
		instaneModel.Regioncode = awsRegion
	} else {
		// Default ap-northeast-1
		instaneModel.Regioncode = models.AwsRegionApNortheast1
	}

	if params.Instancetype != nil {
		instaneModel.Instancetype = *params.Instancetype
	}

	fmt.Println("#####", instaneModel.Regioncode, params.Region)
	var awsEc2Instances api.GetAwsEc2Instances200JSONResponse
	db.DB.Model(instaneModel).Where(&instaneModel).Order("vcpu, memory").Find(&awsEc2Instances)

	awsEc2Instances.VisitGetAwsEc2InstancesResponse(c.Writer)
}

func (s *Server) GetAwsEc2InstancesInstanceSku(c *gin.Context, instanceSKU string) {
	var awsEc2Instance api.GetAwsEc2InstancesInstanceSku200JSONResponse

	var instaneModel models.AwsEc2Inctance
	instaneModel.ID = instanceSKU
	db.DB.Debug().Model(instaneModel).Where(&instaneModel).First(&awsEc2Instance)
	if awsEc2Instance.Id == nil || *awsEc2Instance.Id == "" {
		resp := api.GetAwsEc2InstancesInstanceSku404Response{}.VisitGetAwsEc2InstancesInstanceSkuResponse(c.Writer)
		c.JSON(404, resp)
		return
	}

	awsEc2Instance.VisitGetAwsEc2InstancesInstanceSkuResponse(c.Writer)
}

func (s *Server) GetAwsRdsInstances(c *gin.Context, params api.GetAwsRdsInstancesParams) {
	if err := s.Validate(c, params); err != nil {
		var resp api.GetAwsRdsInstances400Response
		if err := resp.VisitGetAwsRdsInstancesResponse(c.Writer); err != nil {
			c.JSON(500, gin.H{"error": "failed to get response"})
		} else {
			c.JSON(400, resp)
		}
		slog.Debug(err.Error())
		return
	}

	var instaneModel models.AwsRdsInstance
	if params.Engine != nil {
		var engine models.DatabaseEngine
		engine.Scan(params.Engine)
		instaneModel.Databaseengine = engine
	} else {
		// Default linux
		instaneModel.Databaseengine = models.DatabaseEngineAuroraPostgresql
	}

	var awsRegion models.AwsRegion
	awsRegion.Scan(params.Region)
	if awsRegion != 0 {
		instaneModel.Regioncode = awsRegion
	} else {
		// Default ap-northeast-1
		instaneModel.Regioncode = models.AwsRegionApNortheast1
	}

	if params.Instancetype != nil {
		instaneModel.Instancetype = *params.Instancetype

	}

	var awsEc2Instances api.GetAwsRdsInstances200JSONResponse
	db.DB.Model(instaneModel).Where(&instaneModel).Find(&awsEc2Instances)

	awsEc2Instances.VisitGetAwsRdsInstancesResponse(c.Writer)
}

func (s *Server) GetAwsRdsInstancesInstanceSku(c *gin.Context, instanceSKU string) {
	var awsEc2Instance api.GetAwsRdsInstancesInstanceSku200JSONResponse

	var instaneModel models.AwsRdsInstance
	instaneModel.ID = instanceSKU
	db.DB.Debug().Model(instaneModel).Where(&instaneModel).First(&awsEc2Instance)
	if awsEc2Instance.Id == nil || *awsEc2Instance.Id == "" {
		resp := api.GetAwsEc2InstancesInstanceSku404Response{}.VisitGetAwsEc2InstancesInstanceSkuResponse(c.Writer)
		c.JSON(404, resp)
		return
	}

	awsEc2Instance.VisitGetAwsRdsInstancesInstanceSkuResponse(c.Writer)
}

func (s *Server) Validate(c *gin.Context, param interface{}) error {
	if err := validator.New().Struct(param); err != nil {
		return fmt.Errorf("not avaraible parameter: %s", param)
	}
	return nil
}

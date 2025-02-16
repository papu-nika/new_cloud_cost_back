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
	instaneModel.Operatingsystem = string(params.Os.Os)
	instaneModel.Regioncode = string(params.Region.Region)

	var awsEc2Instances api.GetAwsEc2Instances200JSONResponse
	awsEc2Instances.VisitGetAwsEc2InstancesResponse(c.Writer)

	db.DB.Model(instaneModel).Where(&instaneModel).Find(&awsEc2Instances)
	c.JSON(200, awsEc2Instances)
	return
}

func (s *Server) Validate(c *gin.Context, param interface{}) error {
	if err := validator.New().Struct(param); err != nil {
		return fmt.Errorf("not avaraible parameter: %s", param)
	}
	return nil
}

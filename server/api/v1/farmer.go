package v1

import (
	bc "application/blockchain"
	"application/model"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)


func CreateCrops(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var cropsJson model.Crops
	err := c.ShouldBindJSON(&cropsJson)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.CreateCrops(cropsJson.CropsId, cropsJson.CropsName, cropsJson.Address, cropsJson.RegisterTime,
		cropsJson.Year, cropsJson.FarmerName, cropsJson.FarmerID, cropsJson.FarmerTel, cropsJson.FertilizerName,
		cropsJson.PlatMode, cropsJson.BaggingStatus, cropsJson.GrowSeedlingsCycle, cropsJson.IrrigationCycle,
		cropsJson.ApplyFertilizerCycle, cropsJson.WeedCycle, cropsJson.Remarks)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
}
func RecordCropsGrow(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	var cropsGrowInfoJson model.CropsGrowInfo
	err := c.ShouldBindJSON(&cropsGrowInfoJson)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	jsonData, err := bc.RecordCropsGrow(cropsGrowInfoJson.CropsGrowId, cropsGrowInfoJson.CropsBakId, cropsGrowInfoJson.RecordTime,
		cropsGrowInfoJson.CropsGrowPhotoUrl, cropsGrowInfoJson.Temperature, cropsGrowInfoJson.GrowStatus,
		cropsGrowInfoJson.WaterContent, cropsGrowInfoJson.IlluminationStatus, cropsGrowInfoJson.Remarks)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
}
func QueryCropsById(c *gin.Context) {
	resp := app.Response{
		Code: "200",
		Data: nil,
	}

	defer app.ResponseFunc(c, &resp)

	CropsID, ok := c.GetQuery("crops_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("crops_id error")
		return
	}
	jsonData, err := bc.QueryCropsById(CropsID)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
}
func QueryCropsProcessByCropsId(c *gin.Context) {
	resp := app.Response{Code: "200", Msg: "OK"}

	defer app.ResponseFunc(c, &resp)

	CropsID, ok := c.GetQuery("crops_id")
	if !ok {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("crops_id error")
		return
	}
	jsonData, err := bc.QueryCropsProcessByCropsId(CropsID)
	if err != nil {
		resp.Code = "101"
		resp.Msg = fmt.Sprintf("query error: %s", err.Error())
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = "101"
		resp.Msg = err.Error()
		return
	}
}

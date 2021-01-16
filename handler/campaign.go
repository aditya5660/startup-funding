package handler

import (
	"net/http"
	"startup-funding/campaign"
	"startup-funding/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get params di handler
// handler to services
// services to repository
// services menentukan repository yang di call
// repository acccess to db

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

// api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// get user id
	userID, _ := strconv.Atoi(c.Query("user_id"))
	// get data
	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Error to get campaings", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// format response
	formatter := campaign.FormatCampaigns(campaigns)
	// api response
	response := helper.APIResponse("list of campaigns", http.StatusOK, "success", formatter)
	// return json
	c.JSON(http.StatusOK, response)
}

// api/v1/campaigns/:id
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// get input
	var input campaign.GetCampaignDetailInput
	// err input
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// get data
	campaign, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// format response
	formatter := campaign.FormatCampaignDetail(campaign)
	// api response
	response := helper.APIResponse("Get Campaign successfuly!", http.StatusOK, "success", formatter)
	// return json
	c.JSON(http.StatusOK, response)

}

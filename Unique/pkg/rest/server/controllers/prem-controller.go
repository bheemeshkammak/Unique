package controllers

import (
	"github.com/bheemeshkammak/Unique/unique/pkg/rest/server/models"
	"github.com/bheemeshkammak/Unique/unique/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PremController struct {
	premService *services.PremService
}

func NewPremController() (*PremController, error) {
	premService, err := services.NewPremService()
	if err != nil {
		return nil, err
	}
	return &PremController{
		premService: premService,
	}, nil
}

func (premController *PremController) CreatePrem(context *gin.Context) {
	// validate input
	var input models.Prem
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger prem creation
	if _, err := premController.premService.CreatePrem(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Prem created successfully"})
}

func (premController *PremController) UpdatePrem(context *gin.Context) {
	// validate input
	var input models.Prem
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger prem update
	if _, err := premController.premService.UpdatePrem(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Prem updated successfully"})
}

func (premController *PremController) FetchPrem(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger prem fetching
	prem, err := premController.premService.GetPrem(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, prem)
}

func (premController *PremController) DeletePrem(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger prem deletion
	if err := premController.premService.DeletePrem(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Prem deleted successfully",
	})
}

func (premController *PremController) ListPrems(context *gin.Context) {
	// trigger all prems fetching
	prems, err := premController.premService.ListPrems()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, prems)
}

func (*PremController) PatchPrem(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*PremController) OptionsPrem(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*PremController) HeadPrem(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}

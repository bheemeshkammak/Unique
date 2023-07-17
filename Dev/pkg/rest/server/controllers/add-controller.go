package controllers

import (
	"github.com/bheemeshkammak/Unique/dev/pkg/rest/server/models"
	"github.com/bheemeshkammak/Unique/dev/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type AddController struct {
	addService *services.AddService
}

func NewAddController() (*AddController, error) {
	addService, err := services.NewAddService()
	if err != nil {
		return nil, err
	}
	return &AddController{
		addService: addService,
	}, nil
}

func (addController *AddController) CreateAdd(context *gin.Context) {
	// validate input
	var input models.Add
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger add creation
	if _, err := addController.addService.CreateAdd(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Add created successfully"})
}

func (addController *AddController) UpdateAdd(context *gin.Context) {
	// validate input
	var input models.Add
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

	// trigger add update
	if _, err := addController.addService.UpdateAdd(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Add updated successfully"})
}

func (addController *AddController) FetchAdd(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger add fetching
	add, err := addController.addService.GetAdd(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, add)
}

func (addController *AddController) DeleteAdd(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger add deletion
	if err := addController.addService.DeleteAdd(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Add deleted successfully",
	})
}

func (addController *AddController) ListAdds(context *gin.Context) {
	// trigger all adds fetching
	adds, err := addController.addService.ListAdds()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, adds)
}

func (*AddController) PatchAdd(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*AddController) OptionsAdd(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*AddController) HeadAdd(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}

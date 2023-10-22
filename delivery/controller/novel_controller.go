package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"enigmacamp.com/rest-api-novel/domain"
	"enigmacamp.com/rest-api-novel/service"
	"enigmacamp.com/rest-api-novel/utils/common"
	"enigmacamp.com/rest-api-novel/utils/exception"
	"github.com/gin-gonic/gin"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type NovelController struct {
	service service.NovelService
}

func (n *NovelController) GetListNovel(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Status:  http.StatusBadRequest,
			Message: exception.ErrInvalidPage.Error(),
		})
		return
	}

	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Status:  http.StatusBadRequest,
			Message: exception.ErrInvalidPerPage.Error(),
		})
		return
	}

	code := c.DefaultQuery("code", "")
	title := c.DefaultQuery("title", "")
	publisher := c.DefaultQuery("publisher", "")
	year := c.DefaultQuery("year", "")
	author := c.DefaultQuery("author", "")

	novels, totalData, err := n.service.GetAllNovelList(page, perPage, code, title, publisher, year, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	data := domain.Pagination{
		NovelList: novels,
		Page:      page,
		PerPage:   perPage,
		TotalPage: common.CountTotalPage(totalData, perPage),
		TotalData: totalData,
	}

	c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success Get List Novel",
		Data:    data,
	})
}

func (n *NovelController) GetDetailNovel(c *gin.Context) {
	id := c.Param("id")

	novel, err := n.service.GetNovelById(id)
	if err != nil {
		if strings.Contains(err.Error(), exception.ErrNotFound.Error()) {
			c.JSON(http.StatusNotFound, errorResponse{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success Get Detail Novel",
		Data:    novel,
	})
}

func (e *NovelController) CreateNovel(c *gin.Context) {
	var novel domain.Novel
	if err := c.ShouldBindJSON(&novel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	novel.Id = common.GenerateUUID()
	novel.CreatedAt = time.Now().String()
	err := e.service.CreateNovel(novel)
	if err != nil {
		if strings.Contains(err.Error(), exception.ErrCodeAlreadyExist.Error()) || strings.Contains(err.Error(), exception.ErrTitleAlreadyExist.Error()) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{
			Status:  http.StatusInternalServerError,
			Message: exception.ErrFailedCreate.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response{
		Status:  http.StatusCreated,
		Message: "Success Create Novel",
		Data:    novel,
	})
}

func (ctr *NovelController) UpdateNovel(c *gin.Context) {

	id := c.Param("id")

	novel := domain.Novel{}

	if err := c.ShouldBindJSON(&novel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})

		return
	}

	data, err := ctr.service.UpdateNovelByID(id, novel)

	if err != nil {
		if strings.Contains(err.Error(), exception.ErrNotFound.Error()) {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})

			return
		}

		if strings.Contains(err.Error(), exception.ErrCodeAlreadyExist.Error()) || strings.Contains(err.Error(), exception.ErrTitleAlreadyExist.Error()) {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})

			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Successfully Update Novel By ID " + id,
		Data:    data,
	})
}

func (ctr *NovelController) DeleteNovel(c *gin.Context) {
	id := c.Param("id")

	err := ctr.service.DeleteNovelByID(id)
	if err != nil {
		if strings.Contains(err.Error(), exception.ErrNotFound.Error()) {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})

			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Successfully Delete Novel By ID " + id,
	})
}

func NewNovelController(service service.NovelService) *NovelController {
	controller := &NovelController{
		service: service,
	}

	return controller
}

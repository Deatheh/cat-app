package handler

import (
	"net/http"
	"strconv"

	"github.com/Deatheh/cat-app"
	"github.com/gin-gonic/gin"
)

// @Summary Create cat list
// @Security ApiKeyAuth
// @Tags lists
// @Description create cat list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body cat.CatList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, "user id not find")
		return
	}

	var input cat.CatList

	if err := c.BindJSON(&input); err != nil {

		newErroreResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CatList.Create(userId, input)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []cat.CatList `json:"data"`
}

// @Summary Get List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} cat.Cat
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, "user id not find")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.CatList.GetById(userId, id)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, "user id not find")
		return
	}

	lists, err := h.services.CatList.GetAll(userId)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Update List
// @Security ApiKeyAuth
// @Tags lists
// @Description update cat list
// @ID update-list
// @Accept  json
// @Produce  json
// @Param input body cat.CatList true "list info"
// @Success 200 {integer} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [put]
func (h *Handler) updeteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, "user id not find")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input cat.UpdeteListInput
	if err := c.Bind(&input); err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.CatList.Update(userId, id, input); err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete List
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list by id
// @ID delete-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, "user id not find")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.CatList.Delete(userId, id)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

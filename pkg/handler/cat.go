package handler

import (
	"net/http"
	"strconv"

	"github.com/Deatheh/cat-app"
	"github.com/gin-gonic/gin"
)

// @Summary Create Cat
// @Security ApiKeyAuth
// @Tags cats
// @Description create cat
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body cat.Cat true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/cats [post]
func (h *Handler) createCat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, "user id not find")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input cat.Cat
	if err := c.BindJSON(&input); err != nil {
		newErroreResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Cat.Create(userId, listId, input)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get Cat By Id
// @Security ApiKeyAuth
// @Tags cats
// @Description get cat by id
// @ID get-cat-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} cat.Cat
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/cats/:cat_id [get]
func (h *Handler) getCatById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.Cat.GetById(userId, listId, itemId)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Get All Cats
// @Security ApiKeyAuth
// @Tags cats
// @Description get all cats
// @ID get-all-cats
// @Accept  json
// @Produce  json
// @Success 200 {object} []cat.Cat
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/cats [get]
func (h *Handler) getAllCat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.Cat.GetAll(userId, listId)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Update Cat
// @Security ApiKeyAuth
// @Tags cats
// @Description update cat
// @ID update-cat
// @Accept  json
// @Produce  json
// @Param input body cat.Cat true "list info"
// @Success 200 {integer} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/cats/:cat_id [put]
func (h *Handler) updeteCat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	id, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input cat.UpdateCatInput
	if err := c.BindJSON(&input); err != nil {
		newErroreResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Cat.Update(userId, listId, id, input); err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Cat
// @Security ApiKeyAuth
// @Tags cats
// @Description delete cat by id
// @ID delete-cats
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/cats/:cat_id [delete]
func (h *Handler) deleteCat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		newErroreResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.Cat.Delete(userId, itemId)
	if err != nil {
		newErroreResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

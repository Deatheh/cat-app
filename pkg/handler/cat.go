package handler

import (
	"net/http"
	"strconv"

	"github.com/Deatheh/cat-app"
	"github.com/gin-gonic/gin"
)

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

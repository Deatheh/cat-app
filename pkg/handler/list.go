package handler

import (
	"net/http"
	"strconv"

	"github.com/Deatheh/cat-app"
	"github.com/gin-gonic/gin"
)

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

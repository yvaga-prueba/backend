package handle

import (
	"net/http"
	"strconv"

	"core/domain/model"
	"core/domain/repo"
	"github.com/labstack/echo/v4"
)

type ContactMessageHandler struct {
	repo repo.ContactMessageRepository
}

func NewContactMessageHandler(r repo.ContactMessageRepository) *ContactMessageHandler {
	return &ContactMessageHandler{repo: r}
}

func (h *ContactMessageHandler) Create(c echo.Context) error {
	var req struct {
		Name        string `json:"nombre"`
		Email       string `json:"email"`
		Message     string `json:"mensaje"`
		OrderNumber string `json:"numeroOrden"`
		UserID      *int   `json:"user_id"` // <--- CAMBIADO A *int
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	msg := &model.ContactMessage{
		Name:        req.Name,
		Email:       req.Email,
		Message:     req.Message,
		OrderNumber: req.OrderNumber,
		UserID:      req.UserID,
	}

	if err := h.repo.Save(c.Request().Context(), msg); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al guardar el mensaje"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Mensaje enviado correctamente",
		"id":      msg.ID,
	})
}

func (h *ContactMessageHandler) GetAdminMessages(c echo.Context) error {
	orderNumber := c.QueryParam("order_number")

	mensajes, err := h.repo.GetAll(c.Request().Context(), orderNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al obtener mensajes"})
	}

	return c.JSON(http.StatusOK, mensajes)
}

func (h *ContactMessageHandler) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // <--- Usamos Atoi para int
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Estado inválido"})
	}

	if err := h.repo.UpdateStatus(c.Request().Context(), id, req.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar estado"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Estado actualizado"})
}
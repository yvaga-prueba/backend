package handle

import (
	"core/domain/model"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SizeGuideHandler struct {
	DB *sql.DB
}

// Crea una regla nueva en la BD
func (h *SizeGuideHandler) CreateSizeGuide(c echo.Context) error {
	var guide model.SizeGuide
	if err := c.Bind(&guide); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "datos inválidos"})
	}

	query := `INSERT INTO size_guides (
		category, fit_type, size, min_weight, max_weight, min_height, max_height, 
		chest_cm, waist_cm, hip_cm, length_cm
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	result, err := h.DB.ExecContext(c.Request().Context(), query, 
		guide.Category, guide.FitType, guide.Size, guide.MinWeight, guide.MaxWeight, 
		guide.MinHeight, guide.MaxHeight, guide.ChestCm, guide.WaistCm, 
		guide.HipCm, guide.LengthCm)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no se pudo guardar la guía"})
	}

	id, _ := result.LastInsertId()
	guide.ID = id

	return c.JSON(http.StatusCreated, guide)
}

// Trae las guías filtradas por categoría y opcionalmente por fit_type (moldería)
func (h *SizeGuideHandler) GetGuidesByCategory(c echo.Context) error {
	category := c.Param("category")
	fitType := c.QueryParam("fit_type") 
	
	var query string
	var args []interface{}

	if fitType != "" {
		query = `SELECT id, category, fit_type, size, min_weight, max_weight, min_height, max_height, 
				 chest_cm, waist_cm, hip_cm, length_cm, created_at, updated_at 
				 FROM size_guides WHERE category = ? AND fit_type = ?`
		args = append(args, category, fitType)
	} else {
		query = `SELECT id, category, fit_type, size, min_weight, max_weight, min_height, max_height, 
				 chest_cm, waist_cm, hip_cm, length_cm, created_at, updated_at 
				 FROM size_guides WHERE category = ?`
		args = append(args, category)
	}
	
	rows, err := h.DB.QueryContext(c.Request().Context(), query, args...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error al buscar las guías"})
	}
	defer rows.Close()

	guides := []model.SizeGuide{}

	for rows.Next() {
		var g model.SizeGuide
		
		err := rows.Scan(
			&g.ID, &g.Category, &g.FitType, &g.Size, &g.MinWeight, &g.MaxWeight, 
			&g.MinHeight, &g.MaxHeight, &g.ChestCm, &g.WaistCm, 
			&g.HipCm, &g.LengthCm, &g.CreatedAt, &g.UpdatedAt,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error leyendo datos"})
		}
		guides = append(guides, g)
	}

	return c.JSON(http.StatusOK, guides)
}

// Trae todas las reglas juntas (para la tabla del panel admin)
func (h *SizeGuideHandler) GetAllGuides(c echo.Context) error {
	
	query := `SELECT id, category, fit_type, size, min_weight, max_weight, min_height, max_height, 
			  chest_cm, waist_cm, hip_cm, length_cm, created_at, updated_at 
			  FROM size_guides ORDER BY category, fit_type, size`
	
	rows, err := h.DB.QueryContext(c.Request().Context(), query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error al buscar las guías"})
	}
	defer rows.Close()

	guides := []model.SizeGuide{}

	for rows.Next() {
		var g model.SizeGuide
		
		err := rows.Scan(
			&g.ID, &g.Category, &g.FitType, &g.Size, &g.MinWeight, &g.MaxWeight, 
			&g.MinHeight, &g.MaxHeight, &g.ChestCm, &g.WaistCm, 
			&g.HipCm, &g.LengthCm, &g.CreatedAt, &g.UpdatedAt,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error leyendo datos"})
		}
		guides = append(guides, g)
	}

	return c.JSON(http.StatusOK, guides)
}

// Elimina una regla por ID
func (h *SizeGuideHandler) DeleteSizeGuide(c echo.Context) error {
	id := c.Param("id")

	query := `DELETE FROM size_guides WHERE id = ?`
	_, err := h.DB.ExecContext(c.Request().Context(), query, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no se pudo borrar la guía"})
	}

	return c.NoContent(http.StatusNoContent)
}

// Función para actualizar una regla existente
func (h *SizeGuideHandler) UpdateSizeGuide(c echo.Context) error {
	id := c.Param("id")
	var guide model.SizeGuide

	if err := c.Bind(&guide); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "datos inválidos"})
	}

	query := `UPDATE size_guides SET 
			category = ?, fit_type = ?, size = ?, min_weight = ?, max_weight = ?, 
			min_height = ?, max_height = ?, chest_cm = ?, waist_cm = ?, 
			hip_cm = ?, length_cm = ? 
			WHERE id = ?`
	
	_, err := h.DB.ExecContext(c.Request().Context(), query, 
		guide.Category, guide.FitType, guide.Size, guide.MinWeight, guide.MaxWeight, 
		guide.MinHeight, guide.MaxHeight, guide.ChestCm, guide.WaistCm, 
		guide.HipCm, guide.LengthCm, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "no se pudo actualizar la guía"})
	}

	return c.JSON(http.StatusOK, guide)
}
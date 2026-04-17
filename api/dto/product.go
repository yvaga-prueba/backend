package dto

import "core/domain/model"
import "time"

type CreateProductRequest struct {
	BarCode     int64   `json:"bar_code" example:"7501234567890" validate:"required"`
	Title       string  `json:"title" example:"Remera Básica Negra" validate:"required"`
	Description string  `json:"description" example:"Remera de algodón 100% color negro, cuello redondo" validate:"required"`
	Stock       int64   `json:"stock" example:"50" validate:"required,min=0"`
	Size        string  `json:"size" example:"M" validate:"required,oneof=S M L XL XXL"` 
	Color       string  `json:"color"`
	Gender      string  `json:"gender"`
	FitType     string  `json:"fit_type"` 
	Category    string  `json:"category" example:"Remeras" validate:"required"`
	UnitPrice   float64 `json:"unit_price" example:"2500.00" validate:"required,min=0"`
}


func (r *CreateProductRequest) ToEntity() *model.Product {
	return &model.Product{
		BarCode:     r.BarCode,
		Title:       r.Title,
		Description: r.Description,
		Stock:       r.Stock,
		Size:        r.Size,
		Color:       r.Color,
		Gender:      r.Gender,
		FitType:     r.FitType, 
		Category:    r.Category,
		UnitPrice:   r.UnitPrice,
	}
}


type UpdateProductRequest struct {
	BarCode     *int64   `json:"bar_code,omitempty" example:"7501234567890"`
	Title       *string  `json:"title,omitempty" example:"Remera Básica Negra"`
	Description *string  `json:"description,omitempty" example:"Remera de algodón 100% color negro, cuello redondo"`
	Stock       *int64   `json:"stock,omitempty" example:"50"`
	Size        *string  `json:"size,omitempty" example:"M"` 
	Color       *string  `json:"color"`
	Gender      *string  `json:"gender"`
	FitType     *string  `json:"fit_type,omitempty"` 
	Category    *string  `json:"category,omitempty" example:"Remeras"`
	UnitPrice   *float64 `json:"unit_price,omitempty" example:"2500.00"`
}


func (r *UpdateProductRequest) ApplyToEntity(p *model.Product) {
	if r.BarCode != nil {
		p.BarCode = *r.BarCode
	}
	if r.Title != nil {
		p.Title = *r.Title
	}
	if r.Description != nil {
		p.Description = *r.Description
	}
	if r.Stock != nil {
		p.Stock = *r.Stock
	}
	if r.Size != nil {
		p.Size = *r.Size
	}
	if r.Color != nil {
		p.Color = *r.Color
	}
	if r.Gender != nil {
		p.Gender = *r.Gender
	}
	if r.FitType != nil {
		p.FitType = *r.FitType 
	}
	if r.Category != nil {
		p.Category = *r.Category
	}
	if r.UnitPrice != nil {
		p.UnitPrice = *r.UnitPrice
	}
}


type ProductResponse struct {
	ID          int64   `json:"id" example:"1"`
	BarCode     int64   `json:"bar_code" example:"7501234567890"`
	Title       string  `json:"title" example:"Remera Básica Negra"`
	Description string  `json:"description" example:"Remera de algodón 100% color negro, cuello redondo"`
	Stock       int64   `json:"stock" example:"50"`
	Size        string  `json:"size" example:"M"`
	Color       string  `json:"color" example:"Negro"`
	Gender      string  `json:"gender" example:"Unisex"`
	FitType     string  `json:"fit_type" example:"oversize"` 
	Category    string  `json:"category" example:"Remeras"`
	UnitPrice   float64 `json:"unit_price" example:"2500.00"`
	ImageURL    string  `json:"image_url,omitempty"` // URL de la imagen primaria (Google Drive)
	CreatedAt   time.Time `json:"created_at"`
}


func FromEntity(p model.Product) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		BarCode:     p.BarCode,
		Title:       p.Title,
		Description: p.Description,
		Stock:       p.Stock,
		Size:        p.Size,
		Color:       p.Color,
		Gender:      p.Gender,
		FitType:     p.FitType,
		Category:    p.Category,
		UnitPrice:   p.UnitPrice,
		CreatedAt:   p.CreatedAt,
	}
}


func FromEntityWithImage(p model.Product, imageURL string) ProductResponse {
	r := FromEntity(p)
	r.ImageURL = imageURL
	return r
}
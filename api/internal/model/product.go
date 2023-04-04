package model

import "time"

// ProductStatus represents the status of the product
type ProductStatus string

const (
	// ProductStatusActive means the product is active
	ProductStatusActive ProductStatus = "ACTIVE"
	// ProductStatusInactive means the product is inactive
	ProductStatusInactive ProductStatus = "INACTIVE"
	// ProductStatusDeleted means the product is deleted. This is for archival only
	ProductStatusDeleted ProductStatus = "DELETED"
)

// AllProductStatus returns list of all ProductStatus
var AllProductStatus = []ProductStatus{ProductStatusActive, ProductStatusInactive, ProductStatusDeleted}

// String converts to string value
func (p ProductStatus) String() string {
	return string(p)
}

// IsValid checks if product status is valid
func (p ProductStatus) IsValid() bool {
	return p == ProductStatusActive || p == ProductStatusInactive || p == ProductStatusDeleted
}

// Product represents the business model of product
type Product struct {
	ID          int64
	Price       int64
	ExternalID  string
	Description string
	Name        string
	Status      ProductStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

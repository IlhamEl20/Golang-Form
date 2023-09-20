package response

import "time"

type RoleResponseGet struct {
	ID        uint       `json:"id"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeletedBy string     `json:"deleted_by,omitempty"`
}

type RoleCreateRequest struct {
	Role string `json:"role" validate:"required"`
}

type RoleUpdateRequest struct {
	Role string `json:"role" validate:"required"`
}

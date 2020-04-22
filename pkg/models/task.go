package models

type Attribute string
type AttributeValue string

type Task struct {
	ID                    uint64
	CustomerTask          string     `json:"merchantTransaction"`
	PaymentStatus         TaskStatus `json:"state"`
	Category              TaskCategory
	SubCategory           TaskSubCategory
	SubCategoryAttributes map[Attribute]AttributeValue
	Groups                []TaskGroup
	Title                 string
	Description           string
	Cost
	Creator   Worker
	Payer     Worker
	Worker    Worker
	CreatedAt SSTime  `json:"date_created"`
	EndedAt   *SSTime `json:"date_end"`
	PayedAt   *SSTime `json:"date_paid"`
}

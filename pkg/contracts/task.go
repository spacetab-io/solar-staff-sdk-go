package contracts

import (
	"strings"

	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

type TaskResponse struct {
	ID                    uint64
	MerchantTransaction   string
	PaymentStatus         models.TaskStatus `json:"state"`
	Category              models.TaskCategory
	SubCategory           models.TaskSubCategory
	SubCategoryAttributes []string
	Groups                []models.TaskGroup
	Title                 string
	Description           string
	Cost
	Creator   models.Worker
	Payer     models.Worker
	Worker    models.Worker
	CreatedAt models.SSTime  `json:"date_created"`
	EndedAt   *models.SSTime `json:"date_end"`
	PayedAt   *models.SSTime `json:"date_paid"`
}

func (tr TaskResponse) ToModel() (*models.Task, error) {
	cost, err := tr.Cost.ToModel()
	if err != nil {
		return nil, err
	}
	sca := make(map[models.Attribute]models.AttributeValue, 0)
	for _, a := range tr.SubCategoryAttributes {
		aa := strings.Split(a, ":")
		sca[models.Attribute(aa[0])] = models.AttributeValue(aa[1])
	}

	return &models.Task{
		ID:                    tr.ID,
		CustomerTask:          tr.MerchantTransaction,
		PaymentStatus:         tr.PaymentStatus,
		Category:              tr.Category,
		SubCategory:           tr.SubCategory,
		SubCategoryAttributes: sca,
		Groups:                tr.Groups,
		Title:                 tr.Title,
		Description:           tr.Description,
		Cost:                  *cost,
		Creator:               tr.Creator,
		Payer:                 tr.Payer,
		Worker:                tr.Worker,
		CreatedAt:             tr.CreatedAt,
		EndedAt:               tr.EndedAt,
		PayedAt:               tr.PayedAt,
	}, nil
}

// Package orders Validations
package orders

import (
	"github.com/chaincode/ifinca/pkg/core/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Validate Validates the Order Structure
func (data Order) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.OrderNo, validation.Required.Error(utils.OrderNoRequired), validation.NotNil.Error(utils.OrderNoRequired)),
		validation.Field(&data.Quantity, validation.Required.Error(utils.QuantityRequired), validation.NotNil.Error(utils.QuantityRequired)),
		validation.Field(&data.Price, validation.Required.Error(utils.PriceRequired), validation.NotNil.Error(utils.PriceRequired)),
		validation.Field(&data.IfincaBonus, validation.Required.Error(utils.IfincaBonusRequired), validation.NotNil.Error(utils.IfincaBonusRequired)),
		validation.Field(&data.BaseUnit, validation.Required.Error(utils.BaseUnitRequired), validation.NotNil.Error(utils.BaseUnitRequired)),
		validation.Field(&data.PriceUnit, validation.Required.Error(utils.PriceUnitRequired), validation.NotNil.Error(utils.PriceUnitRequired)),
		validation.Field(&data.DeliveryDate, validation.Required.Error(utils.DeliveryDateRequired), validation.NotNil.Error(utils.DeliverDateRequired)),
		validation.Field(&data.ExporterDeliveryDate, validation.Required.Error(utils.ExporterDeliverDateRequired), validation.NotNil.Error(utils.ExporterDeliverDateRequired)),
		validation.Field(&data.ImporterDeliveryDate, validation.Required.Error(utils.ImporterDeliverDateRequired), validation.NotNil.Error(utils.ImporterDeliverDateRequired)),
		// validation.Field(&data.CafeDeliveryDate, validation.Required.Error(utils.CafeDeliverDateRequired), validation.NotNil.Error(utils.CafeDeliverDateRequired)),
		// validation.Field(&data.RoasterDeliveryDate, validation.Required.Error(utils.RoasterDeliverDateRequired), validation.NotNil.Error(utils.RoasterDeliverDateRequired)),
		validation.Field(&data.Importers, validation.Required.Error(utils.RoasterDeliverDateRequired), validation.NotNil.Error(utils.RoasterDeliverDateRequired)),
		// validation.Field(&data.Roasters, validation.Required.Error(utils.RoasterDeliverDateRequired), validation.NotNil.Error(utils.RoasterDeliverDateRequired)),
		// validation.Field(&data.CafeStores, validation.Required.Error(utils.RoasterDeliverDateRequired), validation.NotNil.Error(utils.RoasterDeliverDateRequired)),
		// validation.Field(&data.CupScore, validation.Required.Error(utils.CupScoreRequired), validation.NotNil.Error(utils.CupScoreRequired)),
		// validation.Field(&data.Country, validation.Required.Error(utils.CountryRequired), validation.NotNil.Error(utils.CountryRequired)),
		// validation.Field(&data.Process, validation.Required.Error(utils.ProcessRequired), validation.NotNil.Error(utils.ProcessRequired)),
		// validation.Field(&data.Region, validation.Required.Error(utils.RegionRequired), validation.NotNil.Error(utils.RegionRequired)),
		// validation.Field(&data.Variety, validation.Required.Error(utils.VerietyRequired), validation.NotNil.Error(utils.VerietyRequired)),
		// validation.Field(&data.Certificates, validation.Required.Error(utils.CertificatesRequired), validation.NotNil.Error(utils.CertificatesRequired)),
	)
}

// Validate Validates the Sub-Order Structure
func (data SubOrder) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.OrderNo, validation.Required.Error(utils.OrderNoRequired), validation.NotNil.Error(utils.OrderNoRequired)),
		validation.Field(&data.Supplier, validation.Required.Error(utils.SupplierRequired), validation.NotNil.Error(utils.SupplierRequired)),
		validation.Field(&data.Quantity, validation.Required.Error(utils.QuantityRequired), validation.NotNil.Error(utils.QuantityRequired)),
		// validation.Field(&data.ActionDate, validation.Required.Error(utils.ActionDateRequired), validation.NotNil.Error(utils.ActionDateRequired)),
	)
}

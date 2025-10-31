package validation

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Registrar validaciones personalizadas
	_ = validate.RegisterValidation("ethereum_address", validateEthereumAddress)
	_ = validate.RegisterValidation("ipfs_cid", validateIPFSCid)
}

// ValidateStruct valida una estructura
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return formatValidationErrors(validationErrors)
		}
		return err
	}
	return nil
}

// formatValidationErrors formatea los errores de validación
func formatValidationErrors(errors validator.ValidationErrors) error {
	var errMsg string
	for _, err := range errors {
		switch err.Tag() {
		case "required":
			errMsg += fmt.Sprintf("El campo '%s' es requerido. ", err.Field())
		case "oneof":
			errMsg += fmt.Sprintf("El campo '%s' debe ser uno de: %s. ", err.Field(), err.Param())
		case "ethereum_address":
			errMsg += fmt.Sprintf("El campo '%s' debe ser una dirección Ethereum válida. ", err.Field())
		case "ipfs_cid":
			errMsg += fmt.Sprintf("El campo '%s' debe ser un CID de IPFS válido. ", err.Field())
		default:
			errMsg += fmt.Sprintf("El campo '%s' falló la validación '%s'. ", err.Field(), err.Tag())
		}
	}
	return fmt.Errorf(errMsg)
}

// validateEthereumAddress valida una dirección Ethereum
func validateEthereumAddress(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	if address == "" {
		return true // Permitir vacío si no es required
	}

	// Validar formato 0x + 40 caracteres hexadecimales
	matched, _ := regexp.MatchString("^0x[0-9a-fA-F]{40}$", address)
	return matched
}

// validateIPFSCid valida un CID de IPFS
func validateIPFSCid(fl validator.FieldLevel) bool {
	cid := fl.Field().String()
	if cid == "" {
		return true // Permitir vacío si no es required
	}

	// CIDv0 comienza con Qm y tiene 46 caracteres
	// CIDv1 puede comenzar con diferentes caracteres
	matched, _ := regexp.MatchString("^(Qm[1-9A-HJ-NP-Za-km-z]{44}|b[A-Za-z2-7]{58,})$", cid)
	return matched
}

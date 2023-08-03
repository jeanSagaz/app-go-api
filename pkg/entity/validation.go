package entity

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

const tagCustom = "errormgs"

type DomainNotification struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func errorTagFunc[T interface{}](obj interface{}, snp string, fieldname, actualTag string) error {
	o := obj.(T)

	if !strings.Contains(snp, fieldname) {
		return nil
	}

	fieldArr := strings.Split(snp, ".")
	rsf := reflect.TypeOf(&o).Elem()

	for i := 1; i < len(fieldArr); i++ {
		field, found := rsf.FieldByName(fieldArr[i])
		if found {
			if fieldArr[i] == fieldname {
				customMessage := field.Tag.Get(tagCustom)
				if customMessage != "" {
					return fmt.Errorf("%s: %s (%s)", fieldname, customMessage, actualTag)
				}
				return nil
			} else {
				nestedFieldType := field.Type
				rsf = nestedFieldType
			}
		}
	}
	return nil
}

func errorCustomMessageFunc[T interface{}](obj interface{}, snp string) error {
	o := obj.(T)

	fieldArr := strings.Split(snp, ".")
	rsf := reflect.TypeOf(&o).Elem()

	for i := 1; i < len(fieldArr); i++ {
		field, found := rsf.FieldByName(fieldArr[i])
		if found {
			if fieldArr[i] != "" {
				customMessage := field.Tag.Get(tagCustom)
				if customMessage != "" {
					return fmt.Errorf("%s", customMessage)
				}
				return nil
			} else {
				nestedFieldType := field.Type
				rsf = nestedFieldType
			}
		}
	}
	return nil
}

func ValidateFunc[T interface{}](obj interface{}, validate *validator.Validate) (errs error) {
	o := obj.(T)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Validate:", r)
			errs = fmt.Errorf("can't validate %+v", r)
		}
	}()

	if err := validate.Struct(o); err != nil {
		errorValid := err.(validator.ValidationErrors)
		for _, e := range errorValid {
			// snp  X.Y.Z
			snp := e.StructNamespace()
			errmgs := errorTagFunc[T](obj, snp, e.Field(), e.ActualTag())
			if errmgs != nil {
				errs = errors.Join(errs, fmt.Errorf("%w", errmgs))
			} else {
				errs = errors.Join(errs, fmt.Errorf("%w", e))
			}
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}

func ValidateStruct[T interface{}](obj interface{}, validate *validator.Validate) (errs []DomainNotification) {
	o := obj.(T)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Validate:", r)
			errs = append(errs, DomainNotification{
				Key:   "",
				Value: fmt.Sprintf("can't validate %+v", r),
			})
		}
	}()

	if err := validate.Struct(o); err != nil {
		errorValid := err.(validator.ValidationErrors)
		for _, e := range errorValid {
			// snp  X.Y.Z
			snp := e.StructNamespace()
			errmgs := errorCustomMessageFunc[T](obj, snp)
			if errmgs != nil {
				errs = append(errs, DomainNotification{
					Key:   e.Field(),
					Value: errmgs.Error(),
				})
			} else {
				errs = append(errs, DomainNotification{
					Key:   e.Field(),
					Value: e.Namespace(),
				})
			}
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}

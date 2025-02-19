package enum

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"

	"github.com/xybor-x/enum/internal/core"
	"github.com/xybor-x/enum/internal/xreflect"
	"gopkg.in/yaml.v3"
)

var _ newableEnum = WrapFloatEnum[int](0)
var _ hookAfterEnum = WrapFloatEnum[int](0)

// WrapFloatEnum provides a set of built-in methods to simplify working with
// float64 enums.
type WrapFloatEnum[underlyingEnum any] float64

func (e WrapFloatEnum[underlyingEnum]) IsValid() bool {
	return IsValid(e)
}

func (e WrapFloatEnum[underlyingEnum]) MarshalJSON() ([]byte, error) {
	return MarshalJSON(e)
}

func (e *WrapFloatEnum[underlyingEnum]) UnmarshalJSON(data []byte) error {
	return UnmarshalJSON(data, e)
}

func (e WrapFloatEnum[underlyingEnum]) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return MarshalXML(encoder, start, e)
}

func (e *WrapFloatEnum[underlyingEnum]) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	return UnmarshalXML(decoder, start, e)
}

func (e WrapFloatEnum[underlyingEnum]) MarshalYAML() (any, error) {
	return MarshalYAML(e)
}

func (e *WrapFloatEnum[underlyingEnum]) UnmarshalYAML(node *yaml.Node) error {
	return UnmarshalYAML(node, e)
}

func (e WrapFloatEnum[underlyingEnum]) Value() (driver.Value, error) {
	return ValueSQL(e)
}

func (e *WrapFloatEnum[underlyingEnum]) Scan(a any) error {
	return ScanSQL(a, e)
}

// To returns the underlying representation of this enum.
func (e WrapFloatEnum[underlyingEnum]) To() underlyingEnum {
	return MustTo[underlyingEnum](e)
}

func (e WrapFloatEnum[underlyingEnum]) String() string {
	return ToString(e)
}

func (e WrapFloatEnum[underlyingEnum]) GoString() string {
	if !e.IsValid() {
		return fmt.Sprintf("%f", e)
	}

	return fmt.Sprintf("%f (%s)", e, e)
}

// WARNING: Only use this function if you fully understand its behavior.
// It might cause unexpected results if used improperly.
func (e WrapFloatEnum[underlyingEnum]) newEnum(repr []any) any {
	numeric := core.GetNumericRepresentation(repr)
	if numeric == nil {
		numeric = core.GetAvailableEnumValue[WrapFloatEnum[underlyingEnum]]()
	} else {
		repr = core.RemoveNumericRepresentation(repr)
	}

	return core.MapAny(xreflect.Convert[WrapFloatEnum[underlyingEnum]](numeric), repr)
}

// WARNING: Only use this function if you fully understand its behavior.
// It might cause unexpected results if used improperly.
func (e WrapFloatEnum[underlyingEnum]) hookAfter() {
	mustHaveUnderlyingRepr[underlyingEnum](e)
}

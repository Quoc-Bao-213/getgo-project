package generator

import (
	"github.com/Quoc-Bao-213/getgo-project/api/internal/pkg/snowflake"
)

var (
	ProductSNF snowflake.SnowflakeGenerator
)

func InitSnowflakeGenerators() error {
	ProductSNF = snowflake.New()

	return nil
}

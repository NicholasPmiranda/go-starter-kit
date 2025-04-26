package conversionTypes

import "github.com/jackc/pgx/v5/pgtype"

func ConvertPgInt8Slice(input []pgtype.Int8) []int64 {
	var result []int64
	for _, v := range input {
		if v.Valid {
			result = append(result, v.Int64)
		}
	}
	return result
}

package utils

import "fmt"

func LikeFormat(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}

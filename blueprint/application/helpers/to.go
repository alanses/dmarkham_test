package helpers

func ToString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

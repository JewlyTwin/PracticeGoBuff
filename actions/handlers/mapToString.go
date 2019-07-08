package handlers

func DynamicPostForm(x map[string][]string) map[string]interface{} {
	v := make(map[string]interface{})
	for key, value := range x {
		v[key] = value[0]
	}
	return v
}

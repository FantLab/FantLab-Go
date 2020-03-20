package reflectutils

import (
	"fantlab/base/assert"
	"testing"
)

func Test_SetStructValues(t *testing.T) {
	type x struct {
		Id    int    `json:"id"`
		Value uint   `json:"value"`
		Kek   bool   `json:"kek"`
		Name  string `json:"name"`
	}

	t.Run("positive", func(t *testing.T) {
		var output x
		SetStructValues(&output, "json", func(s string) string {
			switch s {
			case "id":
				return "-100"
			case "value":
				return "100"
			case "kek":
				return "true"
			case "name":
				return "user"
			}
			return ""
		})
		assert.True(t, output.Id == -100)
		assert.True(t, output.Name == "user")
		assert.True(t, output.Value == 100)
		assert.True(t, output.Kek == true)
	})
}

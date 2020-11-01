package TDD

import "testing"

func TestJson(t *testing.T) {
	myStructTestResponse := []struct {
		name string
		want []byte
	}{{name: "Test One", want: []byte(`{"name":"Lucas","code":"Success","body":"Hello"}`)}}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			got := returnJson()
			t.Helper()
			if string(got) != string(tt.want) {
				t.Errorf("body is wrong, got %q want %q\n", got, tt.want)
			}
		})
	}
}

func AssertResponsebody(t *testing.T, got, want string) {

}

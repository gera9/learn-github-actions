package pkg

import "testing"

func TestGetGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test 1 ok",
			want: "Hello, world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGreeting(); got != tt.want {
				t.Errorf("GetGreeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

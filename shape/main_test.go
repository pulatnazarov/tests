package shape

import "testing"

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		name string
		s    Shape
		want float64
	}{
		{
			name: "rectangle",
			s:    Rectangle{10.0, 10.0},
			want: 40.0,
		},
		{
			name: "circle",
			s:    Circle{3.0},
			want: 18.0,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.s, tt.want)
		})
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name string
		s    Shape
		want float64
	}{
		{
			name: "rectangle",
			s:    Rectangle{12.0, 6.0},
			want: 72.0,
		},
		{
			name: "circle",
			s:    Circle{10.0},
			want: 300.0,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.s, tt.want)
		})
	}
}

func checkPerimeter(t *testing.T, s Shape, want float64) {
	t.Helper()

	got := s.Perimeter()
	assertFloat(t, got, want)
}

func checkArea(t *testing.T, s Shape, want float64) {
	t.Helper()

	got := s.Area()
	assertFloat(t, got, want)
}

func assertFloat(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

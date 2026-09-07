package currency

import (
	"fmt"
	"testing"
)

type MockConvertor struct {
	lastAmount float64
	lastFrom   string
	lastTo     string
	calls      int
}

func (m *MockConvertor) Convert(amount float64, from, to string) float64 {
	m.lastAmount, m.lastFrom, m.lastTo = amount, from, to
	m.calls++
	return 42.0
}

func TestPriceIn(t *testing.T) {
	m := &MockConvertor{}
	tests := []struct {
		amount float64
		from   string
		to     string
		c      Converter
		want   float64
	}{
		{
			amount: 1,
			from:   "dollar",
			to:     "rub",
			c:      m,
			want:   42,
		},
		{
			amount: 100,
			from:   "dollar",
			to:     "rub",
			c:      m,
			want:   42,
		},
		{
			amount: 0,
			from:   "dollar",
			to:     "rub",
			c:      m,
			want:   42,
		},
		{
			amount: -5,
			from:   "dollar",
			to:     "rub",
			c:      m,
			want:   42,
		},
	}

	for i, tt := range tests {
		name := fmt.Sprintf("Testing PriceIn with %v", tt.amount)
		t.Run(name, func(t *testing.T) {
			got := PriceIn(tt.amount, tt.from, tt.to, tt.c)
			if got != tt.want {
				t.Errorf("PriceIn() = %v, want %v", got, tt.want)
			}
			if m.calls != i+1 {
				t.Error("Expected 1 call")
			}
			if m.lastAmount != tt.amount || m.lastFrom != tt.from || m.lastTo != tt.to {
				t.Error("Expected Convert with correct parameters")
			}
		})
	}
}

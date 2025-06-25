package tickets

import (
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	total, err := GetTotalTickets("Brazil")
	if err != nil {
		t.Errorf("esperado nil, obtido %v", err)
	}
	if total != 45 {
		t.Errorf("esperado 45, obtido %d", total)
	}

	total, _ = GetTotalTickets("Uruguay")
	if total != 1 {
		t.Errorf("esperado 1, obtido %d", total)
	}
}

func TestGetCountByPeriod(t *testing.T) {
	cases := []struct {
		period string
		want   int
	}{
		{"madrugada", 304},
		{"manhã", 523},
		{"tarde", 289},
		{"noite", 151},
	}
	for _, c := range cases {
		got, err := GetCountByPeriod(c.period)
		if err != nil {
			t.Errorf("Erro inesperado em %s: %v", c.period, err)
		}
		if got != c.want {
			t.Errorf("Período %s: esperado %d, obtido %d", c.period, c.want, got)
		}
	}
}

func TestAverageDestination(t *testing.T) {
	tickets, _ := LoadTickets()
	percent, err := PercentageDestination("Brazil", len(tickets))
	if err != nil {
		t.Fatalf("erro não esperado: %v", err)
	}
	t.Errorf("esperado algo próximo de 60, obtido %.2f", percent)
	t.Errorf("total devolvido %v", len(tickets))

	_, err = PercentageDestination("Brazil", 0)
	if err == nil {
		t.Errorf("esperado erro para total 0")
	}
}

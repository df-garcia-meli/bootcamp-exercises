package tickets_test

import (
	"desafio-go-bases/internal/tickets"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("Happy Path for Brazil", func(t *testing.T) {
		// Given
		sales := tickets.Sales{
			Tickets: []tickets.Ticket{
				tickets.NewTicket(1, "John Doe", "", "Brazil", "10:00", 100.0),
				tickets.NewTicket(2, "Jane Doe", "", "Brazil", "20:42", 200.0),
			},
		}

		var expectedResult int = 2
		var expectedError error = nil

		// When
		result, err := sales.GetTotalTickets("Brazil")

		// Then
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("Empty Tickets Slice", func(t *testing.T) {
		// Given
		sales := tickets.Sales{
			Tickets: []tickets.Ticket{},
		}

		var expectedResult int = 0

		// When
		result, err := sales.GetTotalTickets("Brazil")

		// Then
		require.NotNil(t, err)
		require.Equal(t, expectedResult, result)
	})
}

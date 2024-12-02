package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_exerc1(t *testing.T) {

	t.Run("Funcionario ganha menos de US$ 50.000", func(t *testing.T) {
		expected := 0.00
		result := exerc1(50000)
		require.Equal(t, expected, result)
	})

	t.Run("Funcionario ganha mais de US$ 50.000", func(t *testing.T) {
		expected := 55000 * 0.17
		result := exerc1(55000)
		require.Equal(t, expected, result)
	})

	t.Run("Funcionario ganha mais de US$ 150.000", func(t *testing.T) {
		expected := 155000 * 0.27
		result := exerc1(155000)
		require.Equal(t, expected, result)
	})
}

func Test_exerc2(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		expected := (10 + 8 + 5 + 6.5) / 4
		result := exerc2(10, 8, 5, 6.5)
		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := (10 + 10 + 6.5) / 3
		result := exerc2(10, 10, 6.5)
		require.Equal(t, expected, result)
	})

	t.Run("3", func(t *testing.T) {
		expected := (10 + 10 + 10 + 10.0 + 10) / 5
		result := exerc2(10, 10, 10, 10.0, 10)
		require.Equal(t, expected, result)
	})

	t.Run("4", func(t *testing.T) {
		expected := 0.0
		result := exerc2(-2.6, 4)
		require.Equal(t, expected, result)
	})

}

func Test_exerc3(t *testing.T) {

	t.Run("salário da categoria A", func(t *testing.T) {
		expected := 4500.00
		result := exerc3(60, "A")
		require.Equal(t, expected, result)
	})

	t.Run("salário da categoria B", func(t *testing.T) {
		expected := 1800.00
		result := exerc3(60, "B")
		require.Equal(t, expected, result)
	})

	t.Run("salário da categoria C", func(t *testing.T) {
		expected := 1000.00
		result := exerc3(60, "C")
		require.Equal(t, expected, result)
	})

}

// exercicio 4
func Test_minFunc(t *testing.T) {
	t.Run("o número mínimo de notas", func(t *testing.T) {
		expected := 3.0
		result := minFunc(3.0, 3.4, 5.0, 10.0)
		require.Equal(t, expected, result)
	})

	t.Run("o número mínimo de notas", func(t *testing.T) {
		expected := 2.0
		result := minFunc(2.0, 4.4, 6.0)
		require.Equal(t, expected, result)
	})
}

func Test_maxFunc(t *testing.T) {
	t.Run("o número máximo de notas", func(t *testing.T) {
		expected := 10.0
		result := maxFunc(3.0, 3.4, 5.0, 10.0)
		require.Equal(t, expected, result)
	})

	t.Run("o número máximo de notas", func(t *testing.T) {
		expected := 6.0
		result := maxFunc(4.0, 5.4, 6.0)
		require.Equal(t, expected, result)
	})
}

func Test_averageFunc(t *testing.T) {
	t.Run("o número médio de notas", func(t *testing.T) {
		expected := (10.0 + 10.0 + 10.0) / 3
		result := averageFunc(10.0, 10.0, 10.0)
		require.Equal(t, expected, result)
	})

	t.Run("o número médio de notas", func(t *testing.T) {
		expected := (5.0 + 8.8 + 9.0 + 6.7) / 4
		result := averageFunc(5.0, 8.8, 9.0, 6.7)
		require.Equal(t, expected, result)
	})
}

func Test_cachorro(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 3 * 10.0
		result := cachorro(3)
		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 5 * 10.0
		result := cachorro(5)
		require.Equal(t, expected, result)
	})
}

func Test_gato(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 2 * 5.0
		result := gato(2)
		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 8 * 5.0
		result := gato(8)
		require.Equal(t, expected, result)
	})
}

func Test_hamster(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 5 * 0.250
		result := hamster(5)
		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 1 * 0.250
		result := hamster(1)
		require.Equal(t, expected, result)
	})
}

func Test_tarantula(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 1.5
		result := tarantula(10)
		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 4 * 0.150
		result := tarantula(4)
		require.Equal(t, expected, result)
	})
}

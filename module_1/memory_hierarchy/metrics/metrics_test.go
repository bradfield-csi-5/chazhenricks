package metrics

import (
	"math"
	"testing"
)

func BenchmarkMetrics(b *testing.B) {
	users := LoadData()
  users2 := LoadData2()

	b.Run("Average age", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = AverageAge(users)
		}
		expected := 59.62
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected average age to be around %.2f, not %.3f", expected, actual)
		}
	})
	b.Run("Average age2", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = AverageAge2(users2)
		}
		expected := 59.62
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected average age to be around %.2f, not %.3f", expected, actual)
		}
	})

	b.Run("Average payment", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
	   	actual = AveragePaymentAmount(users)
		}

		expected := 499850.56
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.09 {
			b.Fatalf("Expected average payment amount to be around %.2f, not %.3f", expected, actual)
		}
	})
	b.Run("Average payment2", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = AveragePaymentAmount2(users2)
		}

		expected := 499850.064
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.09 {
			b.Fatalf("Expected average payment amount to be around %.2f, not %.3f", expected, actual)
		}
	})

	b.Run("Payment stddev", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = StdDevPaymentAmount(users)
		}
		expected := 288684.850
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected standard deviation to be around %.2f, not %.3f", expected, actual)
		}
	})
	b.Run("Payment stddev2", func(b *testing.B) {
		actual := 0.0
		for n := 0; n < b.N; n++ {
			actual = StdDevPaymentAmount2(users2)
		}
		expected := 288684.850
		if math.IsNaN(actual) || math.Abs(actual-expected) > 0.01 {
			b.Fatalf("Expected standard deviation to be around %.2f, not %.3f", expected, actual)
		}
	})

}

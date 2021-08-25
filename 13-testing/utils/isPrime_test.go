package utils

import "testing"

func TestIsPrime(t *testing.T) {
	//Arrange
	no := 11
	expectedResult := true

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {
		t.Fail()
	}
}

//sub tests
/* func TestUtils(t *testing.T) {
	t.Run("Number Tests", func(t *testing.T) {
		t.Run("Prime Test", func(t *testing.T) {
			t.Run("Testing Prime for 11", func(t *testing.T) {
				//Arrange
				no := 11
				expectedResult := false

				//Act
				actualResult := IsPrime(no)

				//Assert
				if actualResult != expectedResult {
					t.Fail()
				}
			})
			t.Run("Testing Prime for 13", func(t *testing.T) {
				//Arrange
				no := 13
				expectedResult := true

				//Act
				actualResult := IsPrime(no)

				//Assert
				if actualResult != expectedResult {
					t.Fail()
				}
			})
		})
	})
} */

type PrimeTestCase struct {
	name           string
	no             int
	expectedResult bool
}

func TestUtils(t *testing.T) {
	t.Run("Number Tests", func(t *testing.T) {
		t.Run("Prime Test", func(t *testing.T) {
			testCases := []PrimeTestCase{
				{name: "Testing Prime for 11", no: 11, expectedResult: true},
				{name: "Testing Prime for 13", no: 13, expectedResult: true},
				{name: "Testing Prime for 15", no: 15, expectedResult: false},
				{name: "Testing Prime for 17", no: 17, expectedResult: true},
			}
			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					//Arrange
					//Act
					actualResult := IsPrime(testCase.no)

					//Assert
					if actualResult != testCase.expectedResult {
						t.Fail()
					}
				})
			}
		})
	})
}

func BenchmarkIsPrime_177993(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(177993)
	}
}

func BenchmarkIsPrime_13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(13)
	}
}

func BenchmarkIsPrime_1779(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(1779)
	}
}

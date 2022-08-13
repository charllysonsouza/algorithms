package sort

func Quicksort(numbers []int) []int {

	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]

	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	menores, maiores := partition(n, pivot)
	return append(append(Quicksort(menores), pivot), Quicksort(maiores)...)
}

func partition(numbers []int, pivot int) (menores []int, maiores []int) {
	for _, value := range numbers {
		if value <= pivot {
			menores = append(menores, value)
		} else {
			maiores = append(maiores, value)
		}
	}
	return menores, maiores
}

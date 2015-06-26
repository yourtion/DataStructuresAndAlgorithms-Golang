// SelectionSort  选择排序

package selectionsort

func SelectionSort(values []int) {

	n := len(values)

	for i := 0; i < n; i++ {
		min := i
		//从第i+1个元素开始，找最小值
		for j := i + 1; j < n; j++ {
			if values[min] > values[j] {
				min = j
			}
		}
		//找到之后和第i个元素交换
		values[i], values[min] = values[min], values[i]
	}

}

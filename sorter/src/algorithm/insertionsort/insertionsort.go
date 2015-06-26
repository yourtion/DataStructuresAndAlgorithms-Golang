//InsertionSort 插入排序

package insertionsort

func InsertionSort(values []int) {

	n := len(values)

	//从第二个元素开始
	for i := 1; i < n; i++ {
		//从第i个元素开始，一次和前面已经排好序的i-1个元素比较，如果小于，则交换
		for j := i; j > 0; j-- {
			if values[j] < values[j-1] {
				values[j-1], values[j] = values[j], values[j-1]
			} else {
				//如果大于，则不用继续往前比较了，因为前面的元素已经排好序，比较大的大就是较大的了。
				break
			}
		}

	}

}

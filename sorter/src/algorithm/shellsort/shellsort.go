// ShellSort  希尔排序

package shellsort

func ShellSort(values []int) {

	n := len(values)
	h := 1
	//初始最大步长
	for {
		if h < n/3 {
			h = h*3 + 1
		} else {
			break
		}
	}

	for {
		if h >= 1 {
			//从第二个元素开始
			for i := 1; i < n; i++ {
				//从第i个元素开始，依次次和前面已经排好序的i-h个元素比较，如果小于，则交换
				for j := i; j >= h; j = j - h {
					if values[j] < values[j-h] {
						values[j-h], values[j] = values[j], values[j-h]
					} else {
						break
					}
				}
			}
			//步长除3递减
			h = h / 3
		} else {
			break
		}
	}

}

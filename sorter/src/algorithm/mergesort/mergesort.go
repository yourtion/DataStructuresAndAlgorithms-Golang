// MergeSort  合并排序

package mergesort

import "fmt"

var aux = make([]int, 10)

func merge(array []int, lo, mid, hi int) {
	i := lo
	j := mid + 1
	fmt.Print(lo, hi, array, "\n")
	//把元素拷贝到辅助数组中
	for k := lo; k <= hi; k++ {
		aux[k] = array[k]
	}
	//fmt.Print("Aux", aux)
	//然后按照规则将数据从辅助数组中拷贝回原始的array中
	for k := lo; k <= hi; k++ {
		//如果左边元素没了， 直接将右边的剩余元素都合并到到原数组中
		if i > mid {
			j++
			array[k] = aux[j]
		} else if j > hi {
			//如果右边元素没有了，直接将所有左边剩余元素都合并到原数组中
			i++
			array[k] = aux[i]
		} else if aux[i] < aux[j] {
			//如果左边右边小，则将左边的元素拷贝到原数组中
			i++
			array[k] = aux[i]
		} else {
			j++
			array[k] = aux[j]
		}
	}

}

func sort(array []int, lo, hi int) {
	if lo >= hi {
		return //如果下标大于上标，则返回
	}
	mid := lo + (hi-lo)/2     //平分数组
	sort(array, lo, mid)      //循环对左侧元素排序
	sort(array, mid+1, hi)    //循环对右侧元素排序
	merge(array, lo, mid, hi) //对左右排好的序列进行合并
}

func MergeSort(values []int) {
	//fmt.Print(len(values), "\n")
	sort(values, 0, len(values)-1)
}

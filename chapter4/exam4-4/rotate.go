/*
 * 编写一个rotate函数，通过一次循环完成旋转。
 * 分步:
 *	首先进行原数组的整体反转；
 *	将反转后的数组的当前k个元素进行反转；
 *	再将后n-k个元素进行反转.
 */

package rotate

// Time complexity: O(n)
func rotate(s *[]int, k int) {
	n := len(*s)
	k = k % n
	reverse(s, 0, n-1)
	reverse(s, 0, k-1)
	reverse(s, k, n-1)
}

func reverse(s *[]int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

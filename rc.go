func gcd(x, y int64) int64 { //辗转相除法
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	}
	return y
}
func lcm(x, y int64) int64 {
	return x / gcd(x, y) * y //计算lcm的小技巧，先除gcd，再乘另一个数，有效防止溢出
}
func main() {
	var (
		n    = int64(1000000000000000000)
		i    = uint(1)
		j    = uint(0)
		nums = []int{29516, 34882, 63315, 28983, 7176, 96533, 33422, 84834, 43803, 61310}
		len  = uint(len(nums))
	)
	sum := int64(0)
	for i = 0; i < (1 << len); i++ {
		count := 0
		curLcm := int64(1)
		ok := true
		for j = 0; j < len; j++ {
			if (i & (1 << j)) > 0 {
				count++
				tmp := curLcm
				curLcm = lcm(int64(nums[j]), curLcm)
				//这里判断乘法溢出！！被卡了好久
				if curLcm > n || curLcm < tmp || curLcm % tmp != 0 || curLcm % int64(nums[j]) != 0 {
					ok = false
					break
				}
			}
		}
		if !ok {
			continue
		}
		if count%2 == 1 {
			sum -= n / curLcm
		} else {
			sum += n / curLcm
		}
	}
	fmt.Println(sum)
}

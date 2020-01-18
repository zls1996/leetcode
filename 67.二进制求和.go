/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
    if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	if len(a) > len(b) {
		return addBinaryByStr(a, b)
	}

    return addBinaryByStr(b , a)

}

func addBinaryByStr(a, b string) string {
	temp := 0
	index := 1
    resultArray := make([] byte, len(a))
	for ; (index <= len(a) && index <= len(b)); index++ {
		byteA := a[len(a) - index]
		byteB := b[len(b) - index]
		result, result1 := plusBinary(byteA, byteB, temp)
        temp = result1
        resultArray[len(a) - index] = result
	}

	for ; index <= len(a); index++ {
		byteA := a[len(a) - index]
		intA := getValue(byteA)
		resultArray[len(a) - index] = getByte((intA + temp) % 2)
		temp = (temp + intA) / 2
	}

	if temp == 0 {
		return string(resultArray[:])
	}

	finalResult := make([]byte, len(resultArray)+1)
	for i := 0; i < len(resultArray); i++ {
		finalResult[i+1] = resultArray[i]
	}
	finalResult[0] = '1'
	return string(finalResult[:])
}

func plusBinary(byteA byte, byteB byte, temp int) (byte, int) {
	intA := getValue(byteA)
	intB := getValue(byteB)

	result1 := getByte((intA + intB + temp) % 2)
	result2 := (intA + intB + temp) / 2
	return result1, result2
}

func getValue(byteValue byte) int {
	if byteValue == '0' {
		return 0
	}
	return 1
}

func getByte(intValue int) byte {
	if intValue == 0 {
		return '0'
	}
	return '1'
}
// @lc code=end


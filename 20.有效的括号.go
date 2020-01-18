/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
type Stack struct {
	Values [] byte
}

func NewStack() *Stack {
	return &Stack{Values: make([] byte, 0)}
}

func (stack *Stack) Push(value byte) bool {
	stack.Values = append(stack.Values, value)
	return true
}

func (stack *Stack) Pop() (byte, error) {
    if (len(stack.Values) -1 ) < 0 {
        return ' ', errors.New("Index out of bound")
    }
	value := stack.Values[len(stack.Values) -1]
	stack.Values = stack.Values[:len(stack.Values) - 1]
	return value, nil
}

func (stack *Stack) Size() int {
    return len(stack.Values)
}


func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := NewStack()
	for index := 0; index < len(s); index++ {
		val := s[index] 
		if val == '(' {
			stack.Push(val)
			continue
		}

		if val == ')' {
			topValue, err := stack.Pop()
			if err != nil {
				return false
			}
			if topValue != '(' {
				return false
			}
			continue
		}

		if val == '{' {
			stack.Push(val)
			continue
		}

		if val == '}' {
			topValue, err := stack.Pop()
			if err != nil {
				return false
			}
			if topValue != '{' {
				return false
			}
			continue
		}

		if val == '[' {
			stack.Push(val)
			continue
		}

		if val == ']' {
			topValue, err := stack.Pop()
			if err != nil {
				return false
			}
			if topValue != '[' {
				return false
			}
			continue
		}
	}

	return stack.Size() == 0
    
}

// @lc code=end


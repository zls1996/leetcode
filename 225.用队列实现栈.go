/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	Values[] int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack {
		Values: make([]int, 0), 
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Values = append(this.Values, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	lastIndex := len(this.Values) - 1
	val := this.Values[lastIndex]
	this.Values = this.Values[:lastIndex]
	return val
}


/** Get the top element. */
func (this *MyStack) Top() int {
	lastIndex := len(this.Values) - 1
	return this.Values[lastIndex]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.Values) == 0 {
		return true
	}

	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end


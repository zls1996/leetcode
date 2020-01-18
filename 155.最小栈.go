/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	Elements []int
	Min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
		Elements: make([]int, 0),
		Min : 0,
	}
}


func (this *MinStack) Push(x int)  {
	if len(this.Elements) == 0 {
		this.Min = x
	}

	if x < this.Min {
		this.Min = x
	}

    this.Elements = append(this.Elements, x)
}


func (this *MinStack) Pop()  {
	this.Elements = this.Elements[0:len(this.Elements)-1]
	if len(this.Elements) == 0 {
		return
	}
	min := this.Elements[0]

	for index := 1; index < len(this.Elements); index++ {
		if min > this.Elements[index] {
			min = this.Elements[index]
		}
	}

	this.Min = min
}


func (this *MinStack) Top() int {
    return this.Elements[len(this.Elements) - 1]
}


func (this *MinStack) GetMin() int {
    return this.Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end


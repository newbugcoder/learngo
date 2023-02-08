package Weekday

// Weekday 自定义类型
type Weekday int

// 声明每个枚举项的索引值
const (
	Sunday    Weekday = iota + 1 // Index = 1
	Monday                       // Index = 2
	Tuesday                      // Index = 3
	Wednesday                    // Index = 4
	Thursday                     // Index = 5
	Friday                       // Index = 6
	Saturday                     // Index = 7
)

var weekdayStr = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

// ======== 枚举实现 ========

// String - 返回枚举项的索引值
func (w Weekday) String() string {
	return weekdayStr[w-1]
}

// Index - 返回枚举项的字符值
func (w Weekday) Index() int {
	return int(w)
}

// ======== 辅助函数 ========

// Values 返回枚举的所有值
func Values() []string {
	return weekdayStr
}

// ExistOf 判断某值是否存在枚举值中
func ExistOf(str string) bool {
	for _, v := range weekdayStr {
		if v == str {
			return true
		}
	}
	return false
}
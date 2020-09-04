## 使用场景
找到两个区间的覆盖关系/覆盖内容，并使之合并
这里有用的就是，培养出来对overlap的判断方式的快速使用

判断两个**未排序**区间是否overlap：
``` golang
        var lo, hi int
		if a[0] < b[0] {
			lo = b[0]
		} else {
			lo = a[0]
		}
		if a[1] > b[1] {
			hi = b[1]
		} else {
			hi = a[1]
		}
		if lo <= hi {
			result = append(result, []int{lo, hi})
		}
```

判断两个**已排序**区间是否overlap：
``` golang 
	if a[1] > b[0] {
		// overlap
	}
```
本工具用于简化使用LeetCode Editor插件时构造参数的操作

1. 安装Goland的LeetCode Editor插件


2. 打开插件的设置界面,Custom Template打勾


3. TempFilePath 选择到项目的路径  如：D:\Code\Project


4. Code FileName 输入

```
$!velocityTool.camelCaseName(${question.titleSlug})
```

5. Code Template 输入

```
package cn

import (
	"leetcodehelper/helper"
)

//${question.title}
func $!velocityTool.camelCaseName(${question.titleSlug})Execute() {
	helper.Code("输入", $!velocityTool.smallCamelCaseName(${question.titleSlug}))
}

${question.code}
```

6. 使用LeetCode Editor插件生成一道题目，以第一题两数之和为例
   
   将helper.Code("输入", twoSum)

   改为helper.Code("nums = [2,7,11,15], target = 9", twoSum)

   将LeetCode.go的cn.Execute()

   改为cn.TwoSumExecute()

   直接运行或者以Debug模式运行即可
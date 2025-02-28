# lessonlistchanger
 课表制作


**项目名称**

lessonlistchanger
 课表制作



### 简介

该项目是一个用 Go 语言编写的程序，用于解析 Excel 文件并生成文本输出，以此来实现课表格式的转换。它提供了一种简单的方法来处理 Excel 文件中的数据，并根据特定规则生成输出文本。

### 功能特性

- 从 Excel 文件中提取数据并进行处理
- 支持替换列表功能，根据指定替换规则修改数据
- 生成文本输出文件，将处理后的数据保存到指定文件中

### 如何使用

1. 下载源代码或克隆本仓库到本地。
2. 安装 Go 编程环境。
3. 在终端中进入项目目录。
4. 运行 `go build` 命令编译代码。
5. 运行生成的可执行文件。
6. 根据提示选择一个 Excel 文件。
7. 程序将处理数据并生成名为 `output.txt` 的文本输出文件。

### 使用示例

```bash
$ go build
$ ./lessonlistchanger
请选择一个xls文件：
input.xls
已将结果输出到output.txt
```

### 依赖项

- `github.com/extrame/xls`: 用于读取 Excel 文件的库

### 注意事项

- 程序目前仅支持处理特定格式的 Excel 文件，其他格式可能导致错误。详情参考input.xls
- 替换列表文件格式应为逗号分隔的键值对，如 `old_value,new_value`。

### 贡献者

- [e7g](https://github.com/e7g)

### 许可证

MIT License

### 联系方式

如有任何问题或建议，请联系 [e7g](https://github.com/e7g)。

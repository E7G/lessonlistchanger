package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/extrame/xls"
)

func main() {
    // 用文件选择器选择一个xls文件
    filename, err := selectFile()
    if err != nil {
        fmt.Println(err)
        return
    }

    // 打开xls文件
    file, err := xls.Open(filename, "utf-8")
    if err != nil {
        fmt.Println(err)
        return
    }

    // 获取第一个工作表
    sheet := file.GetSheet(0)
    if sheet == nil {
        fmt.Println("工作表未找到")
        return
    }

    // 创建文件输出
    f, err := os.Create("output.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    // 创建可替换列表
    replacements := map[string]string{}
    file2, err := os.Open("replacements.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file2.Close()
    scanner := bufio.NewScanner(file2)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ",")
        if len(parts) != 2 {
            fmt.Println("替换列表格式错误")
            return
        }
        replacements[parts[0]] = parts[1]
    }

    for _, t := range []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"} {
        var rowStrings []string
	
       rowStrings = append(rowStrings, t)
	
        // 单元格内容为"周一"
        if t == "周一" {
            rowStrings = append(rowStrings, "升旗")
        } else {
            rowStrings = append(rowStrings, "早读")
        }

        // 获取工作表中的所有行
        for rowNum := uint16(0); rowNum <= sheet.MaxRow; rowNum++ {
            // 获取行
            row := sheet.Row(int(rowNum))

            // 遍历行中的列
            for colNum := row.FirstCol(); colNum < row.LastCol(); colNum++ {
                // 获取单元格
                cell := row.Col(int(colNum))

                if cell == t {
                    // 向下获取每一行的单元格内容，直到下面两个单元格都为空
                    for nextRowNum := rowNum + 1; nextRowNum <= sheet.MaxRow; nextRowNum++ {
                        nextRow := sheet.Row(int(nextRowNum))
                        nextCell1 := nextRow.Col(int(colNum))
                        if nextCell1 == "" {
                            continue
                        }

                        // 查找替换列表中的替换内容
                        if replacement, ok := replacements[nextCell1]; ok {
                            nextCell1 = replacement
                        }
                        rowStrings = append(rowStrings, nextCell1)
                    }

                    break
                }
            }
        }

        // 补全剩余单元格内容
        for len(rowStrings) < 10 {
            rowStrings = append(rowStrings, "无")
        }

        rowStrings = append(rowStrings, "晚修")
        rowStrings = append(rowStrings, "晚修")

        // 将获取到的单元格内容用","连接成字符串
        resultString := strings.Join(rowStrings, ",")

        // 输出到文件
        f.WriteString(resultString + ",\r\n")
    }

    fmt.Println("已将结果输出到output.txt")
}

// selectFile 使用文件选择器选择一个文件，并返回文件路径
func selectFile() (string, error) {
    fmt.Println("请选择一个xls文件：")

    var filename string
    fmt.Scanln(&filename)

    return filename, nil
}
/*
	转化C端提供的业绩数据为CRM需要的数据格式
	导入文件是 01.xlsx
	01.xlsx  字段分别为：  C端订单号  订单日期  订单金额  用户ID  工号   类型
	类型  1  2  代表回款   3  4 5  代表 退款
    生成文件   日期_标准化业绩格式.xlsx
*/
package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"
)

func main() {
	c := createData{}
	c.filePath = "./01.xlsx"
	filename := time.Now().Format("20060102150405")
	c.savePath = "./" + filename + "_标准化业绩格式.xlsx"

	rows := c.getExcelData()
	orderAmountList := c.formatExcelData(rows)
	c.createExcel(orderAmountList)
	defer func() {
		fmt.Println("程序执行完成")
	}()

}

type createData struct {
	filePath string // 文件名称
	savePath string
}

/**
*获取EXCEL数据
 */
func (c *createData) getExcelData() (rows [][]string) {
	f, err := excelize.OpenFile(c.filePath)
	c.checkError(err)
	rows, err = f.GetRows(f.GetSheetName(0))
	c.checkError(err)
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return
}

func (c *createData) checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
*格式化数据
 */
func (c *createData) formatExcelData(rows [][]string) map[string]float64 {
	orderAmountList := make(map[string]float64)
	for index, row := range rows {
		if index == 0 {
			continue
		}

		userOrderNum := row[4] + "_" + row[0]
		amount, _ := strconv.ParseFloat(row[2], 64)
		orderType, _ := strconv.Atoi(row[5])

		if orderType >= 3 {
			amount = 0 - amount

		}
		if _, ok := orderAmountList[userOrderNum]; ok {
			orderAmountList[userOrderNum] += amount
		} else {
			orderAmountList[userOrderNum] = amount
		}
		// 数据格式化 还未找到更好的方式 保留2位小数
		t,_ :=strconv.ParseFloat(fmt.Sprintf("%.2f", orderAmountList[userOrderNum]), 64)
		orderAmountList[userOrderNum] = t

	}
	return orderAmountList
}

/**
*创建新的excel
 */
func (c *createData) createExcel(rows map[string]float64) {

	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.SetCellValue(sheetName, "A1", "B端员工号")
	f.SetCellValue(sheetName, "B1", "DD订单号（必填）")
	f.SetCellValue(sheetName, "C1", "C端业绩金额（必填）")
	f.SetCellValue(sheetName, "D1", "是否完成对比（必填）")
	f.SetCellValue(sheetName, "E1", "负责人（必填）")
	f.SetCellValue(sheetName, "F1", "业务类型（必填）")

	//rowNum := len(rows)
	i := 1
	for userOrderStr, amount := range rows {
		i++
		num := strconv.Itoa(i)
		strIndex := strings.Index(userOrderStr, "_")
		userNumStr := userOrderStr[0:strIndex]
		orderNumStr := userOrderStr[(strIndex + 1):]

		f.SetCellValue(sheetName, "A"+num, userNumStr)
		f.SetCellValue(sheetName, "B"+num, orderNumStr)
		f.SetCellValue(sheetName, "C"+num, amount)
		f.SetCellValue(sheetName, "D"+num, "否")
		f.SetCellValue(sheetName, "E"+num, "张永胜")
		f.SetCellValue(sheetName, "F"+num, "预设业务类型")

	}
	if err := f.SaveAs(c.savePath); err != nil {
		fmt.Println(err)
	}

}

package main

import (
	"github.com/andlabs/ui"
	"math"
	"strconv"
)

func main() {
	err := ui.Main(func() {
		qishu := ui.NewEntry()
		zonge := ui.NewEntry()
		meiqijine := ui.NewEntry()
		button := ui.NewButton("计算")
		greeting := ui.NewLabel("计算结果: \n")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("输入期数:"), false)
		box.Append(qishu, false)
		box.Append(ui.NewLabel("输入总额:"), false)
		box.Append(zonge, false)
		box.Append(ui.NewLabel("输入每期应还金额:"), false)
		box.Append(meiqijine, false)
		box.Append(button, false)
		box.Append(greeting, false)

		//创建window窗口。并设置长宽。
		window := ui.NewWindow("等额本息计算器。", 600, 500, false)
		//mac不支持居中。
		//https://github.com/andlabs/ui/issues/162
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			qishuNum, err := strconv.ParseFloat(qishu.Text(), 64)
			zongeNum, err := strconv.ParseFloat(zonge.Text(), 64)
			meiqijineNum, err := strconv.ParseFloat(meiqijine.Text(), 64)
			if err != nil {
				greeting.SetText("出错了！\n " + err.Error())
				return
			}
			yuelilv, nianlilv := compute(qishuNum, zongeNum, meiqijineNum)
			greeting.SetText("月利息: " + strconv.FormatFloat(float64(yuelilv), 'f', 6, 64) + "\n" + "年利息: " + strconv.FormatFloat(float64(nianlilv), 'f', 6, 64) + "\n")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func compute(qishu, zonge, meiqi float64) (yuelilv, nianlilv float64) {
	var cifang = qishu
	var result = meiqi
	var benjin = zonge

	var x = 0.005
	var tag = 0.001
	var x1 = x + tag

	y := math.Pow(x+1, cifang)
	y1 := math.Pow(x1+1, cifang)

	for true {

		if (result-(benjin*x*y)/(y-1)) > 0 && (result-(benjin*x1*y1)/(y1-1)) > 0 {
			x = x + tag
			y = math.Pow(x+1, cifang)
			x1 = x1 + tag
			y1 = math.Pow(x1+1, cifang)
		} else if (result-(benjin*x*y)/(y-1)) < 0 && (result-(benjin*x1*y1)/(y1-1)) < 0 {
			if x-tag == 0 {
				tag = tag * 0.1
				x = x * 0.5
				x1 = x + tag
				y = math.Pow(x+1, cifang)
				y1 = math.Pow(x1+1, cifang)
				continue
			}
			x = x - tag
			y = math.Pow(x+1, cifang)
			x1 = x1 - tag
			y1 = math.Pow(x1+1, cifang)
		} else if (result-(benjin*x*y)/(y-1)) > 0 && (result-(benjin*x1*y1)/(y1-1)) < 0 {
			tag = tag * 0.1
			x = x + tag
			y = math.Pow(x+1, cifang)
			x1 = x + tag
			y1 = math.Pow(x1+1, cifang)
		} else if (result - (benjin*x*y)/(y-1)) == 0 {
			return x, x * 12
		} else if (result - (benjin*x1*y1)/(y1-1)) == 0 {
			return x1, x1 * 12
		} else {
			//错误数据
			return
		}

		if tag < 0.0000000000001 {
			return x, x * 12
		}
	}

	return

}

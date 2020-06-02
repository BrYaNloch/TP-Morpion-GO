package main

import "github.com/tadvi/winc"

func main() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)
	mainWindow.SetText("Hello World Demo")

	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)

	resetButton := winc.NewPushButton(mainWindow)
	resetButton.SetText("Reset")
	resetButton.SetPos(50, 200)
	resetButton.SetSize(100, 40)

	resetButton.OnClick().Bind(func(e *winc.Event) {
		mainWindow.Close()
		main()
	})

	var XouO = true //X c'est 1, O c'est 0

	grid00 := winc.NewPushButton(mainWindow)
	var grid00Filled = false
	grid00.SetText("-")
	grid00.SetPos(50, 50)
	grid00.SetSize(50, 50)
	grid00.OnClick().Bind(func(e *winc.Event) {
		if grid00Filled == false && XouO == true {
			grid00.SetText("X")
			XouO = false
			grid00Filled = true
		}
		if grid00Filled == false && XouO == false {
			grid00.SetText("O")
			XouO = true
			grid00Filled = true
		}
	})
	grid01 := winc.NewPushButton(mainWindow)
	var grid01Filled = false
	grid01.SetText("-")
	grid01.SetPos(100, 50)
	grid01.SetSize(50, 50)
	grid01.OnClick().Bind(func(e *winc.Event) {
		if grid01Filled == false && XouO == true {
			grid01.SetText("X")
			XouO = false
			grid00Filled = true
		}
		if grid01Filled == false && XouO == false {
			grid01.SetText("O")
			XouO = true
			grid01Filled = true
		}
	})
	grid02 := winc.NewPushButton(mainWindow)
	var grid02Filled = false
	grid02.SetText("-")
	grid02.SetPos(150, 50)
	grid02.SetSize(50, 50)
	grid02.OnClick().Bind(func(e *winc.Event) {
		if grid02Filled == false && XouO == true {
			grid02.SetText("X")
			XouO = false
			grid02Filled = true
		}
		if grid02Filled == false && XouO == false {
			grid02.SetText("O")
			XouO = true
			grid02Filled = true
		}
	})
	grid10 := winc.NewPushButton(mainWindow)
	var grid10Filled = false
	grid10.SetText("-")
	grid10.SetPos(50, 100)
	grid10.SetSize(50, 50)
	grid10.OnClick().Bind(func(e *winc.Event) {
		if grid10Filled == false && XouO == true {
			grid10.SetText("X")
			XouO = false
			grid10Filled = true
		}
		if grid10Filled == false && XouO == false {
			grid10.SetText("O")
			XouO = true
			grid10Filled = true
		}
	})
	grid11 := winc.NewPushButton(mainWindow)
	var grid11Filled = false
	grid11.SetText("-")
	grid11.SetPos(100, 100)
	grid11.SetSize(50, 50)
	grid11.OnClick().Bind(func(e *winc.Event) {
		if grid11Filled == false && XouO == true {
			grid11.SetText("X")
			XouO = false
			grid11Filled = true
		}
		if grid11Filled == false && XouO == false {
			grid11.SetText("O")
			XouO = true
			grid11Filled = true
		}
	})
	grid12 := winc.NewPushButton(mainWindow)
	var grid12Filled = false
	grid12.SetText("-")
	grid12.SetPos(150, 100)
	grid12.SetSize(50, 50)
	grid12.OnClick().Bind(func(e *winc.Event) {
		if grid12Filled == false && XouO == true {
			grid12.SetText("X")
			XouO = false
			grid12Filled = true
		}
		if grid12Filled == false && XouO == false {
			grid12.SetText("O")
			XouO = true
			grid12Filled = true
		}
	})
	grid20 := winc.NewPushButton(mainWindow)
	var grid20Filled = false
	grid20.SetText("-")
	grid20.SetPos(50, 150)
	grid20.SetSize(50, 50)
	grid20.OnClick().Bind(func(e *winc.Event) {
		if grid20Filled == false && XouO == true {
			grid20.SetText("X")
			XouO = false
			grid20Filled = true
		}
		if grid20Filled == false && XouO == false {
			grid20.SetText("O")
			XouO = true
			grid20Filled = true
		}
	})
	grid21 := winc.NewPushButton(mainWindow)
	var grid21Filled = false
	grid21.SetText("-")
	grid21.SetPos(100, 150)
	grid21.SetSize(50, 50)
	grid21.OnClick().Bind(func(e *winc.Event) {
		if grid21Filled == false && XouO == true {
			grid21.SetText("X")
			XouO = false
			grid21Filled = true
		}
		if grid21Filled == false && XouO == false {
			grid21.SetText("O")
			XouO = true
			grid21Filled = true
		}
	})
	grid22 := winc.NewPushButton(mainWindow)
	var grid22Filled = false
	grid22.SetText("-")
	grid22.SetPos(150, 150)
	grid22.SetSize(50, 50)
	grid22.OnClick().Bind(func(e *winc.Event) {
		if grid22Filled == false && XouO == true {
			grid22.SetText("X")
			XouO = false
			grid22Filled = true
		}
		if grid22Filled == false && XouO == false {
			grid22.SetText("O")
			XouO = true
			grid22Filled = true
		}
	})

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()
}
func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

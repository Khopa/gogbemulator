package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"khopa.github.io/gogbemulator/emulator"
)

// formatMemory Utility to print a memory section
func formatMemory(mem []uint8, sp uint16) string {
	var b strings.Builder
	for i := 0; i < len(mem); i += 16 {
		_, err := fmt.Fprintf(&b, "%04X: ", uint16(i)+sp)
		if err != nil {
			return ""
		}
		for j := 0; j < 16 && i+j < len(mem); j++ {
			_, err := fmt.Fprintf(&b, "%02X ", mem[i+j])
			if err != nil {
				return ""
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func main() {

	// Create emulator and load initial rom
	dmg := emulator.MakeDMG()
	dmg.Print()
	err := dmg.LoadROM("testrom.gb")
	if err != nil {
		fmt.Printf("Error loading ROM: %v\n, please add your rom with name 'testrom.gb'"+
			"in the working dir, it is not included by default in repo.", err)
		return
	}
	dmg.Gbz80.Pc = 0x150

	dissasembly := emulator.Disassembly("testrom.gb", dmg.Gbz80.Pc, 20)

	// Create Fyne APP
	a := app.New()
	w := a.NewWindow("Go GB Emulator")

	// --- Registers Left Panel ---
	var updateRegisters func()
	var updateMemory func()

	screenImage := canvas.NewImageFromImage(dmg.Snapshot())
	screenImage.ScaleMode = canvas.ImageScalePixels
	screenImage.FillMode = canvas.ImageFillContain
	screenContainer := container.NewVBox(
		widget.NewLabel("LCD"),
		screenImage,
	)
	screenImage.SetMinSize(fyne.NewSize(
		160*3,
		144*3,
	))

	stepButton := widget.NewButton("Step", func() {
		dmg.Step()
		updateRegisters()
		updateMemory()
		screenImage.Image = dmg.Snapshot()
		screenImage.Refresh()
	})

	regLabel := widget.NewLabel("")
	regPanel := container.NewVBox(
		widget.NewLabel("Registers"),
		regLabel,
		widget.NewSeparator(),
		stepButton,
	)

	updateRegisters = func() {
		regLabel.SetText(fmt.Sprintf(
			"AF: 0x%04X\nBC: 0x%04X\nDE: 0x%04X\nHL: 0x%04X\nSP: 0x%04X\nPC: 0x%04X",
			dmg.Gbz80.Af,
			dmg.Gbz80.Bc,
			dmg.Gbz80.De,
			dmg.Gbz80.Hl,
			dmg.Gbz80.Sp,
			dmg.Gbz80.Pc,
		))
	}

	// --- Memory Viewer ---

	memEntry := widget.NewMultiLineEntry()
	memEntry.Wrapping = fyne.TextWrapOff
	memEntry.SetText(dissasembly)

	memPanel := container.NewBorder(
		widget.NewLabel("Memory"),
		nil, nil, nil,
		container.NewScroll(memEntry),
	)

	updateMemory = func() {
		dissasembly = emulator.Disassembly("testrom.gb", dmg.Gbz80.Pc, 15)
		memEntry.SetText(dissasembly)
	}

	// --- Layout Split ---
	topSplit := container.NewHSplit(
		regPanel,
		screenContainer,
	)
	topSplit.Offset = 0.25

	mainLayout := container.NewVSplit(
		topSplit,
		memPanel,
	)
	mainLayout.Offset = 0.45

	w.SetContent(mainLayout)

	// simulate emulator updates
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			// run
			// dmg.Step()
			fyne.Do(updateRegisters)
			//fyne.Do(updateMemory)
		}
	}()

	w.ShowAndRun()
}

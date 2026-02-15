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

	emulator.Dissasembly("testrom.gb")

	// Create Fyne APP
	a := app.New()
	w := a.NewWindow("Go GB Emulator")

	// --- Registers Left Panel ---
	var updateRegisters func()
	var updateMemory func()

	stepButton := widget.NewButton("Step", func() {
		dmg.Step()
		updateRegisters()
		updateMemory()
	})

	image := canvas.NewImageFromFile("./resources/gopher.png")
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(356, 356))

	regLabel := widget.NewLabel("")
	regPanel := container.NewVBox(
		image,
		widget.NewSeparator(),
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
	memEntry.SetText(formatMemory(dmg.Memory[0:512], 0))

	memPanel := container.NewBorder(
		widget.NewLabel("Memory"),
		nil, nil, nil,
		container.NewScroll(memEntry),
	)

	updateMemory = func() {
		memEntry.SetText(formatMemory(dmg.Memory[dmg.Gbz80.PC():dmg.Gbz80.PC()+512], dmg.Gbz80.PC()))
	}

	// --- Layout Split ---
	split := container.NewHSplit(
		regPanel,
		memPanel,
	)
	split.Offset = 0.25

	w.SetContent(split)
	w.Resize(fyne.NewSize(900, 600))

	// simulate emulator updates
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			// run
			// dmg.Step()
			fyne.Do(updateRegisters)
			fyne.Do(updateMemory)
		}
	}()

	w.ShowAndRun()
}

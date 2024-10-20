package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const MARGIN = 10

var (
	Red         = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	Green       = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	Blue        = color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	Yellow      = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	Cyan        = color.NRGBA{R: 0, G: 255, B: 255, A: 255}
	Magenta     = color.NRGBA{R: 255, G: 0, B: 255, A: 255}
	Orange      = color.NRGBA{R: 255, G: 165, B: 0, A: 255}
	Purple      = color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	Brown       = color.NRGBA{R: 165, G: 42, B: 42, A: 255}
	Black       = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	White       = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	Gray        = color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	Silver      = color.NRGBA{R: 192, G: 192, B: 192, A: 255}
	Pink        = color.NRGBA{R: 255, G: 192, B: 203, A: 255}
	LightBlue   = color.NRGBA{R: 173, G: 216, B: 230, A: 255}
	ForestGreen = color.NRGBA{R: 34, G: 139, B: 34, A: 255}
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func drawBox(gtx layout.Context, sizeX int, sizeY int, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(sizeX, sizeY)}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: image.Pt(sizeX, sizeY)}
}

func renderAddMenuItemButton(gtx layout.Context, theme *material.Theme, addMenuItemButton *widget.Clickable, displayMenu *bool) layout.Dimensions {
	buttonWidth := 100
	buttonLength := 50
	defer op.Offset(image.Pt(gtx.Constraints.Min.X-(buttonWidth+(MARGIN*2)), 0)).Push(gtx.Ops).Pop()

	for addMenuItemButton.Clicked(gtx) {
		*displayMenu = true
	}

	return layout.UniformInset(MARGIN).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = buttonWidth
		gtx.Constraints.Min.Y = buttonLength
		gtx.Constraints.Max = gtx.Constraints.Min

		btn := material.Button(theme, addMenuItemButton, "Add menu item")
		btn.Color = Black
		btn.Background = ForestGreen
		return btn.Layout(gtx)
	})
}

func renderConfirmMenuItemButton(gtx layout.Context, theme *material.Theme, addMenuItemButton *widget.Clickable, menuItemButtons *[]*widget.Clickable, displayMenu *bool, menuItemInput *widget.Editor, menuItemInputs *[]string) layout.Dimensions {
	buttonWidth := 100
	buttonLength := 50

	defer op.Offset(image.Pt(gtx.Constraints.Min.X-(buttonWidth+(MARGIN*2)), gtx.Constraints.Min.Y-(buttonLength+(MARGIN*2)))).Push(gtx.Ops).Pop()

	for addMenuItemButton.Clicked(gtx) {
		fmt.Println("Add Menu item menu!")
		addMenuItems(menuItemButtons, menuItemInput, menuItemInputs)
		*displayMenu = false
	}

	return layout.UniformInset(MARGIN).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = buttonWidth
		gtx.Constraints.Min.Y = buttonLength
		gtx.Constraints.Max = gtx.Constraints.Min

		btn := material.Button(theme, addMenuItemButton, "Add menu item")
		btn.Color = Gray
		btn.Background = Black
		return btn.Layout(gtx)
	})
}

func renderTitle(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	title := material.H6(theme, "simplePos")
	title.Color = ForestGreen
	return title.Layout(gtx)
}

func renderMenuItemButtons(theme *material.Theme, menuItemButtons *[]*widget.Clickable, menuItemInputs *[]string, gtx layout.Context) layout.StackChild {
	buttonWidth := 100
	buttonLength := 50
	menuItems := *menuItemButtons // need to dereference first
	menuItemLayouts := []layout.FlexChild{}
	lastSlice := 0

	type Test struct {
		MIL []layout.FlexChild
	}

	menuItemFlexes := []Test{}

	if len(menuItems) > 0 {
		for index, menuItem := range menuItems {
			if index != 0 && ((index+1)%5) == 0 {
				menuItemFlexes = append(menuItemFlexes, Test{MIL: menuItemLayouts[lastSlice:]})
				lastSlice = index + 1
				fmt.Println("we at index", index, lastSlice)
			}
			menuItemLayouts = append(menuItemLayouts,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(MARGIN).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						inputs := *menuItemInputs

						for menuItem.Clicked(gtx) {
							fmt.Println(inputs[index])
						}

						gtx.Constraints.Min.X = buttonWidth
						gtx.Constraints.Min.Y = buttonLength
						gtx.Constraints.Max = gtx.Constraints.Min

						btn := material.Button(theme, menuItem, inputs[index])
						btn.Color = Black
						btn.Background = Blue
						return btn.Layout(gtx)
					})
				}))
		}
	}

	return layout.Stacked(func(gtx layout.Context) layout.Dimensions {
		for index, flex := range menuItemFlexes {
			if index != 0 {
				defer op.Offset(image.Pt((buttonWidth+MARGIN)*index, 0)).Push(gtx.Ops).Pop()
			}
			layout.Flex{Axis: layout.Vertical}.Layout(gtx, flex.MIL...)
		}
		return layout.Dimensions{Size: gtx.Constraints.Max}
	})
}

func addMenuItems(menuItemButtons *[]*widget.Clickable, menuItemInput *widget.Editor, menuItemInputs *[]string) {
	menuItemButton := new(widget.Clickable)
	*menuItemButtons = append(*menuItemButtons, menuItemButton)
	menuItemText := menuItemInput.Text()
	*menuItemInputs = append(*menuItemInputs, menuItemText)
	menuItemTextReset := *menuItemInput
	menuItemTextReset.SetText("")
}

func renderLayout(gtx layout.Context, theme *material.Theme, addMenuItemButton *widget.Clickable, menuItems *[]*widget.Clickable, displayMenu *bool, confirmMenuItemButton *widget.Clickable, menuItemInput *widget.Editor, menuItemInputs *[]string) {
	layout.UniformInset(MARGIN).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			// Main layout backround
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					// Check Viewer
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return renderTitle(gtx, theme)
							}),
							layout.Rigid(layout.Spacer{Height: MARGIN}.Layout),
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return drawBox(gtx, gtx.Constraints.Max.X, gtx.Constraints.Max.Y, Silver)
							}),
						)
					}),
					layout.Rigid(layout.Spacer{Width: MARGIN}.Layout),
					// Menu Viewer
					layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Top: 34}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Stack{}.Layout(gtx,
								layout.Stacked(func(gtx layout.Context) layout.Dimensions {
									return drawBox(gtx, gtx.Constraints.Max.X, gtx.Constraints.Max.Y, Black)
								}),
								layout.Expanded(func(gtx layout.Context) layout.Dimensions {
									return layout.Stack{}.Layout(gtx,
										renderMenuItemButtons(theme, menuItems, menuItemInputs, gtx),
										layout.Expanded(func(gtx layout.Context) layout.Dimensions {
											return renderAddMenuItemButton(gtx, theme, addMenuItemButton, displayMenu)
										}),
									)
								}),
							)
						})
					}),
				)
			}),
			// Menu Editor
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(200).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					if *displayMenu {
						return layout.Stack{}.Layout(gtx,
							layout.Stacked(func(gtx layout.Context) layout.Dimensions {
								return drawBox(gtx, gtx.Constraints.Max.X, gtx.Constraints.Max.Y, Gray)
							}),
							layout.Expanded(func(gtx layout.Context) layout.Dimensions {
								layout.Inset{Top: MARGIN}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										drawBox(gtx, 255, 30, Black)
										menuItemEditorTheme := material.NewTheme()
										menuItemEditorTheme.Palette.Fg = White
										menuItemEditorTheme.TextSize = 30
										menuItemInput.SingleLine = true
										menuItemInput.MaxLen = 18
										menuItemEditor := material.Editor(menuItemEditorTheme, menuItemInput, "Menu item title...")
										return menuItemEditor.Layout(gtx)
									})
								})

								renderConfirmMenuItemButton(gtx, theme, confirmMenuItemButton, menuItems, displayMenu, menuItemInput, menuItemInputs)
								return layout.Dimensions{Size: gtx.Constraints.Max}
							}))
					}
					return layout.Dimensions{Size: gtx.Constraints.Max}
				})
			}),
		)
	})
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	addMenuItemButton := new(widget.Clickable)
	confirmMenuItemButton := new(widget.Clickable)
	menuItems := []*widget.Clickable{}
	displayMenu := false
	menuItemInput := new(widget.Editor)
	var menuItemInputs []string

	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			renderLayout(gtx, theme, addMenuItemButton, &menuItems, &displayMenu, confirmMenuItemButton, menuItemInput, &menuItemInputs)
			e.Frame(gtx.Ops)
		}
	}
}

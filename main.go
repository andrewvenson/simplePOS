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

func renderAddMenuItemButton(gtx layout.Context, theme *material.Theme, addMenuItemButton *widget.Clickable, menuItemButtons *[]*widget.Clickable) layout.Dimensions {
	for addMenuItemButton.Clicked(gtx) {
		fmt.Println("Add Menu item menu!")
		addMenuItems(menuItemButtons)
	}

	buttonWidth := 100
	buttonLength := 50

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

func renderTitle(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	title := material.H6(theme, "simplePos")
	title.Color = ForestGreen
	return title.Layout(gtx)
}

func renderMenuItemButtons(theme *material.Theme, menuItemButtons *[]*widget.Clickable) []layout.FlexChild {
	buttonWidth := 100
	buttonLength := 50
	menuItems := *menuItemButtons // need to dereference first

	menuItemLayouts := []layout.FlexChild{}

	if len(*menuItemButtons) > 0 {
		for _, menuItem := range menuItems {
			menuItemLayouts = append(menuItemLayouts,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(MARGIN).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						for menuItem.Clicked(gtx) {
							fmt.Println("Clicked!")
						}

						gtx.Constraints.Min.X = buttonWidth
						gtx.Constraints.Min.Y = buttonLength
						gtx.Constraints.Max = gtx.Constraints.Min

						btn := material.Button(theme, menuItem, "Menu Item")
						btn.Color = Black
						btn.Background = Blue
						return btn.Layout(gtx)
					})
				}))
		}
	}
	return menuItemLayouts
}

func addMenuItems(menuItemButtons *[]*widget.Clickable) {
	menuItemButton := new(widget.Clickable)
	*menuItemButtons = append(*menuItemButtons, menuItemButton)
	fmt.Println(menuItemButtons)

}

func renderLayout(gtx layout.Context, theme *material.Theme, addMenuItemButton *widget.Clickable, menuItems *[]*widget.Clickable) {
	layout.UniformInset(MARGIN).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
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
			layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: 34}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Background{}.Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							return drawBox(gtx, gtx.Constraints.Max.X, gtx.Constraints.Max.Y, Black)
						},
						func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
								layout.Flexed(3, func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
										renderMenuItemButtons(theme, menuItems)...,
									)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{Axis: layout.Vertical, Alignment: layout.End}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return renderAddMenuItemButton(gtx, theme, addMenuItemButton, menuItems)
										}),
									)
								}),
							)
						},
					)
				})
			}),
		)
	})
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	addMenuItemButton := new(widget.Clickable)
	menuItems := []*widget.Clickable{}
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			renderLayout(gtx, theme, addMenuItemButton, &menuItems)
			e.Frame(gtx.Ops)
		}
	}
}

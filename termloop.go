package termloop

import (
	"github.com/gdamore/tcell"
	"strings"
)

// A Canvas is a 2D array of Cells, used for drawing.
// The structure of a Canvas is an array of columns.
// This is so it can be addressed canvas[x][y].
type Canvas [][]Cell

// NewCanvas returns a new Canvas, with
// width and height defined by arguments.
func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]Cell, height)
	}
	return canvas
}

func (canvas *Canvas) equals(oldCanvas *Canvas) bool {
	c := *canvas
	c2 := *oldCanvas
	if c2 == nil {
		return false
	}
	if len(c) != len(c2) {
		return false
	}
	if len(c[0]) != len(c2[0]) {
		return false
	}
	for i := range c {
		for j := range c[i] {
			equal := c[i][j].equals(&(c2[i][j]))
			if !equal {
				return false
			}
		}
	}
	return true
}

// CanvasFromString returns a new Canvas, built from
// the characters in the string str. Newline characters in
// the string are interpreted as a new Canvas row.
func CanvasFromString(str string) Canvas {
	lines := strings.Split(str, "\n")
	runes := make([][]rune, len(lines))
	width := 0
	for i := range lines {
		runes[i] = []rune(lines[i])
		width = max(width, len(runes[i]))
	}
	height := len(runes)
	canvas := make(Canvas, width)
	for i := 0; i < width; i++ {
		canvas[i] = make([]Cell, height)
		for j := 0; j < height; j++ {
			if i < len(runes[j]) {
				canvas[i][j] = Cell{Ch: runes[j][i]}
			}
		}
	}
	return canvas
}

// Drawable represents something that can be drawn, and placed in a Level.
type Drawable interface {
	Tick(Event)   // Method for processing events, e.g. input
	Draw(*Screen) // Method for drawing to the screen
}

// Physical represents something that can collide with another
// Physical, but cannot process its own collisions.
// Optional addition to Drawable.
type Physical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
}

// DynamicPhysical represents something that can process its own collisions.
// Implementing this is an optional addition to Drawable.
type DynamicPhysical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
	Collide(Physical)     // Handle collisions with another Physical
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abstract Termbox stuff for convenience - users
// should only need Termloop imported

// Represents a character to be drawn on the screen.
type Cell struct {
	Fg Attr // Foreground colour
	Bg Attr // Background color
	Ch rune // The character to draw
}

func (c *Cell) equals(c2 *Cell) bool {
	return c.Fg == c2.Fg &&
		c.Bg == c2.Bg &&
		c.Ch == c2.Ch
}

// Provides an event, for input, errors or resizing.
// Resizing and errors are largely handled by Termloop itself
// - this would largely be used for input.
type Event struct {
	Type   EventType     // The type of event
	Key    Key           // The key pressed, if any
	Ch     rune          // The character of the key, if any
	Mod    tcell.ModMask // A keyboard modifier, if any
	Err    error         // Error, if any
	MouseX int           // Mouse X coordinate, if any
	MouseY int           // Mouse Y coordinate, if any
}

func convertEvent(rawEv interface{}) Event {
	var result Event
	switch ev := rawEv.(type) {
	case *tcell.EventKey:
		result.Type = EventMouse
		result.Key = Key(ev.Key())
		result.Ch = ev.Rune()
		result.Mod = ev.Modifiers()

	case *tcell.EventMouse:
		result.Type = EventMouse
		result.Key = Key(ev.Buttons())
		result.MouseX, result.MouseY = ev.Position()

	}
	return result
}

type (
	Attr      uint16
	Key       tcell.Key
	Modifier  uint8
	EventType uint8
)

// Types of event. For example, a keyboard press will be EventKey.
const (
	EventKey EventType = iota
	EventResize
	EventMouse
	EventError
	EventInterrupt
	EventRaw
	EventNone
)

// Cell colors. You can combine these with multiple attributes using
// a bitwise OR ('|'). Colors can't combine with other colors.
const (
	ColorBlack Attr = iota
	ColorMaroon
	ColorGreen
	ColorOlive
	ColorNavy
	ColorPurple
	ColorTeal
	ColorSilver
	ColorGray
	ColorRed
	ColorLime
	ColorYellow
	ColorBlue
	ColorFuchsia
	ColorAqua
	ColorWhite
	Color16
	Color17
	Color18
	Color19
	Color20
	Color21
	Color22
	Color23
	Color24
	Color25
	Color26
	Color27
	Color28
	Color29
	Color30
	Color31
	Color32
	Color33
	Color34
	Color35
	Color36
	Color37
	Color38
	Color39
	Color40
	Color41
	Color42
	Color43
	Color44
	Color45
	Color46
	Color47
	Color48
	Color49
	Color50
	Color51
	Color52
	Color53
	Color54
	Color55
	Color56
	Color57
	Color58
	Color59
	Color60
	Color61
	Color62
	Color63
	Color64
	Color65
	Color66
	Color67
	Color68
	Color69
	Color70
	Color71
	Color72
	Color73
	Color74
	Color75
	Color76
	Color77
	Color78
	Color79
	Color80
	Color81
	Color82
	Color83
	Color84
	Color85
	Color86
	Color87
	Color88
	Color89
	Color90
	Color91
	Color92
	Color93
	Color94
	Color95
	Color96
	Color97
	Color98
	Color99
	Color100
	Color101
	Color102
	Color103
	Color104
	Color105
	Color106
	Color107
	Color108
	Color109
	Color110
	Color111
	Color112
	Color113
	Color114
	Color115
	Color116
	Color117
	Color118
	Color119
	Color120
	Color121
	Color122
	Color123
	Color124
	Color125
	Color126
	Color127
	Color128
	Color129
	Color130
	Color131
	Color132
	Color133
	Color134
	Color135
	Color136
	Color137
	Color138
	Color139
	Color140
	Color141
	Color142
	Color143
	Color144
	Color145
	Color146
	Color147
	Color148
	Color149
	Color150
	Color151
	Color152
	Color153
	Color154
	Color155
	Color156
	Color157
	Color158
	Color159
	Color160
	Color161
	Color162
	Color163
	Color164
	Color165
	Color166
	Color167
	Color168
	Color169
	Color170
	Color171
	Color172
	Color173
	Color174
	Color175
	Color176
	Color177
	Color178
	Color179
	Color180
	Color181
	Color182
	Color183
	Color184
	Color185
	Color186
	Color187
	Color188
	Color189
	Color190
	Color191
	Color192
	Color193
	Color194
	Color195
	Color196
	Color197
	Color198
	Color199
	Color200
	Color201
	Color202
	Color203
	Color204
	Color205
	Color206
	Color207
	Color208
	Color209
	Color210
	Color211
	Color212
	Color213
	Color214
	Color215
	Color216
	Color217
	Color218
	Color219
	Color220
	Color221
	Color222
	Color223
	Color224
	Color225
	Color226
	Color227
	Color228
	Color229
	Color230
	Color231
	Color232
	Color233
	Color234
	Color235
	Color236
	Color237
	Color238
	Color239
	Color240
	Color241
	Color242
	Color243
	Color244
	Color245
	Color246
	Color247
	Color248
	Color249
	Color250
	Color251
	Color252
	Color253
	Color254
	Color255
	ColorAliceBlue
	ColorAntiqueWhite
	ColorAquaMarine
	ColorAzure
	ColorBeige
	ColorBisque
	ColorBlanchedAlmond
	ColorBlueViolet
	ColorBrown
	ColorBurlyWood
	ColorCadetBlue
	ColorChartreuse
	ColorChocolate
	ColorCoral
	ColorCornflowerBlue
	ColorCornsilk
	ColorCrimson
	ColorDarkBlue
	ColorDarkCyan
	ColorDarkGoldenrod
	ColorDarkGray
	ColorDarkGreen
	ColorDarkKhaki
	ColorDarkMagenta
	ColorDarkOliveGreen
	ColorDarkOrange
	ColorDarkOrchid
	ColorDarkRed
	ColorDarkSalmon
	ColorDarkSeaGreen
	ColorDarkSlateBlue
	ColorDarkSlateGray
	ColorDarkTurquoise
	ColorDarkViolet
	ColorDeepPink
	ColorDeepSkyBlue
	ColorDimGray
	ColorDodgerBlue
	ColorFireBrick
	ColorFloralWhite
	ColorForestGreen
	ColorGainsboro
	ColorGhostWhite
	ColorGold
	ColorGoldenrod
	ColorGreenYellow
	ColorHoneydew
	ColorHotPink
	ColorIndianRed
	ColorIndigo
	ColorIvory
	ColorKhaki
	ColorLavender
	ColorLavenderBlush
	ColorLawnGreen
	ColorLemonChiffon
	ColorLightBlue
	ColorLightCoral
	ColorLightCyan
	ColorLightGoldenrodYellow
	ColorLightGray
	ColorLightGreen
	ColorLightPink
	ColorLightSalmon
	ColorLightSeaGreen
	ColorLightSkyBlue
	ColorLightSlateGray
	ColorLightSteelBlue
	ColorLightYellow
	ColorLimeGreen
	ColorLinen
	ColorMediumAquamarine
	ColorMediumBlue
	ColorMediumOrchid
	ColorMediumPurple
	ColorMediumSeaGreen
	ColorMediumSlateBlue
	ColorMediumSpringGreen
	ColorMediumTurquoise
	ColorMediumVioletRed
	ColorMidnightBlue
	ColorMintCream
	ColorMistyRose
	ColorMoccasin
	ColorNavajoWhite
	ColorOldLace
	ColorOliveDrab
	ColorOrange
	ColorOrangeRed
	ColorOrchid
	ColorPaleGoldenrod
	ColorPaleGreen
	ColorPaleTurquoise
	ColorPaleVioletRed
	ColorPapayaWhip
	ColorPeachPuff
	ColorPeru
	ColorPink
	ColorPlum
	ColorPowderBlue
	ColorRebeccaPurple
	ColorRosyBrown
	ColorRoyalBlue
	ColorSaddleBrown
	ColorSalmon
	ColorSandyBrown
	ColorSeaGreen
	ColorSeashell
	ColorSienna
	ColorSkyblue
	ColorSlateBlue
	ColorSlateGray
	ColorSnow
	ColorSpringGreen
	ColorSteelBlue
	ColorTan
	ColorThistle
	ColorTomato
	ColorTurquoise
	ColorViolet
	ColorWheat
	ColorWhiteSmoke
	ColorYellowGreen

	ColorDefault = ColorBlack
)

// Cell attributes. These can be combined with OR.
const (
	AttrBold Attr = 1 << (iota + 9)
	AttrUnderline
	AttrReverse
)

const ModAltModifier = 0x01




const (
	KeyNUL Key = iota
	KeySOH
	KeySTX
	KeyETX
	KeyEOT
	KeyENQ
	KeyACK
	KeyBEL
	KeyBS
	KeyTAB
	KeyLF
	KeyVT
	KeyFF
	KeyCR
	KeySO
	KeySI
	KeyDLE
	KeyDC1
	KeyDC2
	KeyDC3
	KeyDC4
	KeyNAK
	KeySYN
	KeyETB
	KeyCAN
	KeyEM
	KeySUB
	KeyESC
	KeyFS
	KeyGS
	KeyRS
	KeyUS
	KeyDEL Key = 0x7F
)

// These keys are aliases for other names.
const (
	KeyBackspace  = KeyBS
	KeyTab        = KeyTAB
	KeyEsc        = KeyESC
	KeyEscape     = KeyESC
	KeyEnter      = KeyCR
	KeyBackspace2 = KeyDEL
)


// These are the control keys.  Note that they overlap with other keys,
// perhaps.  For example, KeyCtrlH is the same as KeyBackspace.
const (
	KeyCtrlSpace Key = iota
	KeyCtrlA
	KeyCtrlB
	KeyCtrlC
	KeyCtrlD
	KeyCtrlE
	KeyCtrlF
	KeyCtrlG
	KeyCtrlH
	KeyCtrlI
	KeyCtrlJ
	KeyCtrlK
	KeyCtrlL
	KeyCtrlM
	KeyCtrlN
	KeyCtrlO
	KeyCtrlP
	KeyCtrlQ
	KeyCtrlR
	KeyCtrlS
	KeyCtrlT
	KeyCtrlU
	KeyCtrlV
	KeyCtrlW
	KeyCtrlX
	KeyCtrlY
	KeyCtrlZ
	KeyCtrlLeftSq // Escape
	KeyCtrlBackslash
	KeyCtrlRightSq
	KeyCtrlCarat
	KeyCtrlUnderscore
)

const (
	KeyRune Key = iota + 256
	KeyArrowUp
	KeyArrowDown
	KeyArrowRight
	KeyArrowLeft
	KeyUpLeft
	KeyUpRight
	KeyDownLeft
	KeyDownRight
	KeyCenter
	KeyPgUp
	KeyPgDn
	KeyHome
	KeyEnd
	KeyInsert
	KeyDelete
	KeyHelp
	KeyExit
	KeyClear
	KeyCancel
	KeyPrint
	KeyPause
	KeyBacktab
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyF26
	KeyF27
	KeyF28
	KeyF29
	KeyF30
	KeyF31
	KeyF32
	KeyF33
	KeyF34
	KeyF35
	KeyF36
	KeyF37
	KeyF38
	KeyF39
	KeyF40
	KeyF41
	KeyF42
	KeyF43
	KeyF44
	KeyF45
	KeyF46
	KeyF47
	KeyF48
	KeyF49
	KeyF50
	KeyF51
	KeyF52
	KeyF53
	KeyF54
	KeyF55
	KeyF56
	KeyF57
	KeyF58
	KeyF59
	KeyF60
	KeyF61
	KeyF62
	KeyF63
	KeyF64


	MouseLeft         = KeyF63 // arbitrary assignments
	MouseRight        = KeyF62
	MouseMiddle       = KeyF61
	MouseRelease      = KeyF60
	MouseWheelUp      = KeyF59
	MouseWheelDown    = KeyF58
)
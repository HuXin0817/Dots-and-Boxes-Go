package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/color"
	"math"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	widget2 "fyne.io/x/fyne/widget"
	"github.com/HuXin0817/dots-and-boxes/server/api"
	"github.com/HuXin0817/dots-and-boxes/src/ai"
	"github.com/HuXin0817/dots-and-boxes/src/audio"
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/mock"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	svg "github.com/ajstarks/svgo"
	"github.com/bytedance/sonic"
)

const (
	ServerAddr = "127.0.0.1:8080"

	EdgeWidth    = 80
	DotWidth     = EdgeWidth / 5
	HalfDotWidth = DotWidth / 2
	BoxWidth     = EdgeWidth - DotWidth
	MinMargin    = EdgeWidth / 3 * 2
)

var (
	LightThemeGameInterfaceBackGroundColor = &color.NRGBA{R: 242, G: 242, B: 242, A: 255}
	DarkThemeGameInterfaceBackGroundColor  = &color.NRGBA{R: 43, G: 43, B: 43, A: 255}
	LightThemeEmptyEdgeColor               = &color.NRGBA{R: 217, G: 217, B: 217, A: 255}
	DarkThemeEmptyEdgeColor                = &color.NRGBA{R: 65, G: 65, B: 65, A: 255}
	Player1EdgeHighlightColor              = &color.NRGBA{R: 64, G: 64, B: 255, A: 255}
	Player2EdgeHighlightColor              = &color.NRGBA{R: 255, G: 64, B: 64, A: 255}
	LightThemeScoreableTipsColor           = &color.NRGBA{R: 250, G: 250, B: 200, A: 255}
	DarkThemeScoreableTipsColor            = &color.NRGBA{R: 65, G: 65, B: 15, A: 255}
	Player1BoxFilledColor                  = &color.NRGBA{R: 64, G: 64, B: 255, A: 64}
	Player2BoxFilledColor                  = &color.NRGBA{R: 255, G: 64, B: 64, A: 64}
	DarkThemeStartMenuBackgroundColor      = &color.NRGBA{R: 16, G: 16, B: 16, A: 255}
	LightThemeStartMenuBackgroundColor     = &color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	LinkColor                              = &color.NRGBA{R: 73, G: 148, B: 236, A: 255}

	//go:embed "assets/font/Times New Roman Bold.ttf"
	TimesNewRomanBold []byte
	//go:embed "assets/font/Times New Roman Italic.ttf"
	TimesNewRomanItalic []byte
	//go:embed "assets/icon/icon.png"
	IconPng []byte
	//go:embed "assets/icon/spinner_dark.gif"
	SpinnerDarkGif []byte
	//go:embed "assets/icon/spinner_light.gif"
	SpinnerLightGif []byte
	//go:embed "assets/music/EnterGame.MP3"
	EnterGameMP3 []byte
	//go:embed "assets/music/NormalMove.MP3"
	NormalMoveMP3 []byte
	//go:embed "assets/music/Score.MP3"
	ScoreMP3 []byte
	//go:embed "assets/music/TouchButton.MP3"
	TouchButtonMP3 []byte
	//go:embed "assets/music/Win.MP3"
	WinMP3 []byte

	IconResource                = fyne.NewStaticResource("Icon", IconPng)
	TimesNewRomanItalicResource = fyne.NewStaticResource("TimesNewRomanItalic", TimesNewRomanItalic)
	TimesNewRomanBoldResource   = fyne.NewStaticResource("TimesNewRomanBold", TimesNewRomanBold)
	SpinnerDarkResource         = fyne.NewStaticResource("SpinnerDark", SpinnerDarkGif)
	SpinnerLightResource        = fyne.NewStaticResource("SpinnerLight", SpinnerLightGif)

	MyID       *uint64
	Cli        = api.New(ServerAddr)
	MainWindow = app.NewWithID("io.github.dotsandboxes").NewWindow("Dots and Boxes")
)

type Config struct {
	AI1     bool
	AI2     bool
	AI1Name string
	AI2Name string
}

var ConfigFilePath = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "config.json"
	}
	dir := filepath.Join(home, ".dots-and-boxes")
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "config.json"
		}
	} else if err != nil {
		return "config.json"
	}
	return filepath.Join(dir, "config.json")
}()

func (c *Config) Save() error {
	if runtime.GOOS == `js` {
		return nil
	}
	content, err := sonic.Marshal(c)
	if err != nil {
		return err
	}
	if err = os.WriteFile(ConfigFilePath, content, 0644); err != nil {
		return err
	}
	return nil
}

var defaultConf = Config{
	AI1:     false,
	AI2:     false,
	AI1Name: "L4",
	AI2Name: "L4",
}

var Conf = func() (c Config) {
	content, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return defaultConf
	}
	if err = sonic.Unmarshal(content, &c); err != nil {
		return defaultConf
	}
	return c
}()

func InterpolationColor(start, end color.NRGBA, f float32) *color.NRGBA {
	return &color.NRGBA{
		R: uint8(float32(start.R) + f*(float32(end.R)-float32(start.R))),
		G: uint8(float32(start.G) + f*(float32(end.G)-float32(start.G))),
		B: uint8(float32(start.B) + f*(float32(end.B)-float32(start.B))),
		A: uint8(float32(start.A) + f*(float32(end.A)-float32(start.A))),
	}
}

func GetColor(lightColor, darkColor *color.NRGBA) *color.NRGBA {
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		return darkColor
	} else {
		return lightColor
	}
}

func BackGroundColor() *color.NRGBA {
	return GetColor(LightThemeGameInterfaceBackGroundColor, DarkThemeGameInterfaceBackGroundColor)
}

func EmptyEdgeColor() *color.NRGBA {
	return GetColor(LightThemeEmptyEdgeColor, DarkThemeEmptyEdgeColor)
}

func BoxTipsColor() *color.NRGBA {
	return GetColor(LightThemeScoreableTipsColor, DarkThemeScoreableTipsColor)
}

func StartMenuBackgroundColor() *color.NRGBA {
	return GetColor(LightThemeStartMenuBackgroundColor, DarkThemeStartMenuBackgroundColor)
}

func TextColor() color.Color {
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		return color.White
	} else {
		return color.Black
	}
}

func BoxFilledColor(turn model.Turn) *color.NRGBA {
	if turn == model.Player1Turn {
		return Player1BoxFilledColor
	} else {
		return Player2BoxFilledColor
	}
}

func EdgeHighlightColor(turn model.Turn) *color.NRGBA {
	if turn == model.Player1Turn {
		return Player1EdgeHighlightColor
	} else {
		return Player2EdgeHighlightColor
	}
}

func CircleResource() fyne.Resource {
	var buf bytes.Buffer
	c := svg.New(&buf)
	c.Start(200, 200)
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		c.Circle(100, 100, 80, "fill:rgb(202, 202, 202)")
	} else {
		c.Circle(100, 100, 80, "fill:rgb(255, 255, 255)")
	}
	c.End()
	return fyne.NewStaticResource("Icon", buf.Bytes())
}

func SpinnerGIFResource() fyne.Resource {
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		return SpinnerDarkResource
	} else {
		return SpinnerLightResource
	}
}

type CenterLayout struct {
	Min fyne.Size
}

func (l *CenterLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	px := (size.Width - l.Min.Width) / 2
	py := (size.Height - l.Min.Height) / 2
	for _, w := range objects {
		if c, ok := w.(*fyne.Container); ok {
			c.Move(fyne.NewPos(px, py))
			c.Refresh()
		} else if r, ok := w.(*canvas.Rectangle); ok {
			r.Resize(size)
			r.Move(fyne.NewPos(0, 0))
		}
	}
}

func (l *CenterLayout) MinSize([]fyne.CanvasObject) fyne.Size { return l.Min }

type fyneDefaultTheme struct{}

func (t *fyneDefaultTheme) Icon(icon fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(icon)
}

func (t *fyneDefaultTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *fyneDefaultTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

type GameTheme struct {
	game    *GameInterface
	variant fyne.ThemeVariant
	fyneDefaultTheme
}

func (t *GameTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if t.variant != variant {
		t.variant = variant
		for d := range model.MaxDot {
			if t.game.DotCanvases[d] == nil {
				break
			}
			t.game.DotCanvases[d].Resource = CircleResource()
			t.game.DotCanvases[d].Refresh()
		}
		for e := range model.MaxEdge {
			if t.game.EdgeCanvases[e] == nil {
				break
			}
			if t.game.Board.NotContains(e) {
				t.game.EdgeCanvases[e].StrokeColor = EmptyEdgeColor()
			}
		}
	}

	switch name {
	case theme.ColorNameButton:
		return &color.NRGBA{}
	case theme.ColorNameBackground:
		return BackGroundColor()
	}
	return theme.DefaultTheme().Color(name, variant)
}

type StartMenuTheme struct {
	fyneDefaultTheme
	variant fyne.ThemeVariant
	gif     *widget2.AnimatedGif
	title   *canvas.Text
}

func (t *StartMenuTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if t.variant != variant && t.gif != nil {
		t.variant = variant
		t.gif.Stop()
		if err := t.gif.LoadResource(SpinnerGIFResource()); err != nil {
			dialog.NewError(err, MainWindow).Show()
		}
		t.gif.Start()
		t.gif.Refresh()

		t.title.Color = TextColor()
		t.title.Refresh()
	}
	if name == theme.ColorNameBackground {
		return StartMenuBackgroundColor()
	}
	if name == theme.ColorNameButton {
		return EmptyEdgeColor()
	}
	return theme.DefaultTheme().Color(name, variant)
}

func ConvertPosition(x int) float32 { return MinMargin + float32(x)*EdgeWidth }

func NewDotCanvas(d model.Dot) *canvas.Image {
	r := canvas.NewImageFromResource(CircleResource())
	r.Resize(fyne.NewSize(DotWidth*math.Sqrt2, DotWidth*math.Sqrt2))
	x := ConvertPosition(d.X()) - DotWidth*(math.Sqrt2-1)/2
	y := ConvertPosition(d.Y()) - DotWidth*(math.Sqrt2-1)/2
	r.Move(fyne.NewPos(x, y))
	return r
}

func NewEdgeCanvas(e model.Edge, c color.Color) *canvas.Line {
	x1 := ConvertPosition(e.Dot1().X()) + HalfDotWidth
	y1 := ConvertPosition(e.Dot1().Y()) + HalfDotWidth
	x2 := ConvertPosition(e.Dot2().X()) + HalfDotWidth
	y2 := ConvertPosition(e.Dot2().Y()) + HalfDotWidth
	nc := *c.(*color.NRGBA)
	l := canvas.NewLine(&nc)
	l.Position1 = fyne.NewPos(x1, y1)
	l.Position2 = fyne.NewPos(x2, y2)
	l.StrokeWidth = DotWidth
	return l
}

func NewBoxCanvas(b model.Box) *canvas.Rectangle {
	d := b.LeftTopDot()
	x := ConvertPosition(d.X()) + DotWidth - 1
	y := ConvertPosition(d.Y()) + DotWidth - 1
	r := canvas.NewRectangle(&color.NRGBA{})
	r.Move(fyne.NewPos(x, y))
	r.Resize(fyne.NewSize(BoxWidth+1, BoxWidth+1))
	return r
}

func NewButtonCanvas(e model.Edge, tapped func()) *widget.Button {
	b := widget.NewButton("", tapped)
	var s fyne.Size
	if e.Dot1().X() == e.Dot2().X() {
		s = fyne.NewSize(DotWidth, EdgeWidth)
	} else {
		s = fyne.NewSize(EdgeWidth, DotWidth)
	}
	b.Resize(s)
	x1 := ConvertPosition(e.Dot1().X())
	x2 := ConvertPosition(e.Dot2().X())
	y1 := ConvertPosition(e.Dot1().Y())
	y2 := ConvertPosition(e.Dot2().Y())
	x := (x1+x2)/2 - s.Width/2 + HalfDotWidth
	y := (y1+y2)/2 - s.Height/2 + HalfDotWidth
	b.Move(fyne.NewPos(x, y))
	return b
}

type GameInterface struct {
	Board         board.V2
	UserInputEdge model.Edge
	LastEdge      model.Edge
	Paused        bool
	Online        bool
	IsFirst       bool
	GameExit      string
	closed        bool
	Container     *fyne.Container
	Dialog        dialog.Dialog
	DotCanvases   [model.MaxDot]*canvas.Image
	EdgeCanvases  [model.MaxEdge]*canvas.Line
	EdgeAnimation [model.MaxEdge]*fyne.Animation
	EdgeLastColor [model.MaxEdge]color.NRGBA
	BoxCanvases   [model.MaxBox]*canvas.Rectangle
	BoxAnimation  [model.MaxBox]*fyne.Animation
	BoxTipped     [model.MaxBox]bool
	Buttons       [model.MaxEdge]*widget.Button
}

func NewGameInterface() *GameInterface {
	game := &GameInterface{
		Board:    *board.NewV2(),
		LastEdge: -1,
	}
	fyne.CurrentApp().Settings().SetTheme(&GameTheme{
		game: game,
	})
	w := container.NewWithoutLayout()
	for b := range model.MaxBox {
		game.BoxCanvases[b] = NewBoxCanvas(b)
		w.Add(game.BoxCanvases[b])
	}
	for e := range model.MaxEdge {
		game.EdgeCanvases[e] = NewEdgeCanvas(e, EmptyEdgeColor())
		w.Add(game.EdgeCanvases[e])
	}
	for e := range model.MaxEdge {
		game.Buttons[e] = NewButtonCanvas(e, func() {
			if game.UserInputEdge == -1 {
				game.UserInputEdge = e
			}
		})
		w.Add(game.Buttons[e])
	}
	for d := range model.MaxDot {
		game.DotCanvases[d] = NewDotCanvas(d)
		w.Add(game.DotCanvases[d])
	}
	game.Container = container.New(&CenterLayout{
		Min: fyne.NewSize(EdgeWidth*model.DotSize+MinMargin, EdgeWidth*model.DotSize+MinMargin),
	}, w)
	go func() {
		<-FadeIn(700*time.Millisecond, 0, game.Container)
		game.UpdateTitle()
	}()
	MainWindow.SetContent(game.Container)
	MainWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyR:
			if game.GameExit != "" {
				return
			}
			if game.Dialog != nil {
				return
			}
			game.Paused = true
			game.Dialog = dialog.NewCustomConfirm("Restart?", "Yes", "No", game.ScoreLabel(), func(b bool) {
				if b {
					game.GameExit = "restarted"
					if game.Online && MyID != nil {
						_ = Cli.DropID(*MyID)
					}
					Restart(game.Online)
				}
				game.Paused = false
			}, MainWindow)
			game.Dialog.SetOnClosed(func() { game.Dialog = nil })
			game.Dialog.Show()
		case fyne.KeyZ:
			if game.GameExit != "" {
				return
			}
			if game.Dialog != nil {
				return
			}
			game.Paused = true
			game.Dialog = dialog.NewCustomConfirm("Return?", "Yes", "No", game.ScoreLabel(), func(b bool) {
				if b {
					game.Container.Hide()
					game.GameExit = "return"
					if game.Online && MyID != nil {
						_ = Cli.DropID(*MyID)
					}
					ShowStartMenu()
				}
				game.Paused = false
			}, MainWindow)
			game.Dialog.SetOnClosed(func() { game.Dialog = nil })
			game.Dialog.Show()
		case fyne.KeySpace:
			if game.GameExit != "" {
				return
			}
			if game.Dialog != nil {
				game.Dialog.Hide()
				if game.Paused {
					game.Paused = false
				}
			} else {
				game.Dialog = dialog.NewCustom("Paused", "Continue", game.ScoreLabel(), MainWindow)
				game.Dialog.SetOnClosed(func() {
					game.Paused = false
				})
				game.Paused = true
				game.Dialog.SetOnClosed(func() { game.Dialog = nil })
				game.Dialog.Show()
			}
		case fyne.KeyQ:
			if game.Dialog != nil {
				return
			}
			game.Paused = true
			game.Dialog = dialog.NewCustomConfirm("Quit?", "Yes", "No", game.ScoreLabel(), func(b bool) {
				if b {
					MainWindow.Close()
					if game.Online && MyID != nil {
						_ = Cli.DropID(*MyID)
					}
				}
				game.Paused = false
			}, MainWindow)
			game.Dialog.SetOnClosed(func() { game.Dialog = nil })
			game.Dialog.Show()
		case fyne.KeyM:
			audio.On = !audio.On
		}
	})
	return game
}

func (game *GameInterface) ScoreLabel() *widget.Label {
	m := fmt.Sprintf("Player 1 Score: %2d\n", game.Board.Player1Score)
	m += fmt.Sprintf("Player 2 Score: %2d\n", game.Board.Player2Score)
	l := widget.NewLabel(m)
	l.Alignment = fyne.TextAlignCenter
	return l
}

func (game *GameInterface) GetUserEdge() model.Edge {
	game.UserInputEdge = -1
	for game.UserInputEdge == -1 {
		runtime.Gosched()
	}
	return game.UserInputEdge
}

func (game *GameInterface) RemoveHighlight() {
	e := game.LastEdge
	if e != -1 {
		game.EdgeAnimation[e].Stop()
		nc := game.EdgeLastColor[e]
		game.EdgeCanvases[e].StrokeColor = &nc
		game.EdgeCanvases[e].Refresh()
		A0 := game.EdgeCanvases[e].StrokeColor.(*color.NRGBA).A
		game.EdgeAnimation[e] = fyne.NewAnimation(300*time.Millisecond, func(f float32) {
			game.EdgeCanvases[e].StrokeColor.(*color.NRGBA).A = A0 - uint8(127.0*f)
			game.EdgeCanvases[e].Refresh()
		})
		game.EdgeAnimation[e].Start()
	}
}

func (game *GameInterface) Add(e model.Edge) {
	defer game.UpdateTitle()
	for game.Paused {
		runtime.Gosched()
		if game.GameExit != "" {
			return
		}
	}
	if game.GameExit != "" {
		return
	}
	for _, box := range model.NearBoxes[e] {
		if game.Board.EdgeCountOfBox[box] == 3 {
			game.BoxAnimation[box].Stop()
			BoxStartColor := *game.BoxCanvases[box].FillColor.(*color.NRGBA)
			BoxGameExitColor := *BoxFilledColor(game.Board.Turn)
			game.BoxAnimation[box] = fyne.NewAnimation(600*time.Millisecond, func(f float32) {
				game.BoxCanvases[box].FillColor = InterpolationColor(BoxStartColor, BoxGameExitColor, f)
				game.BoxCanvases[box].Refresh()
			})
			game.BoxAnimation[box].Start()
		}
	}
	EdgeStartColor := *game.EdgeCanvases[e].StrokeColor.(*color.NRGBA)
	game.EdgeLastColor[e] = *EdgeHighlightColor(game.Board.Turn)
	game.EdgeAnimation[e] = fyne.NewAnimation(300*time.Millisecond, func(f float32) {
		game.EdgeCanvases[e].StrokeColor = InterpolationColor(EdgeStartColor, game.EdgeLastColor[e], f)
		game.EdgeCanvases[e].Refresh()
	})
	game.EdgeAnimation[e].Start()
	game.Buttons[e].Hide()
	game.RemoveHighlight()
	game.LastEdge = e
	game.EdgeCanvases[e].Refresh()
	score := game.Board.Add(e)
	if game.GameExit == "" {
		if score > 0 {
			audio.Play(ScoreMP3)
		} else {
			audio.Play(NormalMoveMP3)
		}
	}
	t := false
	for box := range model.MaxBox {
		if game.Board.EdgeCountOfBox[box] == 3 && !game.BoxTipped[box] {
			t = true
			break
		}
	}
	if t {
		for box := range model.MaxBox {
			if game.Board.EdgeCountOfBox[box] == 3 {
				if game.BoxTipped[box] {
					game.BoxAnimation[box].Stop()
				}
				BoxStartColor := *BackGroundColor()
				BoxGameExitColor := *BoxTipsColor()
				game.BoxAnimation[box] = fyne.NewAnimation(1200*time.Millisecond, func(f float32) {
					game.BoxCanvases[box].FillColor = InterpolationColor(BoxStartColor, BoxGameExitColor, f)
					game.BoxCanvases[box].Refresh()
				})
				game.BoxAnimation[box].AutoReverse = true
				game.BoxAnimation[box].RepeatCount = -1
				game.BoxAnimation[box].Start()
				game.BoxTipped[box] = true
			}
		}
	}
}

func (game *GameInterface) UpdateTitle() {
	if game.Online {
		nowStep := game.Board.Step
		var title string
		if game.Board.Turn == model.Player1Turn {
			if game.IsFirst {
				title = "Your Turn"
			} else {
				title = "Opponent's Turn"
			}
		} else {
			if game.IsFirst {
				title = "Opponent's Turn"
			} else {
				title = "Your Turn"
			}
		}
		go func() {
			t := config.PlayerTimeOut / time.Second
			for range time.Tick(time.Second) {
				if t == 0 || game.GameExit != "" || game.Board.Step != nowStep {
					return
				}
				MainWindow.SetTitle(fmt.Sprintf("%s (%ds)", title, t))
				t--
			}
		}()
	} else {
		if game.Board.Turn == model.Player1Turn {
			MainWindow.SetTitle("Blue Turn")
		} else {
			MainWindow.SetTitle("Red Turn")
		}
	}
}

func (game *GameInterface) Gaming() bool {
	if game.GameExit != "" {
		return false
	}
	return game.Board.NotOver()
}

func (game *GameInterface) Close() {
	if game.closed {
		return
	}
	game.closed = true
	audio.Play(WinMP3)
	MainWindow.SetTitle("dots and boxes")
	game.RemoveHighlight()
	title := ""
	if game.Online {
		if game.GameExit != "" {
			title = game.GameExit
		} else {
			if game.Board.Player1Score > game.Board.Player2Score {
				if game.IsFirst {
					title = "You Win!"
				} else {
					title = "You Lose."
				}
			} else if game.Board.Player1Score < game.Board.Player2Score {
				if game.IsFirst {
					title = "You Lose."
				} else {
					title = "You Win!"
				}
			} else {
				title = "Draw!"
			}
		}
	} else {
		if game.Board.Player1Score > game.Board.Player2Score {
			title = "Player 1 Win!"
		} else if game.Board.Player1Score < game.Board.Player2Score {
			title = "Player 2 Win!"
		} else {
			title = "Draw!"
		}
	}
	tapped := false
	banReturn := false
	GameOverDialog := dialog.NewCustomConfirm(title, "", "Return", game.ScoreLabel(), func(b bool) {
		if tapped {
			return
		}
		tapped = true
		if !b && !banReturn {
			ShowStartMenu()
		} else {
			Restart(game.Online)
		}
	}, MainWindow)
	go func() {
		t := 5
		for t > 0 && GameOverDialog != nil {
			GameOverDialog.SetConfirmText(fmt.Sprintf("Restart (%ds)", t))
			t--
			time.Sleep(time.Second)
		}
		banReturn = true
		if GameOverDialog != nil {
			GameOverDialog.Hide()
		}
	}()
	GameOverDialog.SetOnClosed(func() {
		GameOverDialog = nil
	})
	GameOverDialog.Show()
}

func Fade(d, wait time.Duration, Container *fyne.Container, StartA uint8) <-chan time.Time {
	BackgroundColor := *StartMenuBackgroundColor()
	BackgroundColor.A = StartA
	Rectangle := canvas.NewRectangle(&BackgroundColor)
	Rectangle.Resize(MainWindow.Canvas().Size())
	go func() {
		time.Sleep(wait)
		if StartA == 0 {
			fyne.NewAnimation(d, func(f float32) {
				BackgroundColor.A = uint8(255 * f)
				Rectangle.Refresh()
			}).Start()
		} else {
			fyne.NewAnimation(d, func(f float32) {
				BackgroundColor.A = uint8(255 * (1 - f))
				Rectangle.Refresh()
			}).Start()
		}
	}()
	Container.Add(Rectangle)
	return time.After(d + wait)
}

func FadeOut(d, wait time.Duration, Container *fyne.Container) <-chan time.Time {
	return Fade(d, wait, Container, 0)
}

func FadeIn(d, wait time.Duration, Container *fyne.Container) <-chan time.Time {
	return Fade(d, wait, Container, 255)
}

func ShowIntroduceInterface() {
	MainWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeyM {
			audio.On = !audio.On
		}
	})
	Container := container.NewWithoutLayout()
	MenuSize := fyne.NewSize(612, 600)
	Container.Resize(MenuSize)
	Icon := canvas.NewImageFromResource(IconResource)
	Icon.Resize(fyne.NewSize(150, 150))
	Icon.Move(fyne.NewPos(100, 225))
	Container.Add(Icon)
	Title := canvas.NewText("Dots and Boxes", TextColor())
	Title.Alignment = fyne.TextAlignCenter
	Title.TextSize = 33
	Title.FontSource = TimesNewRomanBoldResource
	Title.Resize(fyne.NewSize(250, 70))
	Title.Move(fyne.NewPos(280, 265))
	Container.Add(Title)
	fyne.CurrentApp().Settings().SetTheme(&StartMenuTheme{
		title: Title,
	})
	MainContainer := container.New(&CenterLayout{
		Min: MenuSize,
	}, Container)
	MainContainer.Resize(MenuSize)
	c := FadeIn(time.Second, time.Second, MainContainer)
	MainWindow.SetContent(MainContainer)
	go func() {
		time.Sleep(800 * time.Millisecond)
		audio.Play(EnterGameMP3)
		<-c
		<-FadeOut(time.Second, 700*time.Millisecond, MainContainer)
		ShowStartMenu()
	}()
}

func ShowStartMenu() {
	MainWindow.SetTitle("Dots and Boxes")
	Container := container.NewWithoutLayout()
	MenuSize := fyne.NewSize(612, 600)
	Container.Resize(MenuSize)
	Icon := canvas.NewImageFromResource(IconResource)
	Icon.Resize(fyne.NewSize(200, 200))
	Icon.Move(fyne.NewPos(206, 50))
	Container.Add(Icon)
	Title := canvas.NewText("Dots and Boxes", TextColor())
	Title.Alignment = fyne.TextAlignCenter
	Title.TextSize = 26
	Title.FontSource = TimesNewRomanBoldResource
	Title.Resize(fyne.NewSize(200, 50))
	Title.Move(fyne.NewPos(206, 250))
	Container.Add(Title)
	Link := canvas.NewText("https://github.com/HuXin0817/Dots-and-Boxes", LinkColor)
	Link.Resize(fyne.NewSize(200, 50))
	Link.Alignment = fyne.TextAlignCenter
	Link.TextSize = 17
	Link.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	Link.FontSource = TimesNewRomanItalicResource
	Link.Move(fyne.NewPos(206, 285))
	Container.Add(Link)
	Spinner, err := widget2.NewAnimatedGifFromResource(SpinnerGIFResource())
	if err != nil {
		dialog.NewError(err, MainWindow).Show()
		return
	}
	Spinner.Start()
	Spinner.Resize(fyne.NewSize(70, 70))
	Spinner.Move(fyne.NewPos(271, 335))
	Container.Add(Spinner)
	fyne.CurrentApp().Settings().SetTheme(&StartMenuTheme{
		gif:   Spinner,
		title: Title,
	})
	started := false
	PlayOnlineButton := widget.NewButton("Play Online", func() {
		if started {
			return
		}
		started = true
		audio.Play(TouchButtonMP3)
		Restart(true)
	})
	PlayOnlineButton.Resize(fyne.NewSize(200, 50))
	PlayOnlineButton.Move(fyne.NewPos(206, 435))
	Container.Add(PlayOnlineButton)
	PlayLocalButton := widget.NewButton("Play Local", func() {
		TmpConfig := Conf
		Player1RadioGroup := widget.NewRadioGroup([]string{
			"People",
			"Computer",
		}, func(selected string) {
			if selected == "People" {
				TmpConfig.AI1 = false
			} else if selected == "Computer" {
				TmpConfig.AI1 = true
			}
		})
		if Conf.AI1 {
			Player1RadioGroup.SetSelected("Computer")
		} else {
			Player1RadioGroup.SetSelected("People")
		}
		Player2RadioGroup := widget.NewRadioGroup([]string{
			"People",
			"Computer",
		}, func(selected string) {
			if selected == "People" {
				TmpConfig.AI2 = false
			} else if selected == "Computer" {
				TmpConfig.AI2 = true
			}
		})
		if Conf.AI2 {
			Player2RadioGroup.SetSelected("Computer")
		} else {
			Player2RadioGroup.SetSelected("People")
		}
		Entry1 := widget.NewEntry()
		Entry1.SetText(Conf.AI1Name)
		Entry2 := widget.NewEntry()
		Entry2.SetText(Conf.AI2Name)
		PlayLocalButtonSettingDialog := dialog.NewForm("Play Local", "Start", "Cancel", []*widget.FormItem{
			widget.NewFormItem("Player1:", Player1RadioGroup),
			widget.NewFormItem("Player2:", Player2RadioGroup),
			widget.NewFormItem("AI1:", Entry1),
			widget.NewFormItem("AI2:", Entry2),
		}, func(b bool) {
			if !b {
				return
			}
			if _, err := ai.New(Entry1.Text); err != nil {
				dialog.NewError(err, MainWindow).Show()
				return
			} else {
				TmpConfig.AI1Name = Entry1.Text
			}
			if _, err := ai.New(Entry2.Text); err != nil {
				dialog.NewError(err, MainWindow).Show()
				return
			} else {
				TmpConfig.AI2Name = Entry2.Text
			}
			Conf = TmpConfig
			if err := Conf.Save(); err != nil {
				dialog.NewError(err, MainWindow).Show()
			}
			if started {
				return
			}
			started = true
			audio.Play(TouchButtonMP3)
			Restart(false)
		}, MainWindow)
		PlayLocalButtonSettingDialog.Resize(fyne.NewSize(300, 360))
		PlayLocalButtonSettingDialog.Show()
	})
	PlayLocalButton.Resize(fyne.NewSize(200, 50))
	PlayLocalButton.Move(fyne.NewPos(206, 495))
	Container.Add(PlayLocalButton)
	MainContainer := container.New(&CenterLayout{
		Min: MenuSize,
	}, Container)
	MainContainer.Resize(MenuSize)
	FadeIn(time.Second, 300*time.Millisecond, MainContainer)
	MainWindow.SetContent(MainContainer)
	MainWindow.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyM {
			audio.On = !audio.On
		}
	})
}

func Restart(Online bool) {
	MyID = nil
	go func() {
		Container := MainWindow.Content().(*fyne.Container)
		<-FadeOut(700*time.Millisecond, 0, Container)
		if Online {
			SpinnerCanvas, err := widget2.NewAnimatedGifFromResource(SpinnerGIFResource())
			if err != nil {
				dialog.NewError(err, MainWindow).Show()
				return
			}
			SpinnerCanvas.Start()
			SpinnerCanvas.Resize(fyne.NewSize(70, 70))
			wait := true
			Size := Container.Size()
			MatchingTimeText := canvas.NewText("Matching", LinkColor)
			MatchingTimeText.Resize(fyne.NewSize(200, 50))
			MatchingTimeText.Alignment = fyne.TextAlignCenter
			MatchingTimeText.TextSize = 20
			MatchingTimeText.FontSource = TimesNewRomanBoldResource
			Size = Container.Size()
			SpinnerCanvas.Move(fyne.NewPos(Size.Width/2-35, Size.Height/2-55))
			MatchingTimeText.Move(fyne.NewPos(Size.Height/2-100, Size.Height/2+15))
			Container.Add(SpinnerCanvas)
			Container.Add(MatchingTimeText)
			go func() {
				startTime := time.Now()
				for range time.Tick(10 * time.Millisecond) {
					if !wait {
						return
					}
					Size = Container.Size()
					SpinnerCanvas.Move(fyne.NewPos(Size.Width/2-35, Size.Height/2-55))
					MatchingTimeText.Move(fyne.NewPos(Size.Width/2-100, Size.Height/2+15))
					t := time.Since(startTime)
					MatchingTimeText.Text = fmt.Sprintf("Matching: %d min %d s", int(t.Minutes()), int(t.Seconds())%60)
					SpinnerCanvas.Refresh()
					MatchingTimeText.Refresh()
				}
			}()
			id, err := Cli.StartGame()
			MyID = &id
			cancel := false
			MainWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
				if ev.Name == fyne.KeyZ {
					cancel = true
					_ = Cli.DropID(id)
					ShowStartMenu()
				}
			})
			if err != nil {
				d := dialog.NewError(err, MainWindow)
				d.SetOnClosed(ShowStartMenu)
				d.Show()
				return
			}
			isFirst, err := Cli.WaitJoin(id, &cancel)
			if cancel {
				return
			}
			if err != nil {
				d := dialog.NewError(err, MainWindow)
				d.SetOnClosed(ShowStartMenu)
				d.Show()
				return
			}
			wait = false
			UI := NewGameInterface()
			userAdd := func() model.Edge {
				e := UI.GetUserEdge()
				gameExit, err := Cli.AddEdge(id, e)
				if err != nil {
					d := dialog.NewError(err, MainWindow)
					d.SetOnClosed(ShowStartMenu)
					d.Show()
					return 0
				}
				if gameExit != "" {
					UI.GameExit = gameExit
					UI.Close()
				}
				return e
			}
			UI.IsFirst = isFirst
			UI.Online = Online
			enemyAdd := func() model.Edge {
				edge, gameExit, err := Cli.GetOnlinePlayerEdge(id, UI.Board.Step)
				if err != nil {
					d := dialog.NewError(err, MainWindow)
					d.SetOnClosed(ShowStartMenu)
					d.Show()
				}
				if gameExit != "" {
					UI.GameExit = gameExit
					UI.Close()
				}
				return edge
			}
			if isFirst {
				go func() {
					mock.Run(&UI.Board, UI.Gaming, UI.Add, userAdd, enemyAdd)
					UI.Close()
				}()
			} else {
				go func() {
					mock.Run(&UI.Board, UI.Gaming, UI.Add, enemyAdd, userAdd)
					UI.Close()
				}()
			}
		} else {
			AI1, err := ai.New(Conf.AI1Name)
			if err != nil {
				dialog.NewError(err, MainWindow).Show()
			}
			AI2, err := ai.New(Conf.AI2Name)
			if err != nil {
				dialog.NewError(err, MainWindow).Show()
			}
			UI := NewGameInterface()
			func1 := UI.GetUserEdge
			func2 := UI.GetUserEdge
			if Conf.AI1 {
				func1 = func() model.Edge { return AI1(&UI.Board) }
			}
			if Conf.AI2 {
				func2 = func() model.Edge { return AI2(&UI.Board) }
			}
			go func() {
				mock.Run(&UI.Board, UI.Gaming, UI.Add, func1, func2)
				UI.Close()
			}()
		}
	}()
}

var handled bool

func HandleSignal() {
	if handled {
		return
	}
	handled = true
	if MyID != nil {
		_ = Cli.DropID(*MyID)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGABRT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		sig := <-sigs
		HandleSignal()
		os.Exit(int(sig.(syscall.Signal)))
	}()
	ShowIntroduceInterface()
	MainWindow.SetOnClosed(HandleSignal)
	MainWindow.ShowAndRun()
}

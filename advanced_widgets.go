package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// CustomHangmanWidget draws the hangman using canvas graphics instead of ASCII
type CustomHangmanWidget struct {
	widget.BaseWidget
	attempts int
	renderer *customHangmanRenderer
}

func NewCustomHangmanWidget() *CustomHangmanWidget {
	w := &CustomHangmanWidget{attempts: 6}
	w.ExtendBaseWidget(w)
	return w
}

func (w *CustomHangmanWidget) SetAttempts(attempts int) {
	w.attempts = attempts
	w.Refresh()
}

func (w *CustomHangmanWidget) CreateRenderer() fyne.WidgetRenderer {
	w.renderer = &customHangmanRenderer{widget: w}
	return w.renderer
}

type customHangmanRenderer struct {
	widget  *CustomHangmanWidget
	objects []fyne.CanvasObject
}

func (r *customHangmanRenderer) Layout(size fyne.Size) {
	// Layout the canvas objects within the given size
	for _, obj := range r.objects {
		obj.Resize(size)
		obj.Move(fyne.NewPos(0, 0))
	}
}

func (r *customHangmanRenderer) MinSize() fyne.Size {
	return fyne.NewSize(200, 250)
}

func (r *customHangmanRenderer) Refresh() {
	r.objects = r.createHangmanDrawing()
	canvas.Refresh(r.widget)
}

func (r *customHangmanRenderer) Objects() []fyne.CanvasObject {
	if r.objects == nil {
		r.objects = r.createHangmanDrawing()
	}
	return r.objects
}

func (r *customHangmanRenderer) Destroy() {}

func (r *customHangmanRenderer) createHangmanDrawing() []fyne.CanvasObject {
	objects := []fyne.CanvasObject{}

	// Base (always visible)
	base := canvas.NewLine(color.RGBA{R: 139, G: 69, B: 19, A: 255}) // Brown
	base.StrokeWidth = 4
	base.Position1 = fyne.NewPos(50, 240)
	base.Position2 = fyne.NewPos(150, 240)
	objects = append(objects, base)

	// Pole (always visible)
	pole := canvas.NewLine(color.RGBA{R: 139, G: 69, B: 19, A: 255})
	pole.StrokeWidth = 4
	pole.Position1 = fyne.NewPos(80, 240)
	pole.Position2 = fyne.NewPos(80, 20)
	objects = append(objects, pole)

	// Top beam (always visible)
	topBeam := canvas.NewLine(color.RGBA{R: 139, G: 69, B: 19, A: 255})
	topBeam.StrokeWidth = 4
	topBeam.Position1 = fyne.NewPos(80, 20)
	topBeam.Position2 = fyne.NewPos(140, 20)
	objects = append(objects, topBeam)

	// Noose (always visible)
	noose := canvas.NewLine(color.RGBA{R: 139, G: 69, B: 19, A: 255})
	noose.StrokeWidth = 2
	noose.Position1 = fyne.NewPos(140, 20)
	noose.Position2 = fyne.NewPos(140, 40)
	objects = append(objects, noose)

	// Draw hangman parts based on wrong attempts (6 - attempts)
	wrongAttempts := 6 - r.widget.attempts

	if wrongAttempts >= 1 {
		// Head
		head := canvas.NewCircle(color.RGBA{R: 255, G: 220, B: 177, A: 255}) // Skin color
		head.StrokeColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
		head.StrokeWidth = 2
		head.Resize(fyne.NewSize(30, 30))
		head.Move(fyne.NewPos(125, 40))
		objects = append(objects, head)
	}

	if wrongAttempts >= 2 {
		// Body
		body := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
		body.StrokeWidth = 3
		body.Position1 = fyne.NewPos(140, 70)
		body.Position2 = fyne.NewPos(140, 150)
		objects = append(objects, body)
	}

	if wrongAttempts >= 3 {
		// Left arm
		leftArm := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
		leftArm.StrokeWidth = 2
		leftArm.Position1 = fyne.NewPos(140, 90)
		leftArm.Position2 = fyne.NewPos(120, 110)
		objects = append(objects, leftArm)
	}

	if wrongAttempts >= 4 {
		// Right arm
		rightArm := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
		rightArm.StrokeWidth = 2
		rightArm.Position1 = fyne.NewPos(140, 90)
		rightArm.Position2 = fyne.NewPos(160, 110)
		objects = append(objects, rightArm)
	}

	if wrongAttempts >= 5 {
		// Left leg
		leftLeg := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
		leftLeg.StrokeWidth = 2
		leftLeg.Position1 = fyne.NewPos(140, 150)
		leftLeg.Position2 = fyne.NewPos(120, 180)
		objects = append(objects, leftLeg)
	}

	if wrongAttempts >= 6 {
		// Right leg
		rightLeg := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
		rightLeg.StrokeWidth = 2
		rightLeg.Position1 = fyne.NewPos(140, 150)
		rightLeg.Position2 = fyne.NewPos(160, 180)
		objects = append(objects, rightLeg)
	}

	return objects
}

// AnimatedProgressBar shows progress with smooth animations
type AnimatedProgressBar struct {
	widget.BaseWidget
	progress       float64
	targetProgress float64
	animating      bool
	renderer       *animatedProgressRenderer
}

func NewAnimatedProgressBar() *AnimatedProgressBar {
	w := &AnimatedProgressBar{progress: 0.0, targetProgress: 0.0}
	w.ExtendBaseWidget(w)
	return w
}

func (w *AnimatedProgressBar) SetValue(value float64) {
	w.targetProgress = value
	if !w.animating {
		w.animateToTarget()
	}
}

func (w *AnimatedProgressBar) animateToTarget() {
	w.animating = true
	go func() {
		for math.Abs(w.progress-w.targetProgress) > 0.01 {
			diff := w.targetProgress - w.progress
			w.progress += diff * 0.1 // Smooth animation
			w.Refresh()
			time.Sleep(16 * time.Millisecond) // ~60 FPS
		}
		w.progress = w.targetProgress
		w.animating = false
		w.Refresh()
	}()
}

func (w *AnimatedProgressBar) CreateRenderer() fyne.WidgetRenderer {
	w.renderer = &animatedProgressRenderer{widget: w}
	return w.renderer
}

type animatedProgressRenderer struct {
	widget     *AnimatedProgressBar
	background *canvas.Rectangle
	fill       *canvas.Rectangle
	text       *canvas.Text
}

func (r *animatedProgressRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.background.Move(fyne.NewPos(0, 0))

	fillWidth := float32(r.widget.progress) * size.Width
	r.fill.Resize(fyne.NewSize(fillWidth, size.Height))
	r.fill.Move(fyne.NewPos(0, 0))

	r.text.Move(fyne.NewPos(size.Width/2-20, size.Height/2-10))
}

func (r *animatedProgressRenderer) MinSize() fyne.Size {
	return fyne.NewSize(200, 20)
}

func (r *animatedProgressRenderer) Refresh() {
	fillWidth := float32(r.widget.progress) * r.background.Size().Width
	r.fill.Resize(fyne.NewSize(fillWidth, r.background.Size().Height))

	percentage := int(r.widget.progress * 100)
	r.text.Text = fmt.Sprintf("%d%%", percentage)
	canvas.Refresh(r.widget)
}

func (r *animatedProgressRenderer) Objects() []fyne.CanvasObject {
	if r.background == nil {
		r.background = canvas.NewRectangle(color.RGBA{R: 200, G: 200, B: 200, A: 255})
		r.fill = canvas.NewRectangle(color.RGBA{R: 76, G: 175, B: 80, A: 255}) // Green
		r.text = canvas.NewText("0%", color.RGBA{R: 0, G: 0, B: 0, A: 255})
		r.text.Alignment = fyne.TextAlignCenter
	}
	return []fyne.CanvasObject{r.background, r.fill, r.text}
}

func (r *animatedProgressRenderer) Destroy() {}

// CustomButton with hover effects and animations
type CustomButton struct {
	widget.BaseWidget
	text     string
	onTapped func()
	hovered  bool
	pressed  bool
	renderer *customButtonRenderer
}

func NewCustomButton(text string, onTapped func()) *CustomButton {
	b := &CustomButton{text: text, onTapped: onTapped}
	b.ExtendBaseWidget(b)
	return b
}

func (b *CustomButton) Tapped(*fyne.PointEvent) {
	if b.onTapped != nil {
		b.onTapped()
	}
}

func (b *CustomButton) MouseIn(*fyne.PointEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *CustomButton) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (b *CustomButton) MouseDown(*fyne.PointEvent) {
	b.pressed = true
	b.Refresh()
}

func (b *CustomButton) MouseUp(*fyne.PointEvent) {
	b.pressed = false
	b.Refresh()
}

func (b *CustomButton) CreateRenderer() fyne.WidgetRenderer {
	b.renderer = &customButtonRenderer{button: b}
	return b.renderer
}

type customButtonRenderer struct {
	button     *CustomButton
	background *canvas.Rectangle
	text       *canvas.Text
}

func (r *customButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.background.Move(fyne.NewPos(0, 0))

	textSize := r.text.MinSize()
	r.text.Move(fyne.NewPos((size.Width-textSize.Width)/2, (size.Height-textSize.Height)/2))
}

func (r *customButtonRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 40)
}

func (r *customButtonRenderer) Refresh() {
	// Change colors based on state
	if r.button.pressed {
		r.background.FillColor = color.RGBA{R: 100, G: 100, B: 100, A: 255} // Dark gray
	} else if r.button.hovered {
		r.background.FillColor = color.RGBA{R: 180, G: 180, B: 180, A: 255} // Light gray
	} else {
		r.background.FillColor = color.RGBA{R: 150, G: 150, B: 150, A: 255} // Medium gray
	}
	canvas.Refresh(r.button)
}

func (r *customButtonRenderer) Objects() []fyne.CanvasObject {
	if r.background == nil {
		r.background = canvas.NewRectangle(color.RGBA{R: 150, G: 150, B: 150, A: 255})
		r.text = canvas.NewText(r.button.text, color.RGBA{R: 0, G: 0, B: 0, A: 255})
		r.text.Alignment = fyne.TextAlignCenter
		r.text.TextStyle = fyne.TextStyle{Bold: true}
	}
	return []fyne.CanvasObject{r.background, r.text}
}

func (r *customButtonRenderer) Destroy() {}

// WordRevealWidget animates letter reveals
type WordRevealWidget struct {
	widget.BaseWidget
	word     string
	revealed []bool
	renderer *wordRevealRenderer
}

func NewWordRevealWidget(word string) *WordRevealWidget {
	w := &WordRevealWidget{
		word:     word,
		revealed: make([]bool, len(word)),
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *WordRevealWidget) RevealLetter(letter rune) {
	for i, char := range w.word {
		if char == letter {
			w.revealed[i] = true
		}
	}
	w.Refresh()
}

func (w *WordRevealWidget) RevealAll() {
	for i := range w.revealed {
		w.revealed[i] = true
	}
	w.Refresh()
}

func (w *WordRevealWidget) Reset(newWord string) {
	w.word = newWord
	w.revealed = make([]bool, len(newWord))
	w.Refresh()
}

func (w *WordRevealWidget) CreateRenderer() fyne.WidgetRenderer {
	w.renderer = &wordRevealRenderer{widget: w}
	return w.renderer
}

type wordRevealRenderer struct {
	widget     *WordRevealWidget
	letters    []*canvas.Text
	underlines []*canvas.Line
}

func (r *wordRevealRenderer) Layout(size fyne.Size) {
	if len(r.widget.word) == 0 {
		return
	}

	letterWidth := size.Width / float32(len(r.widget.word))
	letterHeight := size.Height

	for i := range r.widget.word {
		x := float32(i) * letterWidth

		if i < len(r.letters) {
			r.letters[i].Move(fyne.NewPos(x+letterWidth/4, letterHeight/4))
		}

		if i < len(r.underlines) {
			r.underlines[i].Position1 = fyne.NewPos(x+5, letterHeight-10)
			r.underlines[i].Position2 = fyne.NewPos(x+letterWidth-5, letterHeight-10)
		}
	}
}

func (r *wordRevealRenderer) MinSize() fyne.Size {
	return fyne.NewSize(float32(len(r.widget.word)*30), 50)
}

func (r *wordRevealRenderer) Refresh() {
	r.createLetterObjects()
	canvas.Refresh(r.widget)
}

func (r *wordRevealRenderer) Objects() []fyne.CanvasObject {
	if r.letters == nil {
		r.createLetterObjects()
	}

	objects := []fyne.CanvasObject{}
	for _, letter := range r.letters {
		objects = append(objects, letter)
	}
	for _, underline := range r.underlines {
		objects = append(objects, underline)
	}
	return objects
}

func (r *wordRevealRenderer) createLetterObjects() {
	r.letters = []*canvas.Text{}
	r.underlines = []*canvas.Line{}

	for i, char := range r.widget.word {
		// Create letter text
		var letterText string
		if r.widget.revealed[i] || char == ' ' {
			letterText = string(char)
		} else {
			letterText = ""
		}

		letter := canvas.NewText(letterText, color.RGBA{R: 0, G: 0, B: 0, A: 255})
		letter.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}
		letter.TextSize = 24
		r.letters = append(r.letters, letter)

		// Create underline (only for letters, not spaces)
		if char != ' ' {
			underline := canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
			underline.StrokeWidth = 2
			r.underlines = append(r.underlines, underline)
		} else {
			// Add empty underline to maintain indexing
			r.underlines = append(r.underlines, canvas.NewLine(color.Transparent))
		}
	}
}

func (r *wordRevealRenderer) Destroy() {}

// ScoreWidget with animated score changes
type ScoreWidget struct {
	widget.BaseWidget
	score        int
	displayScore int
	animating    bool
	renderer     *scoreRenderer
}

func NewScoreWidget() *ScoreWidget {
	w := &ScoreWidget{score: 0, displayScore: 0}
	w.ExtendBaseWidget(w)
	return w
}

func (w *ScoreWidget) SetScore(score int) {
	w.score = score
	if !w.animating {
		w.animateScore()
	}
}

func (w *ScoreWidget) animateScore() {
	w.animating = true
	go func() {
		for w.displayScore != w.score {
			if w.displayScore < w.score {
				w.displayScore++
			} else {
				w.displayScore--
			}
			w.Refresh()
			time.Sleep(50 * time.Millisecond)
		}
		w.animating = false
	}()
}

func (w *ScoreWidget) CreateRenderer() fyne.WidgetRenderer {
	w.renderer = &scoreRenderer{widget: w}
	return w.renderer
}

type scoreRenderer struct {
	widget     *ScoreWidget
	text       *canvas.Text
	background *canvas.Rectangle
}

func (r *scoreRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.background.Move(fyne.NewPos(0, 0))

	textSize := r.text.MinSize()
	r.text.Move(fyne.NewPos((size.Width-textSize.Width)/2, (size.Height-textSize.Height)/2))
}

func (r *scoreRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 30)
}

func (r *scoreRenderer) Refresh() {
	r.text.Text = fmt.Sprintf("🏆 %d", r.widget.displayScore)
	canvas.Refresh(r.widget)
}

func (r *scoreRenderer) Objects() []fyne.CanvasObject {
	if r.text == nil {
		r.background = canvas.NewRectangle(color.RGBA{R: 255, G: 215, B: 0, A: 100}) // Gold background
		r.text = canvas.NewText("🏆 0", color.RGBA{R: 0, G: 0, B: 0, A: 255})
		r.text.TextStyle = fyne.TextStyle{Bold: true}
		r.text.TextSize = 18
	}
	return []fyne.CanvasObject{r.background, r.text}
}

func (r *scoreRenderer) Destroy() {}

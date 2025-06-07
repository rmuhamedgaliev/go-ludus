package hello

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			PaddingTop(2).
			PaddingLeft(4).
			Width(22)

	nameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 2)
)

type App struct {
	input  io.Reader
	output io.Writer
}

func NewApp() *App {
	return &App{
		input:  os.Stdin,
		output: os.Stdout,
	}
}

func NewAppWithIO(input io.Reader, output io.Writer) *App {
	return &App{
		input:  input,
		output: output,
	}
}

func (a *App) ReadInput(prompt string) (string, error) {
	fmt.Fprint(a.output, prompt)
	reader := bufio.NewReader(a.input)
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func (a *App) PrintError(message string) {
	fmt.Fprintln(a.output, errorStyle.Render(message))
}

func (a *App) SayHello(name string) {
	title := titleStyle.Render("Привет!")
	greeting := fmt.Sprintf("Добро пожаловать, %s!", nameStyle.Render(name))
	info := infoStyle.Render(greeting)

	fmt.Fprintln(a.output, title)
	fmt.Fprintln(a.output, info)
}

func (a *App) ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("имя не может быть пустым")
	}
	if len(name) > 100 {
		return fmt.Errorf("имя слишком длинное (максимум 100 символов)")
	}
	return nil
}

func (a *App) Run() error {
	name, err := a.ReadInput("Введите ваше имя: ")
	if err != nil {
		a.PrintError(fmt.Sprintf("Ошибка при чтении ввода: %v", err))
		return err
	}

	if err := a.ValidateName(name); err != nil {
		a.PrintError(err.Error())
		return err
	}

	a.SayHello(name)
	return nil
}

func ReadInput(prompt string) (string, error) {
	app := NewApp()
	return app.ReadInput(prompt)
}

func PrintError(message string) {
	app := NewApp()
	app.PrintError(message)
}

func SayHello(name string) {
	app := NewApp()
	app.SayHello(name)
}

func Run() {
	app := NewApp()
	app.Run()
}

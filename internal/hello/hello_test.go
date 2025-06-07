package hello

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestApp_ReadInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		prompt   string
		expected string
		hasError bool
	}{
		{
			name:     "valid input with newline",
			input:    "Test User\n",
			prompt:   "Enter your name: ",
			expected: "Test User",
			hasError: false,
		},
		{
			name:     "empty input",
			input:    "\n",
			prompt:   "Enter your name: ",
			expected: "",
			hasError: false,
		},
		{
			name:     "input with leading/trailing spaces",
			input:    "  John Doe  \n",
			prompt:   "Enter your name: ",
			expected: "John Doe",
			hasError: false,
		},
		{
			name:     "input without newline (EOF case)",
			input:    "Test User",
			prompt:   "Enter your name: ",
			expected: "Test User",
			hasError: false,
		},
		{
			name:     "empty input without newline",
			input:    "",
			prompt:   "Enter your name: ",
			expected: "",
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}
			app := NewAppWithIO(input, output)

			result, err := app.ReadInput(tt.prompt)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %q, got %q", tt.expected, result)
				}
			}

			outputStr := output.String()
			if !strings.Contains(outputStr, tt.prompt) {
				t.Errorf("Expected output to contain prompt %q, got %q", tt.prompt, outputStr)
			}
		})
	}
}

func TestApp_PrintError(t *testing.T) {
	tests := []struct {
		name          string
		errorMessage  string
		shouldContain []string
	}{
		{
			name:          "simple error message",
			errorMessage:  "Test error message",
			shouldContain: []string{"Test error message"},
		},
		{
			name:          "russian error message",
			errorMessage:  "Имя не может быть пустым",
			shouldContain: []string{"Имя не может быть пустым"},
		},
		{
			name:          "formatted error message",
			errorMessage:  "Ошибка при чтении ввода: invalid input",
			shouldContain: []string{"Ошибка при чтении ввода", "invalid input"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader("")
			output := &bytes.Buffer{}
			app := NewAppWithIO(input, output)

			app.PrintError(tt.errorMessage)

			outputStr := output.String()
			for _, expected := range tt.shouldContain {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Expected output to contain %q, got %q", expected, outputStr)
				}
			}
		})
	}
}

func TestApp_SayHello(t *testing.T) {
	tests := []struct {
		name          string
		inputName     string
		shouldContain []string
	}{
		{
			name:          "valid name",
			inputName:     "Test User",
			shouldContain: []string{"Привет!", "Добро пожаловать", "Test User"},
		},
		{
			name:          "name with special characters",
			inputName:     "Анна-Мария",
			shouldContain: []string{"Привет!", "Добро пожаловать", "Анна-Мария"},
		},
		{
			name:          "single character name",
			inputName:     "A",
			shouldContain: []string{"Привет!", "Добро пожаловать", "A"},
		},
		{
			name:          "name with unicode",
			inputName:     "🎉 Test",
			shouldContain: []string{"Привет!", "Добро пожаловать", "🎉 Test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader("")
			output := &bytes.Buffer{}
			app := NewAppWithIO(input, output)

			app.SayHello(tt.inputName)

			outputStr := output.String()
			for _, expected := range tt.shouldContain {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Expected output to contain %q, got %q", expected, outputStr)
				}
			}
		})
	}
}

func TestApp_ValidateName(t *testing.T) {
	tests := []struct {
		name      string
		inputName string
		hasError  bool
		errorMsg  string
	}{
		{
			name:      "valid name",
			inputName: "John Doe",
			hasError:  false,
		},
		{
			name:      "empty name",
			inputName: "",
			hasError:  true,
			errorMsg:  "имя не может быть пустым",
		},
		{
			name:      "only spaces",
			inputName: "   ",
			hasError:  true,
			errorMsg:  "имя не может быть пустым",
		},
		{
			name:      "too long name",
			inputName: strings.Repeat("A", 101),
			hasError:  true,
			errorMsg:  "имя слишком длинное",
		},
		{
			name:      "exactly 100 characters",
			inputName: strings.Repeat("A", 100),
			hasError:  false,
		},
		{
			name:      "name with leading/trailing spaces",
			inputName: "  John  ",
			hasError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := NewApp()
			err := app.ValidateName(tt.inputName)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				} else if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error to contain %q, got %q", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestApp_Run(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectError      bool
		shouldContain    []string
		shouldNotContain []string
	}{
		{
			name:          "successful run",
			input:         "John Doe\n",
			expectError:   false,
			shouldContain: []string{"Привет!", "Добро пожаловать", "John Doe", "Введите ваше имя"},
		},
		{
			name:             "empty name",
			input:            "\n",
			expectError:      true,
			shouldContain:    []string{"имя не может быть пустым", "Введите ваше имя"},
			shouldNotContain: []string{"Привет!", "Добро пожаловать"},
		},
		{
			name:             "name with spaces only",
			input:            "   \n",
			expectError:      true,
			shouldContain:    []string{"имя не может быть пустым"},
			shouldNotContain: []string{"Привет!", "Добро пожаловать"},
		},
		{
			name:             "too long name",
			input:            strings.Repeat("A", 101) + "\n",
			expectError:      true,
			shouldContain:    []string{"имя слишком длинное"},
			shouldNotContain: []string{"Привет!", "Добро пожаловать"},
		},
		{
			name:          "valid russian name",
			input:         "Анна\n",
			expectError:   false,
			shouldContain: []string{"Привет!", "Добро пожаловать", "Анна"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}
			app := NewAppWithIO(input, output)

			err := app.Run()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}

			outputStr := output.String()

			for _, expected := range tt.shouldContain {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Expected output to contain %q, got %q", expected, outputStr)
				}
			}

			for _, notExpected := range tt.shouldNotContain {
				if strings.Contains(outputStr, notExpected) {
					t.Errorf("Expected output NOT to contain %q, but it does. Got %q", notExpected, outputStr)
				}
			}
		})
	}
}

func TestGlobalFunctions(t *testing.T) {
	t.Run("global SayHello", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Global SayHello panicked: %v", r)
			}
		}()

		SayHello("Test")
	})

	t.Run("global PrintError", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Global PrintError panicked: %v", r)
			}
		}()

		PrintError("Test error")
	})
}

func BenchmarkApp_SayHello(b *testing.B) {
	input := strings.NewReader("")
	output := io.Discard
	app := NewAppWithIO(input, output)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.SayHello("Test User")
	}
}

func BenchmarkApp_ValidateName(b *testing.B) {
	app := NewApp()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.ValidateName("Test User")
	}
}

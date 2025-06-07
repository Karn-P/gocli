package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var ttsCmd = &cobra.Command{
	Use:   "tts",
	Short: "Text to Speech",
	Long: `Convert text to speech using TTS engine.
	For example:
	gocli tts "Hello, world!"
	gocli tts -v Alex -r 180 "This is a spoken with a different voice and rate."`,
	Run: textToSpeech,
}

func init() {
	rootCmd.AddCommand(ttsCmd)

	ttsCmd.Flags().StringP("voice", "v", "", "Voice to use for speech (macOS voices like Alex, Samantha, etc.)")
	ttsCmd.Flags().IntP("rate", "r", 175, "Speech rate (words per minute)")
	ttsCmd.Flags().StringP("output", "o", "", "Save speech to audio file (AIFF format)")
}

func textToSpeech(cmd *cobra.Command, args []string) {

	text := strings.Join(args, " ")
	voice, _ := cmd.Flags().GetString("voice")
	rate, _ := cmd.Flags().GetInt("rate")
	output, _ := cmd.Flags().GetString("output")

	sayCmd := []string{}

	if voice != "" {
		sayCmd = append(sayCmd, "-v", voice)
	}

	sayCmd = append(sayCmd, "-r", fmt.Sprintf("%d", rate))

	if output != "" {
		sayCmd = append(sayCmd, "-o", output)
	}

	sayCmd = append(sayCmd, text)

	fmt.Printf("Speaking: %s\n", text)
	cmd2 := exec.Command("say", sayCmd...)
	err := cmd2.Run()
	if err != nil {
		fmt.Printf("Error speaking text: %v\n", err)
	}
}

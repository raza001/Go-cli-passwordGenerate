package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var generatecmd = &cobra.Command{
	Use:   "generate",
	Short: "generate random password",
	Long: `Generate random password with customize options.
	For Example:

	cobra generate  -l 12 -d -s`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generatecmd)
	generatecmd.Flags().IntP("length", "l", 8, "Length  of the generate")
	generatecmd.Flags().BoolP("digits", "d", false, "Include the digits")
	generatecmd.Flags().BoolP("special-chars", "s", false, "Include the special char is here")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")

	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if isDigits {
		charset += "0123456789"
	}
	if isSpecialChars {
		charset += "!@#$%^&*()"
	}
	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Println("Generate Password")
	fmt.Println(string(password))

}

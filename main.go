package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeDuplicatesFromFile(inputFile, outputFile string) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("gagal membuka file input: %w", err)
	}
	defer inFile.Close()

	uniqueNumbers := make(map[string]struct{})
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() { // Baca file secara linear
		number := scanner.Text()
		uniqueNumbers[number] = struct{}{} // Tidak perlu lock
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("kesalahan saat membaca file input: %w", err)
	}

	var outputBuilder strings.Builder
	for number := range uniqueNumbers {
		outputBuilder.WriteString(number)
		outputBuilder.WriteString("\n")
	}
	outputString := outputBuilder.String()

	err = os.WriteFile(outputFile, []byte(outputString), 0644)
	if err != nil {
		return fmt.Errorf("gagal menulis ke file output: %w", err)
	}

	return nil
}

func main() {
	inputFile := "input.txt"
	outputFile := "nomor_unik.txt"

	err := removeDuplicatesFromFile(inputFile, outputFile)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
	} else {
		fmt.Println("Nomor telepon unik telah ditulis ke", outputFile)
	}
}

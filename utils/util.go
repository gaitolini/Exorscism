package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type InputField struct {
	Prompt string
	Type   string
	Value  interface{}
}

func ReadInputs(fields []InputField) error {
	reader := bufio.NewReader(os.Stdin)

	for i, field := range fields {
		fmt.Print(field.Prompt + ": ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)

		switch field.Type {
		case "int":
			intValue, err := strconv.Atoi(input)
			if err != nil {
				return err
			}
			fields[i].Value = intValue
		case "string":
			fields[i].Value = input
		case "float":
			floatValue, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return err
			}
			fields[i].Value = floatValue
		case "date":
			dateValue, err := time.Parse("02/01/2006", input)
			if err != nil {
				return err
			}
			fields[i].Value = dateValue
		default:
			return fmt.Errorf("unsupported input type: %s", field.Type)
		}
	}

	return nil
}

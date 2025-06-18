package object

import (
	"fmt"
	"strings"
)

func GetOwnerAndNameFromId(id string, separator string) (string, string, error) {
	tokens := strings.Split(id, separator)
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("GetOwnerAndNameFromId() error, wrong token count for ID: %s", id)
	}

	return tokens[0], tokens[1], nil
}

func GetIdFromOwnerAndName(owner string, name string, separator string) string {
	return fmt.Sprintf("%s%s%s", owner, separator, name)
}

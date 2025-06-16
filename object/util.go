package object

import (
	"errors"
	"fmt"
	"strings"
)

func GetOwnerAndNameFromId(id string, separator string) (string, string) {
	tokens := strings.Split(id, separator)
	if len(tokens) != 2 {
		panic(errors.New("GetOwnerAndNameFromId() error, wrong token count for ID: " + id))
	}

	return tokens[0], tokens[1]
}

func GetIdFromOwnerAndName(owner string, name string, separator string) string {
	return fmt.Sprintf("%s%s%s", owner, separator, name)
}

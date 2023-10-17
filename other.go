package goutils

import "github.com/google/uuid"

/* return Prestring + _uuidv4 */
func uuidv4gen(preString string) (string, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uuidStr := newUUID.String()
	return preString + uuidStr, err
}

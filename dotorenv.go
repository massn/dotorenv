package dotorenv

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(dotEnvFirst bool, fileNames ...string) error {
	envMap, err := Read(dotEnvFirst, fileNames...)
	if err != nil {
		return err
	}
	for k, v := range envMap {
		os.Setenv(k, v)
	}
	return nil
}

func Read(dotEnvFirst bool, fileNames ...string) (map[string]string, error) {
	dotEnvMap, err := godotenv.Read(fileNames...)
	finalMap := make(map[string]string)
	if err != nil {
		return finalMap, err
	}
	for k, v := range dotEnvMap {
		envVar := os.Getenv(k)
		if envVar == "" {
			finalMap[k] = v
		} else {
			if dotEnvFirst {
				finalMap[k] = v
			} else {
				finalMap[k] = envVar
			}
		}
	}
	return finalMap, nil
}

package reader

import (
	"bufio"
	"os"
)

//*********************************************************

type URLList struct {
	List []string
	Size uint32
}

//*********************************************************

func (u *URLList) Read(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		u.List = append(u.List, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	u.Size = uint32(len(u.List))

	return nil
}

//*********************************************************

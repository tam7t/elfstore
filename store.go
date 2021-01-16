// Package elfstore is a persistent key/value string storage system without
// unnecessary things like a database or a file.
package elfstore

import (
	"debug/elf"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unsafe"
)

var storage = `                                                                `

// Load loads the data in storage into a map.
func Load() (map[string]string, error) {
	d := make(map[string]string)
	if strings.HasPrefix(storage, "{") {
		err := json.Unmarshal([]byte(storage), &d)
		return d, err
	}
	return d, nil
}

// Save keeps data safe.
func Save(data map[string]string) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(b) > MaxSize() {
		return fmt.Errorf("cannot persist data, size exceeds %d", MaxSize())
	}

	// determine elf SEGMENT_START for program
	bin, err := os.OpenFile(os.Args[0], os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	e, err := elf.NewFile(bin)
	if err != nil {
		return err
	}

	var offset int
	if sect := e.Section(".text"); sect != nil {
		offset = int(sect.Addr - sect.Offset)
	} else {
		return errors.New("could not determin offset")
	}

	// follow the pointer to storage to find offset in the file where we can
	// write data.
	offset = *(*int)(unsafe.Pointer(&storage)) - offset

	// read and modify executible
	d, err := ioutil.ReadAll(bin)
	if err != nil {
		return err
	}
	for i := 0; i < len(b); i++ {
		d[offset+i] = b[i]
	}

	// replace binary with modified one
	if err := os.Remove(os.Args[0]); err != nil {
		return err
	}
	return ioutil.WriteFile(os.Args[0], d, 0775)
}

// MaxSize of save data. If you need more recompile with:
//
// -ldflags "-X 'github.com/tam7t/elfstore.storage=$(python3 -c "print(' '*131032)")'"
//
// Note: values over 131032 result in an error:
// bash: /usr/local/go/bin/go: Argument list too long
func MaxSize() int {
	return len(storage)
}

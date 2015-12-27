/**
 * The MIT License (MIT)
 *
 * Copyright (c) 2014 Yani Iliev <yani@iliev.me>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package wpress

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

const (
	headerSize   = 4377 // length of the header
	filenameSize = 255  // maximum number of bytes allowed for filename
	contentSize  = 14   // maximum number of bytes allowed for content size
	mtimeSize    = 12   // maximum number of bytes allowed for last modified date
	prefixSize   = 4096 // maximum number of bytes allowed  for prefix
)

// Header block format of a file
// Field Name    Offset    Length    Contents
// Name               0       255    filename (no path, no slash)
// Size             255        14    length of file contents
// Mtime            269        12    last modification date
// Prefix           281      4096    path name, no trailing slashes
type Header struct {
	Name   []byte
	Size   []byte
	Mtime  []byte
	Prefix []byte
}

// PopulateFromBytes populates header struct from bytes array
func (h *Header) PopulateFromBytes(block []byte) {
	h.Name = block[0:255]
	h.Size = block[255:269]
	h.Mtime = block[269:281]
	h.Prefix = block[281:4377]
}

// PopulateFromFilename populates header struct from passed filename
func (h *Header) PopulateFromFilename(filename string) error {

	// try to open the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	// get the fileinfo
	fi, err := file.Stat()
	if err != nil {
		return err
	}

	// validate if filename fits the allowed length
	if len(fi.Name()) > filenameSize {
		return errors.New("filename is longer than max allowed")
	}
	// create filename buffer
	h.Name = make([]byte, filenameSize)
	// copy filename to the buffer leaving available space as zero-bytes
	copy(h.Name, fi.Name())

	// get filesize as string
	size := strconv.FormatInt(fi.Size(), 10)
	// validate if filesize fits the allowed length
	if len(size) > contentSize {
		return errors.New("file size is larger than max allowed")
	}
	// create size buffer
	h.Size = make([]byte, contentSize)
	// copy content size length to the buffer
	copy(h.Size, size)

	// get last modified date as string
	unixTime := strconv.FormatInt(fi.ModTime().Unix(), 10)
	if len(unixTime) > mtimeSize {
		return errors.New("last modified date is after than max allowed")
	}
	// create mtime buffer
	h.Mtime = make([]byte, mtimeSize)
	// copy mtime to the buffer
	copy(h.Mtime, unixTime)

	// get the path to the file
	_path := filepath.Dir(filename)
	// validate if path fits the allowed length
	if len(_path) > prefixSize {
		return errors.New("prefix size is longer than max allowed")
	}
	// create buffer to put the prefix in
	h.Prefix = make([]byte, prefixSize)
	// put the prefix in the buffer
	copy(h.Prefix, _path)

	// close the file
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// GetHeaderBlock returns byte sequence of header block populated with data
func (h Header) GetHeaderBlock() []byte {
	block := append(h.Name, h.Size...)
	block = append(block, h.Mtime...)
	block = append(block, h.Prefix...)
	return block
}

// GetSize returns content size
func (h Header) GetSize() (int, error) {
	// remove any trailing zero bytes, convert to string, then convert to integer
	return strconv.Atoi(string(bytes.Trim(h.Size, "\x00")))
}

// GetEOFBlock returns byte sequence describing EOF
func (h Header) GetEOFBlock() []byte {
	// generate zero-byte sequence of length headerSize
	return bytes.Repeat([]byte("\x00"), headerSize)
}

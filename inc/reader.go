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
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

const PATH_SEPARATOR_WIN = '\\'
const PATH_SEPARATOR_UNIX = '/'

// Reader structure
type Reader struct {
	Filename      string
	File          *os.File
	NumberOfFiles int
}

// NewReader creates a new Reader instance and calls its constructor
func NewReader(filename string) (*Reader, error) {
	// create a new instance of Reader
	r := &Reader{filename, nil, 0}

	// call the constructor
	err := r.Init()
	if err != nil {
		return nil, err
	}

	// return Reader instance
	return r, nil
}

// Init is the constructor of Reader struct
func (r *Reader) Init() error {
	// try to open the file
	file, err := os.Open(r.Filename)
	if err != nil {
		return err
	}

	// file was openned, assign the handle to the holding variable
	r.File = file

	return nil
}

// ExtractFile extracts file that matches tha filename and path from archive
func (r Reader) ExtractFile(filename string, path string) ([]byte, error) {
	// TODO: implement
	return nil, nil
}

// Extract all files from archive
func (r Reader) Extract() (int, error) {
	// put pointer at the beginning of the file
	r.File.Seek(0, 0)

	// loop until end of file was reached
	iteration := 0;
	for {
		iteration++;
		// read header block
		block, err := r.GetHeaderBlock()
		if err != nil {
			return 0, err
		}

		// initialize new header
		h := &Header{}

		// check if block equals EOF sequence
		if bytes.Compare(block, h.GetEOFBlock()) == 0 {
			// EOF reached, stop the loop
			break
		}

		// populate header from our block bytes
		h.PopulateFromBytes(block)

		pathToFile := path.Clean("." + string(os.PathSeparator) + string(bytes.Trim(h.Prefix, "\x00")) + string(os.PathSeparator) + string(bytes.Trim(h.Name, "\x00")))
		if runtime.GOOS == "windows" {
		    sep := fmt.Sprintf("%c", PATH_SEPARATOR_UNIX)
		    pathToFile = strings.Replace(pathToFile,"\\",sep,-1)
		    fmt.Println(pathToFile)
		}

		err = os.MkdirAll(path.Dir(pathToFile), 0777)
		if err != nil {
			fmt.Println(err)
			return r.NumberOfFiles, err
		}

		// try to open the file
		


		file, err := os.Create(pathToFile)
		if err != nil {
			return r.NumberOfFiles, err
		}

		totalBytesToRead, _ := h.GetSize()
		for {
			bytesToRead := 512
			if bytesToRead > totalBytesToRead {
				bytesToRead = totalBytesToRead
			}

			if bytesToRead == 0 {
				break
			}

			content := make([]byte, bytesToRead)
			bytesRead, err := r.File.Read(content)
			if err != nil {
				return r.NumberOfFiles, err
			}

			totalBytesToRead -= bytesRead
			contentRead := content[0:bytesRead]

			_, err = file.Write(contentRead)
			if err != nil {
				return r.NumberOfFiles, err
			}
		}

		file.Close()

		// increment file counter
		r.NumberOfFiles++
	}

	return r.NumberOfFiles, nil
}

// GetHeaderBlock reads and returns header block from archive
func (r Reader) GetHeaderBlock() ([]byte, error) {
	// create buffer to keep the header block
	block := make([]byte, headerSize)

	// read the header block
	bytesRead, err := r.File.Read(block)
	if err != nil {
		return nil, err
	}

	if bytesRead != headerSize {
		return nil, errors.New("unable to read header block size")
	}

	return block, nil
}

// GetFilesCount returns the number of files in archive
func (r Reader) GetFilesCount() (int, error) {
	// test if we have enumerated the archive already
	if r.NumberOfFiles != 0 {
		return r.NumberOfFiles, nil
	}

	// put pointer at the beginning of the file
	r.File.Seek(0, 0)

	// loop until end of file was reached
	for {
		// read header block
		block, err := r.GetHeaderBlock()
		if err != nil {
			return 0, err
		}

		// initialize new header
		h := &Header{}

		// check if block equals EOF sequence
		if bytes.Compare(block, h.GetEOFBlock()) == 0 {
			// EOF reached, stop the loop
			break
		}

		// populate header from our block bytes
		h.PopulateFromBytes(block)

		// set pointer after file content, to the next header block
		size, err := h.GetSize()
		if err != nil {
			return 0, err
		}
		r.File.Seek(int64(size), 1)

		// increment file counter
		r.NumberOfFiles++
	}

	return r.NumberOfFiles, nil
}

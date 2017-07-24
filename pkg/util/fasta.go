// Copyright 2017 Kenji Kaneda.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

// SequenceSet contains a set of DNA sequences. Constructed from a
// file of the FASTA format (https://en.wikipedia.org/wiki/FASTA_format).
type SequenceSet struct {
	Seqs []string
}

// maxSequenceInFile defines the maximum number of DNA
// sequences defined in one file.
const maxSequencePerFile = 50

//maxLengthPerSeq defines the maximum length of each DNA sequence.
const maxLengthPerSeq = 1000

// LoadFromFile loads a file of the (restricted) FASTA format and
// returns a SequenceSet. It also validates the output.
//
// The file consists of a set of DNA sequences. Each DNA sequence
// starts from the description line beginning with '>'. The actual
// sequence can be one or more than one line.
func LoadFromFile(filename string) (*SequenceSet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Scan each line of the file and construct a SequenceSet.
	scanner := bufio.NewScanner(file)
	seqSet := &SequenceSet{}
	// buf keeps one sequence. It is created when a descriptor line is found
	// and its content is written to seqSet when we reach another descriptor line
	// or the end of the file.
	var buf *bytes.Buffer
	lineNo := 0
	const descriptorCode byte = '>'
	for scanner.Scan() {
		s := scanner.Text()
		lineNo++
		if s[0] == descriptorCode {
			// Found a descriptor line.
			if buf != nil {
				seqSet.Seqs = append(seqSet.Seqs, buf.String())
			}
			buf = new(bytes.Buffer)
		} else {
			// Found a sequence line.
			if buf == nil {
				return nil, fmt.Errorf("sequence started without a descriptor line: line no: %d", lineNo)
			}
			// err is always nil, but checking here to pass errcheck.
			if _, err := buf.WriteString(s); err != nil {
				return nil, err
			}
		}
	}
	if buf != nil {
		seqSet.Seqs = append(seqSet.Seqs, buf.String())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if err := validate(seqSet); err != nil {
		return nil, err
	}
	return seqSet, nil
}

// validate validates the output SequenceSet. Returns an error when
// the validation fails.
func validate(seqSet *SequenceSet) error {
	if len(seqSet.Seqs) == 0 {
		return errors.New("no sequence is found")
	}
	if len(seqSet.Seqs) > maxSequencePerFile {
		return fmt.Errorf("too many sequences are found: %d > %d",
			len(seqSet.Seqs), maxSequencePerFile)
	}

	expectedChars := map[rune]bool{
		'T': true,
		'C': true,
		'G': true,
		'A': true,
	}
	for _, seq := range seqSet.Seqs {
		if len(seq) == 0 {
			return errors.New("sequence must have at least one character")
		}
		if len(seq) > maxLengthPerSeq {
			return fmt.Errorf("too long sequence found: %d > %d", len(seq), maxLengthPerSeq)
		}
		for _, c := range seq {
			if !expectedChars[c] {
				return fmt.Errorf("unexpected character %c", c)
			}
		}
	}
	return nil
}

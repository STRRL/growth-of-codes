package source

import (
	"bufio"
	"github.com/STRRL/growth-of-codes/pkg/persistent/entity"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type CSVSource struct {
	project string
	file    string
	f       *os.File
}

func NewCSVSource(project string, file string) *CSVSource {
	return &CSVSource{project: project, file: file}
}

func (it *CSVSource) Close() error {
	if it.f != nil {
		return it.f.Close()
	}
	return nil
}

func (it *CSVSource) FetchIterator() (Iterator, error) {
	if it.f == nil {
		f, err := os.Open(it.file)
		if err != nil {
			return nil, err
		}
		it.f = f
	}
	return NewCSVIterator(it.project, it.f), nil
}

func (it *CSVSource) ListAll() ([]entity.FileComplexitySnapshot, error) {
	iterator, err := it.FetchIterator()
	if err != nil {
		return nil, err
	}
	var result []entity.FileComplexitySnapshot

	for {
		hasNext, err := iterator.HasNext()
		if err != nil {
			return nil, err
		}
		if !hasNext {
			break
		}
		item, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		result = append(result, *item)
	}
	return result, nil
}

type CSVIterator struct {
	project  string
	upstream io.Reader
	b        *bufio.Reader
	line     string
}

func NewCSVIterator(project string, upstream io.Reader) *CSVIterator {
	return &CSVIterator{project: project, upstream: upstream}
}

func (it *CSVIterator) HasNext() (bool, error) {
	if it.b == nil {
		it.b = bufio.NewReader(it.upstream)
		// discard first line
		it.b.ReadLine()
	}
	temp := ""
	for {
		line, prefix, err := it.b.ReadLine()
		if err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, err
		}
		temp += string(line)
		if !prefix {
			break
		}
	}
	it.line = temp
	return true, nil
}

func (it *CSVIterator) Next() (*entity.FileComplexitySnapshot, error) {
	split := strings.Split(it.line, ",")
	commitTime, err := time.Parse(time.RFC3339, strings.TrimSpace(split[0]))
	if err != nil {
		return nil, err
	}
	language := strings.TrimSpace(split[1])
	filePath := strings.TrimSpace(split[2])
	complexity, err := strconv.Atoi(strings.TrimSpace(split[3]))
	if err != nil {
		return nil, err
	}
	commitHash := strings.TrimSpace(split[4])
	return &entity.FileComplexitySnapshot{
		Project:    it.project,
		CommitHash: commitHash,
		CommitTime: commitTime,
		Language:   language,
		FilePath:   filePath,
		Complexity: uint32(complexity),
	}, nil
}

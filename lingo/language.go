package lingo

import (
	"sort"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type Language struct {
	Name  string
	Files []string
	Count uint32
	Color string
}

func (l *Language) FileCount() int {
	return len(l.Files)
}

func (l *Language) CountLines() {
	for _, fp := range l.Files {
		if strings.HasSuffix(fp, "package-lock.json") {
			continue
		}
		wg.Add(1)
		go l.addLineCount(fp)
	}
}

func (l *Language) addLineCount(fp string) {
	defer wg.Done()
	l.Count += GetLineCount(fp)
}

type Languages map[string]*Language

func (l Languages) FileCount() int {
	total := 0
	for _, lang := range l {
		total += len(lang.Files)
	}

	return total
}

func (l Languages) GetTotal() uint32 {
	var total uint32 = 0
	for _, lang := range l {
		total += lang.Count
	}
	return total
}

func (l Languages) CountLines() {
	for _, lang := range l {
		lang.CountLines()
	}
}

func (l Languages) Wait() {
	wg.Wait()
}

func (l Languages) CountLinesNow() {
	l.CountLines()
	l.Wait()
}

func (l Languages) Slice() LangArray {
	langs := LangArray{}
	for _, lang := range l {
		langs = append(langs, *lang)
	}
	return langs
}

func (l Languages) Sorted() LangArray {
	langs := l.Slice()
	sort.Sort(langs)
	return langs
}

type LangArray []Language

func (l LangArray) Len() int {
	return len(l)
}

func (l LangArray) Less(i, j int) bool {
	return l[i].Count > l[j].Count
}

func (l LangArray) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

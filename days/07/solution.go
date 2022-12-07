package main

import (
	"fmt"
	"strconv"
	strings "strings"

	"github.com/Tch1b0/polaris/array"
	"github.com/Tch1b0/polaris/input"
	"github.com/Tch1b0/polaris/math"
)

type Directory struct {
	Name  string
	Files []File
	Dirs  []Directory
}

// not working correctly right now
func (d Directory) String() string {
	str := fmt.Sprintf("- %s (dir)", d.Name)
	for _, v := range d.Files {
		str += fmt.Sprintf("\n   - %s (file, size=%d)", v.Name, v.Size)
	}

	for _, v := range d.Dirs {
		str += "\n" + strings.ReplaceAll(v.String(), "\n", "\n  ")
	}

	return str
}

func (d Directory) HasDir(name string) bool {
	return array.Some(d.Dirs, func(dir Directory, _ int) bool { return dir.Name == name })
}

func (d Directory) HasFile(name string) bool {
	return array.Some(d.Files, func(file File, _ int) bool { return file.Name == name })
}

func (d Directory) GetDirIndex(name string) int {
	for i, v := range d.Dirs {
		if v.Name == name {
			return i
		}
	}

	return -1
}

func (d *Directory) AddDir(dir Directory) {
	d.Dirs = append(d.Dirs, dir)
}

func (d *Directory) AddFile(file File) {
	d.Files = append(d.Files, file)
}

func (d Directory) GetSize() int {
	sum := 0
	for _, v := range d.Files {
		sum += v.Size
	}

	for _, v := range d.Dirs {
		sum += v.GetSize()
	}

	return sum
}

func (d Directory) GetAllFiles() []File {
	files := d.Files
	for _, dir := range d.Dirs {
		files = append(files, dir.GetAllFiles()...)
	}
	return files
}

func (d Directory) GetSizes() []int {
	sizes := []int{d.GetSize()}
	for _, v := range d.Dirs {
		sizes = append(sizes, v.GetSizes()...)
	}

	return sizes
}

type File struct {
	Name string
	Size int
}

type OS struct {
	Fs         Directory
	CurrentDir []int
}

func (os *OS) Execute(stdio []string) {
	cmds := strings.Split(stdio[0], " ")
	if cmds[1] == "cd" {
		os.cd(cmds[2])
	} else {
		os.ls(stdio[1:])
	}
}

func (os *OS) cd(path string) {
	if path == ".." {
		os.CurrentDir = os.CurrentDir[:len(os.CurrentDir)-1]
	} else if path == "/" {
		os.CurrentDir = []int{}
	} else {
		parent := os.getCurrDir()
		os.checkOrCreateDir(path)

		os.CurrentDir = append(os.CurrentDir, parent.GetDirIndex(path))
	}
}

func (os *OS) checkOrCreateDir(path string) {
	parent := os.getCurrDir()

	if parent.HasDir(path) {
		return
	} else {
		dir := Directory{
			Name: path,
		}
		parent.AddDir(dir)
	}
}

func (os *OS) getCurrDir() *Directory {
	x := &os.Fs
	for _, v := range os.CurrentDir {
		x = &x.Dirs[v]
	}

	return x
}

func (os *OS) ls(stdout []string) {
	dir := os.getCurrDir()
	for _, line := range stdout {
		content := strings.Split(line, " ")
		if content[0] == "dir" {
			os.checkOrCreateDir(content[1])
		} else {
			size, err := strconv.Atoi(content[0])
			if err != nil {
				panic(err)
			}
			file := File{
				Name: content[1],
				Size: size,
			}
			dir.AddFile(file)
		}
	}
}

func NewOS() OS {
	return OS{
		Fs: Directory{
			Name: "/",
		},
	}
}

func getInput() [][]string {
	return input.Process("./days/07/input.txt", func(str string) [][]string {
		splitted := strings.Split(str, "\n")
		stdios := [][]string{}
		for _, v := range splitted {
			if v[0] == '$' {
				stdios = append(stdios, []string{v})
			} else {
				x := &stdios[len(stdios)-1]
				*x = append(*x, v)
			}
		}
		return stdios
	})
}

func sumDir(dir Directory) int {
	sum := 0
	if size := dir.GetSize(); size <= 100000 {
		sum += size
	}

	for _, child := range dir.Dirs {
		sum += sumDir(child)
	}

	return sum
}

func part1() int {
	os := NewOS()
	for _, v := range getInput() {
		os.Execute(v)
	}

	return sumDir(os.Fs)
}

func part2() int {
	os := NewOS()
	for _, v := range getInput() {
		os.Execute(v)
	}

	x := os.Fs.GetSizes()
	dirs := []int{}

	for _, v := range x {
		if 70000000-os.Fs.GetSize()+v >= 30000000 {
			dirs = append(dirs, v)
		}
	}

	return math.Min(dirs)
}

func main() {
	fmt.Printf("PART 1: %d\n", part1())

	fmt.Printf("PART 2: %d\n", part2())
}

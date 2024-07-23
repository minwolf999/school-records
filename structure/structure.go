package structure

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

const IdCtx IDContextKey = "id"

var (
	App        Application
	Tpl        *template.Template
	SeededRand *rand.Rand = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
)

type IDContextKey string

type Application struct {
	Port    string
	Ip      string
	LogPath string
}

type Result struct {
	Success   string
	Error     string
	ImagePath string

	Student     Student
	YearList    []string
	Class       []string
	Categories  []Categorie
	Competence  Competence
	Competences []Competence
}

type Teacher struct {
	Id            string
	Username      string
	Password      string
	Class         string
	SigningUpPath string
	Key           string
}

type Student struct {
	Id    string
	Name  string
	Class string
	Year  string

	Competences []Competence
	Teacher     Teacher
}

type Categorie struct {
	Id    string
	Name  string
	Class string

	SubCategories []SubCategorie
}

type SubCategorie struct {
	Id        string
	Name      string
	Class     string
	Categorie Categorie
}

type Competence struct {
	Id        string
	Name      string
	ImagePath string

	Categorie    Categorie
	SubCategorie SubCategorie

	Students []Student
}

type JustFilesFilesystem struct {
	Fs http.FileSystem
}

func (fs JustFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.Fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, _ := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}

	return f, nil
}

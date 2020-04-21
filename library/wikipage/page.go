package wikipage

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save(path string) error {
	err := ioutil.WriteFile(path+p.Title+".txt", p.Body, 0755)
	return err
}

func ReadPage(path string, title string) (*Page, error) {
	body, err := ioutil.ReadFile(path + title + ".txt")
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

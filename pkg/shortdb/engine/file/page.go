package file

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"

	table "github.com/batazor/shortlink/pkg/shortdb/table/v1"
)

func (f *file) getPage(nameTable string, page int32) error {
	t := f.database.Tables[nameTable]

	// read page
	pageFile, err := f.loadPage(f.pageName(nameTable))
	if err != nil {
		return err
	}

	t.Pages[page] = pageFile
	return nil
}

func (f *file) addPage(nameTable string) (int32, error) {
	t := f.database.Tables[nameTable]

	if t.Stats.RowsCount%t.Option.PageSize == 0 {
		if t.Pages == nil {
			t.Pages = make(map[int32]*table.Page, 0)
		}

		t.Stats.PageCount += 1
		t.Pages[t.Stats.PageCount] = &table.Page{Rows: []*table.Row{}}

		// create a page file
		newPageFile, err := f.createFile(f.pageName(nameTable))
		if err != nil {
			return t.Stats.PageCount, err
		}

		err = newPageFile.Close()
		if err != nil {
			return t.Stats.PageCount, err
		}

		// if this not first page, save current date
		if t.Stats.PageCount > 0 {
			// save data after clear memory page
			err = f.savePage(nameTable, t.Stats.PageCount-1)
			if err != nil {
				return t.Stats.PageCount, err
			}

			// clear old page
			err = f.clearPage(nameTable, t.Stats.PageCount-1)
			if err != nil {
				return t.Stats.PageCount, err
			}
		}
	}

	return t.Stats.PageCount, nil
}

func (f *file) savePage(nameTable string, pageCount int32) error {
	t := f.database.Tables[nameTable]

	// save date
	openFile, err := f.createFile(fmt.Sprintf("%s_%s_%d.page", f.database.Name, nameTable, pageCount))
	if err != nil {
		return err
	}

	defer func() {
		_ = openFile.Close() // #nosec
	}()

	payload, err := proto.Marshal(t.Pages[pageCount])
	if err != nil {
		return err
	}

	// Write something
	err = f.writeFile(openFile.Name(), payload)
	if err != nil {
		return err
	}

	return nil
}

func (f *file) clearPage(nameTable string, pageCount int32) error {
	f.database.Tables[nameTable].Pages[pageCount] = nil
	return nil
}

func (f *file) clearPages(nameTable string) error {
	f.database.Tables[nameTable].Pages = nil
	return nil
}

func (f *file) loadPage(path string) (*table.Page, error) {
	page := table.Page{}

	payload, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(payload) != 0 {
		err = proto.Unmarshal(payload, &page)
		if err != nil {
			return nil, err
		}
	}

	return &page, nil
}

func (f *file) pageName(nameTable string) string {
	return fmt.Sprintf("%s_%s_%d.page", f.database.Name, nameTable, f.database.Tables[nameTable].Stats.PageCount)
}

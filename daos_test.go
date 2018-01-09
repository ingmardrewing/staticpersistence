package staticPersistence

import "testing"

func TestFindVersion_returns_zero_if_no_version_defined(t *testing.T) {
	json := []byte(`{"noversiongiven":""}`)
	expected := 0

	actual := FindJsonVersion(json)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestFindVersion_returns_one_if_version_one_is_defined(t *testing.T) {
	json := []byte(`{"version":1}`)
	expected := 1

	actual := FindJsonVersion(json)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestFindVersion_returns_one_if_new_empty_dao_is_needed(t *testing.T) {
	expected := 1

	actual := FindJsonVersion(nil)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestNewPostDAO_without_json_data_returns_newest_version(t *testing.T) {
	expected := `{
	"version":1,
	"thumbImg":"",
	"postImg":"",
	"filename":"",
	"id":0,
	"date":"",
	"url":"",
	"title":"",
	"title_plain":"",
	"excerpt":"",
	"content":"",
	"dsq_thread_id":""
}`
	d := NewPostDAO(nil, "", "")
	d.Id(0)
	d.Title("")
	d.TitlePlain("")
	d.ThumbUrl("")
	d.ImageUrl("")
	d.Description("")
	d.DisqusId("")
	d.CreateDate("")
	d.Content("")
	d.Url("")
	d.PathFromDocRoot("")
	d.HtmlFilename("")

	actual := string(d.FillJson())

	if actual != expected {
		t.Error("Expected", expected, ", but got", actual)
	}

}

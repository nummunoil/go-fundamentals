package example

import "testing"

type fakeRepo struct {
}

func (fakeRepo) QueryLang(id uint) Language {
	return Language{
		ID:   1,
		Name: "golang",
	}
}

func TestHandle(t *testing.T) {
	h := NewHandler(fakeRepo{})
	s := h.Do(0)

	want := "golang language"
	if s != want {
		t.Errorf("want %q but get %s", want, s)
	}

}

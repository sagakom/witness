package test

import (
	"strconv"
	"testing"
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func TestClickHit(t *testing.T) {
	t.Parallel()

	var expectedRate int64 = 60 // 60% click hit

	chrome, err := chrome.New()
	check(t, err)
	defer chrome.Close()
	page, err := chrome.CDP.DefaultPage()

	get := func(sel string) witness.Element {
		t.Helper()
		el, err := page.Doc().Seek(sel)
		check(t, err)
		return el
	}

	check(t, err)
	check(t, page.Navigate(getFilepath("click_playground.html")))

	target := get("#target")

	var pass int64
	var miss int64
	for i := 0; i < 50; i++ {
		err := target.Click()
		switch err {
		case nil:
			pass++
		case witness.ErrElementMissClick:
			miss++
		default:
			t.Fatal(err)
		}
		time.Sleep(time.Millisecond * 300)
	}

	clickedText, err := target.GetText()
	check(t, err)

	clicked, err := strconv.ParseInt(clickedText, 10, 64)
	check(t, err)

	rate := (100 * pass) / (miss + pass)
	t.Logf("pass = %d, miss = %d, rate = %d", pass, miss, rate)
	if rate <= expectedRate {
		t.Fatalf("miss click degradation - expected at least %d%% success click, but was %d", expectedRate, rate)
	}
	if clicked != pass {
		t.Fatalf("%d flaky clicks", pass-clicked)
	}

}
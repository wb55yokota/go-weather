package libs_test

import (
	"testing"

	"github.com/wb55yokota/go-weather/libs"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func TestCp932(t *testing.T) {
	tests := []struct {
		str string
	}{
		{str: "ascii"},
		{str: "日本語"},
		{str: "①㈱㎝髙﨑"},
	}
	for _, tt := range tests {
		// Cp932して、utf8に戻して、一致するかを確認
		convertedStr, err := libs.Cp932(tt.str)
		if err != nil {
			t.Error("panic.")
		}
		reConvertedStr, _, err := transform.String(japanese.ShiftJIS.NewDecoder(), convertedStr)
		if err != nil {
			t.Error(err)
		}
		if reConvertedStr != tt.str {
			t.Error("convert error.")
		}
		t.Log(reConvertedStr)
	}
}

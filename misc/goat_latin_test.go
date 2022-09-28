package misc

import (
	"testing"
)

func TestGoatLatin(t *testing.T) {
	cases := map[string]struct {
		given string
		want  string
	}{
		"I speak Goat Latin": {
			"I speak Goat Latin",
			"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		"The quick brown fox jumped over the lazy dog": {
			"The quick brown fox jumped over the lazy dog",
			"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
		"Each word consists of lowercase and uppercase letters only": {
			"Each word consists of lowercase and uppercase letters only",
			"Eachmaa ordwmaaa onsistscmaaaa ofmaaaaa owercaselmaaaaaa andmaaaaaaa uppercasemaaaaaaaa etterslmaaaaaaaaa onlymaaaaaaaaaa",
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := ToGoatLatin(tt.given)
			if actual != tt.want {
				t.Errorf("During %s; expected %v and got %v",
					name,
					tt.want,
					actual,
				)
			}
		})
	}
}

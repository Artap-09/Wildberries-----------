package main

import "testing"

var slice = []string{"ва","кУлон", "мостки", "мостки", "стека", "клоун", "автор", "аскет", "ситец", "мостик", "ропак", "ситец", "теска", "истец", "сетка", "тесак", "мостки", "капор", "уклон", "москит", "втора", "кулон", "ситце", "отвар", "колун", "секта", "лукно", "парок", "рвота", "рвота", "копра", "тавро", "товар", "порка"}

func TestAnagram(t *testing.T) {
	result := Anagram(slice)
	if len(result) < 6 {
		t.Error("Нехватает множества")
	}

	for str, val := range result {
		switch str {
		case "ропак":
			if !Equal(*val, []string{"капор", "копра", "парок", "порка", "ропак"}) {
				t.Error(
					"For", slice, "\n",
					"expected", []string{"капор", "копра", "парок", "порка", "ропак"}, "\n",
					"got", val,
				)
			}
		case "кулон":
			if !Equal(*val, []string{"клоун", "колун", "кулон", "лукно", "уклон"}) {
				t.Error(
					"For", slice, "\n",
					"expected", []string{"клоун", "колун", "кулон", "лукно", "уклон"}, "\n",
					"got", val,
				)
			}
		case "мостки":
			if !Equal(*val, []string{"москит", "мостик", "мостки"}) {
				t.Error(
					"For", slice, "\n",
					"expected", []string{"москит", "мостик", "мостки"}, "\n",
					"got", val,
				)
			}
		case "стека":
			if !Equal(*val, []string{"аскет", "секта", "сетка", "стека", "тесак", "теска"}) {
				t.Error(
					"For", slice, "\n",
					"expected", []string{"аскет", "секта", "сетка", "стека", "тесак", "теска"}, "\n",
					"got", val,
				)
			}
		case "автор":
			if !Equal(*val, []string{"автор", "втора", "отвар", "рвота", "тавро", "товар"}) {
				t.Error(
					"For", slice, "\n",
					"expected", []string{"автор", "втора", "отвар", "рвота", "тавро", "товар"}, "\n",
					"got", val,
				)
			}
		case "ситец":
			if !Equal(*val, []string{"истец", "ситец", "ситце"}) {
				t.Error(
					"For", slice, "\n",
					"expected", []string{"истец", "ситец", "ситце"}, "\n",
					"got", val,
				)
			}
		default:
			t.Error("Лишнее множество")
		}
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

package excavator

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TopCallback ...
type TopCallback func(url string, ch *RootRadicalCharacter)

//RootRegular 常用字
func RootRegular(host string, cb TopCallback) error {
	url := strings.Join([]string{host, "z/zb/cc1.htm"}, "/")
	html, e := parseDocument(url)
	if e != nil {
		return e
	}
	html.Find(".bs_index3").Each(func(i int, s1 *goquery.Selection) {
		s1.Find("li").Each(func(i int, s2 *goquery.Selection) {
			a := s2.Find("a").Text()
			link, _ := s2.Find("a").Attr("href")
			pinyin, _ := s2.Find("a").Attr("title")
			cc := RootRadicalCharacter{
				Class:     ClassRegular,
				Character: a,
				Link:      link,
				Pinyin:    strings.Split(pinyin, ","),
			}
			log.Infof("%+v", cc)
			cb(host, &cc)
		})
	})
	return nil
}

func dummyLog(c *StandardCharacter, i int, selection *goquery.Selection) (err error) {
	log.With("index", i).Info(selection.Text())
	return nil
}

func basicExplanation(ex *StandardCharacter, i int, selection *goquery.Selection) (err error) {
	selection.Find("ol li").Each(func(i int, selection *goquery.Selection) {
		log.With("index", i).Info(selection.Text())
		ex.BasicExplanation.BasicMeaning = append(ex.BasicExplanation.BasicMeaning, selection.Text())
	})
	log.Infof("%+v", ex)
	return nil
}

// ProcessFunc ...
type ProcessFunc func(*StandardCharacter, int, *goquery.Selection) error

// DataTypeBlock ...
var DataTypeBlock = map[string]ProcessFunc{
	"基本解释": basicExplanation,
	"详细解释": dummyLog,
	"國語詞典": dummyLog,
	"康熙字典": dummyLog,
	"说文解字": dummyLog,
	"音韵方言": dummyLog,
	"字源字形": dummyLog,
	"网友讨论": dummyLog,
}

//CommonlyBase ...
func CommonlyBase(url string, character *RootRadicalCharacter) {
	url = url + character.Link
	html, e := parseDocument(url)
	if e != nil {
		log.Errorf("%s error with:%s", e, url)
		return
	}
	bc := StandardCharacter{
		Radical:         character.Character,
		CharacterDetail: map[string]string{},
	}

	html.Find("div[data-type-block]").Each(func(i int, s1 *goquery.Selection) {
		log.Info(s1.Html())
		dtb, b := s1.Attr("data-type-block")
		if !b {
			return
		}
		if fn, b := DataTypeBlock[dtb]; b {
			e := fn(&bc, i, s1)
			log.Error(e)
		}

		//log.Info(s1.Text())
		//k := strings.TrimSpace(s1.Text())
		//v, b := s1.Find("data-type-block").Attr("href")
		//if !b || k == "基本解释" {
		//	//基本解释
		//	bc.CharacterDetail[k] = character.Link
		//	return
		//}
		////other
		//bc.CharacterDetail[k] = v
	})
	log.Infof("character detail:%+v", bc)
}

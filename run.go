package excavator

import (
	"fmt"

	"github.com/godcong/go-trait"
)

var log = trait.InitGlobalZapSugar().With("package", "excavator")

func MainRun(url string) error {
	//"http://www.zdic.net/z/zb/cc1.htm"
	url = "http://www.zdic.net/z/zb/cc1.htm"
	chars := CommonlyTop(url)

	fmt.Println("start total:", len(chars))
	return nil
	for idx, v := range chars {
		bc := CommonlyBase("http://www.zdic.net", v)
		//db.DB("base").Insert(bc)
		log.Info(bc)
		fmt.Println("current is :", idx, v.Character)
	}
	fmt.Println("end")
	//rootCmd.AddCommand()
	return nil
}

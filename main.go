package main

import (
	"fmt"

	cld3 "./cld3lib"
)

func main() {
	langId, err := cld3.NewLanguageIdentifier(0, 512)
	if err != nil {
		fmt.Println("whoops, couldn't create a new LanguageIdentifier:", err)
	}
	defer cld3.FreeLanguageIdentifier(langId)
	res := langId.FindLanguage("Hey, this is an english sentence")
	if res.IsReliable {
		fmt.Println("pretty sure we've got text written in", res.Language)
	}
	res = langId.FindLanguage("Guten Tag, wie geht es?")
	if res.IsReliable {
		fmt.Println("ah, and this one is", res.Language)
	}
}

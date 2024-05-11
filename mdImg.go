package main

// example for https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html

import (
	"os"
	"fmt"
	"log"
//	"bytes"

	"github.com/goccy/go-yaml"
	util "github.com/prr123/utility/utilLib"
)

type ImgSum struct {
	MdFile string `yaml:"md file"`
	ImgList []ImgItem `yaml:"image list"`
}

type ImgItem struct {
	Capt string `yaml:"caption"`
	Filnam string `yaml:"file name"`
	Alt string	`yaml:"alt"`
}

func main() {

    numArgs := len(os.Args)

	flags:=[]string{"dbg","md", "img"}

    useStr := "/md=<markdown file> [/img=<img yaml>] [/dbg]"
    helpStr := fmt.Sprintf("help: The program reads an md file and creates a yaml file with a list of images\n")

    if numArgs > len(flags)+1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s %s\n", os.Args[0], useStr)
        os.Exit(1)
    }


    if numArgs == 2 {
        if os.Args[1] == "help" {
            fmt.Printf("usage is: %s %s\n", os.Args[0], useStr)
            fmt.Printf("%s\n", helpStr)
            os.Exit(1)
        }
    }

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    dbg := false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}

    mdFilnam := ""
    outFilnam := ""
    mdval, ok := flagMap["md"]
    if ok {
        if mdval.(string) == "none" {log.Fatalf("error -- no markdown file provided with /md flag!")}
        mdFilnam = mdval.(string)
		outFilnam = "yaml/" + mdFilnam + ".yaml"
		mdFilnam = "mdFiles/" + mdFilnam + ".md"
//		mds = string(md)
	} else {
		log.Fatalf("error -- no md file provided!\n")
	}

    oval, ok := flagMap["img"]
    if ok {
        if oval.(string) == "none" {log.Fatalf("error -- no yaml file provided with /img flag!")}
        outFilnam = oval.(string)
//	idx := bytes.Index(mdFilnam,".md")
		outFilnam += ".yaml"
	}

	mdByt, err := os.ReadFile(mdFilnam)
	if err != nil {log.Fatalf("error -- cannot read md file: %v\n", mdFilnam)}

//	outFilnam = "imgList/" + outFilnam

	outfil, err := os.Create(outFilnam)
	if err != nil {log.Fatalf("error -- cannot create yaml file: %v", err)}
	defer outfil.Close()

    if dbg {
		if len(mdFilnam)  == 0 {
			log.Printf("debug -- no md file!\n")
		} else {
        	log.Printf("debug -- md file: %s\n",mdFilnam)
		}
		log.Printf("debug -- yaml file: %s\n",outFilnam)
    }

	imgList, err := createImgList(mdByt)
	if err != nil {log.Fatalf("error -- creating yaml img list: %v\n", err)}

	for i:=0; i< len(imgList); i++ {
		fmt.Printf(" --%d: %20s %20s %s\n", i, imgList[i].Capt, imgList[i].Alt, imgList[i].Filnam)

	}

	imgSum := ImgSum{
		MdFile: mdFilnam,
		ImgList: imgList,
	}

	imgDat, err := yaml.Marshal(imgSum)
	if err != nil {log.Fatalf("error -- yaml Marshal: %v\n", err)}

	if outfil != nil && len(imgDat) > 0 {
		_, err = outfil.Write(imgDat)
		if err !=nil {log.Fatalf("error -- writing to yaml file! %v\n",err)}
	} else {
		log.Fatalf("error -- no outfil or yamldata!\n")
	}

}



func createImgList(md []byte) (imgList []ImgItem, err error) {

	var img ImgItem

	imgCount := 0
	for i:=0; i< len(md) -1; i++ {
		if md[i] != '!' {continue}
		if md[i+1] != '[' {continue}
		// parse image
		state := 0
		imgFilSt := 0
		imgParse := true
		for j:=i+1; j<len(md) -2; j++ {
			switch state {
			case 0:
				if md[j] == ']' {
					if md[j+1] != '(' {return nil, fmt.Errorf("error -- parsing no img file string!")}
					imgFilSt = j+2
					img.Capt = string(md[i+2:j])
					state = 1
				}
			case 1:
				if md[j] == ' ' {
					img.Filnam = string(md[imgFilSt:j])
					imgFilSt=j+1
					state = 2
				}
				if md[j] == ')' {
					img.Filnam = string(md[imgFilSt:j])
					i = j+1
					state = 0
					imgParse = false
				}

			case 2:
				if md[j] == ' ' {
					imgFilSt=j+1
				}
				if md[j] == ')' {
					img.Alt = string(md[imgFilSt:j])
					i = j+1
					state = 0
					imgParse = false
				}
			default:
			}
			if !imgParse {
				imgCount++
				break
			}
		}
		fmt.Printf("img file: %s %s %s\n", img.Capt, img.Alt, img.Filnam)
		imgList = append(imgList,img)
	}

	return imgList, nil
}


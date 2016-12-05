package main

/*
File: glyphpass.go
Authors: GlyphPass Programming Team (Holden Jones (Leader), Aaron Waner, Will
Jones)
Project: GlyphPass Graphical Password System
Description: Near-complete implementation using Sciter interface. Still needs Icon Set functionality and Export module
Submission Date: 4/18/2016
*/

import (
	"fmt"
	"github.com/hjones/glyphpass/lib"
	"github.com/oskca/sciter"
	"github.com/oskca/sciter/window"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"os"
	"strconv"
)

var BasicSet g2p.GlyphSet
var AbstractSet g2p.GlyphSet
var MyPassword g2p.GlyphPassword
var IconSet string = "basic"
var PassPhrase string = ""

func main() {
	// Create a Sciter window
	BasicSet = MakeBasicGlyphSet()
	AbstractSet = MakeAbstractGlyphSet()
	winsize := sciter.NewRect(10, 10, 850, 550)
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_MAIN, winsize)
	if err != nil {
		fmt.Println("Error. Could not open window:" , err)
	}
	w.LoadFile("interface.html")
	setEventHandler(w)
	w.Show()
	w.Run()
}

func setEventHandler(w *window.Window) {

	// Define a Sciter function for adding a Glyph to the list
	w.DefineFunction("addGlyph", func(args ...*sciter.Value) *sciter.Value {
		fmt.Println("ARGS " , args[0])

		// Stringify parameters sent back from HTML
		glyphClicked := args[0].String()

		XCoord , YCoord := glyphClicked[2:] , glyphClicked[0:1]

		XVal , _ := strconv.ParseInt(XCoord,0,64)
		YVal , _ := strconv.ParseInt(YCoord,0,64)

		// Subtract one
		XVal -= 1
		YVal -= 1

		if IconSet == "basic" {
			err := MyPassword.AddGlyph(BasicSet.GlyphArray[ XVal + (YVal * 6)])
			if err != nil {
				fmt.Println("Error. Could not add glyph, already full.")
			}
		}
		if IconSet == "abstract" {
			err := MyPassword.AddGlyph(AbstractSet.GlyphArray[ XVal + (YVal * 6)])
			if err != nil {
				fmt.Println("Error. Could not add glyph, already full.")
			}
		}


		var imgList [10]string
		fmt.Println( MyPassword.GlyphSlice )
		for i := 0 ; i < 10 ; i++ {
			if len(MyPassword.GlyphSlice) > i {
				imgList[i] = "<img class=\"histimg\" src=\"iconsets/" + IconSet + "/" + strconv.Itoa(MyPassword.GlyphSlice[i].YVal) + "x" + strconv.Itoa(MyPassword.GlyphSlice[i].XVal) + ".png\">"
				fmt.Println( imgList[i] )
			} else {
				imgList[i] = ""
			}
		}

		// Send data with callback function
		fn := args[1]
		fn.Invoke(sciter.NullValue(), "[Native Script]", sciter.NewValue(imgList[0]),sciter.NewValue(imgList[1]),sciter.NewValue(imgList[2]),sciter.NewValue(imgList[3]),sciter.NewValue(imgList[4]),sciter.NewValue(imgList[5]),sciter.NewValue(imgList[6]),sciter.NewValue(imgList[7]),sciter.NewValue(imgList[8]),sciter.NewValue(imgList[9]))
		ret := sciter.NewValue()
		ret.Set("ip", sciter.NewValue("127.0.0.1"))
		return ret
	})

	// Define a Sciter function for clearing the Glyph list
	w.DefineFunction("clearGlyphs", func(args ...*sciter.Value) *sciter.Value {

		MyPassword.GlyphSlice = nil

		var imgList [10]string
		for i := 0 ; i < 10 ; i++ {
			imgList[i] = "";
		}

		// Send data with callback function
		fn := args[0]
		fn.Invoke(sciter.NullValue(), "[Native Script]", sciter.NewValue(imgList[0]),sciter.NewValue(imgList[1]),sciter.NewValue(imgList[2]),sciter.NewValue(imgList[3]),sciter.NewValue(imgList[4]),sciter.NewValue(imgList[5]),sciter.NewValue(imgList[6]),sciter.NewValue(imgList[7]),sciter.NewValue(imgList[8]),sciter.NewValue(imgList[9]))
		ret := sciter.NewValue()
		ret.Set("ip", sciter.NewValue("127.0.0.1"))
		return ret
	})

	// Define a Sciter function for generating the password
	w.DefineFunction("generatePass", func(args ...*sciter.Value) *sciter.Value {
		var passPhrase string

		// Stringify parameters sent back from HTML
		passPhrase = args[0].String()

		fmt.Println( passPhrase , "<>" , MyPassword.GlyphSlice)

		var output, err string

		if(len(passPhrase) >= 1){
			// Make into GlyphPass
			output = MyPassword.GeneratePass( passPhrase )
		} else {
			err = "Must enter a passphrase"
		}

		// Send data with callback function
		fn := args[1]
		fn.Invoke(sciter.NullValue(), "[Native Script]", sciter.NewValue(output), sciter.NewValue(err))
		ret := sciter.NewValue()
		ret.Set("ip", sciter.NewValue("127.0.0.1"))
		return ret
	})

	w.DefineFunction("changeSet", func(args ...*sciter.Value) *sciter.Value {

		// Stringify parameters sent back from HTML
		IconSet = args[0].String()

		fmt.Println("Jesus")

		ioutil.WriteFile("iconset.txt",[]byte(IconSet),os.ModeExclusive)

		ret := sciter.NewValue()
		ret.Set("ip", sciter.NewValue("127.0.0.1"))
		return ret
	})

	w.DefineFunction("copyClip", func(args ...*sciter.Value) *sciter.Value {

		CopyClip := args[0].String()

		clipboard.WriteAll(CopyClip)

		ret := sciter.NewValue()
		ret.Set("ip", sciter.NewValue("127.0.0.1"))
		return ret
	})

}

// Function for creating the Basic glyphset
func MakeBasicGlyphSet() (g g2p.GlyphSet) {

	// First Row
	g1x1 := g2p.Glyph{1,1,"fish","fishfi01"}
	g1x2 := g2p.Glyph{1,2,"duck","duckdu02"}
	g1x3 := g2p.Glyph{1,3,"hedgehog","hedgeh03"}
	g1x4 := g2p.Glyph{1,4,"squirrel","squirr04"}
	g1x5 := g2p.Glyph{1,5,"dove","dovedo05"}
	g1x6 := g2p.Glyph{1,6,"sheep","sheeps06"}
	// Second Row
	g2x1 := g2p.Glyph{2,1,"snail","snails07"}
	g2x2 := g2p.Glyph{2,2,"whale","whalew08"}
	g2x3 := g2p.Glyph{2,3,"ladybug","ladybu09"}
	g2x4 := g2p.Glyph{2,4,"rabbit","rabbit10"}
	g2x5 := g2p.Glyph{2,5,"bear","bearbe11"}
	g2x6 := g2p.Glyph{2,6,"chicken","chicke12"}
	// Third Row
	g3x1 := g2p.Glyph{3,1,"turtle","turtle13"}
	g3x2 := g2p.Glyph{3,2,"horse","horseh14"}
	g3x3 := g2p.Glyph{3,3,"bat","batbat15"}
	g3x4 := g2p.Glyph{3,4,"pig","pigpig16"}
	g3x5 := g2p.Glyph{3,5,"frog","frogfr17"}
	g3x6 := g2p.Glyph{3,6,"lion","lionli18"}
	// Fourth Row
	g4x1 := g2p.Glyph{4,1,"giraffe","giraff19"}
	g4x2 := g2p.Glyph{4,2,"crab","crabcr20"}
	g4x3 := g2p.Glyph{4,3,"cat","catcat21"}
	g4x4 := g2p.Glyph{4,4,"octopus","octopu22"}
	g4x5 := g2p.Glyph{4,5,"paw","pawpaw23"}
	g4x6 := g2p.Glyph{4,6,"bone","bonebo24"}
	// Fifth Row
	g5x1 := g2p.Glyph{5,1,"mouse","mousem25"}
	g5x2 := g2p.Glyph{5,2,"shark","sharks26"}
	g5x3 := g2p.Glyph{5,3,"shell","shells27"}
	g5x4 := g2p.Glyph{5,4,"lobster","lobste28"}
	g5x5 := g2p.Glyph{5,5,"cow","cowcow29"}
	g5x6 := g2p.Glyph{5,6,"elephant","elepha30"}
	// Sixth Row
	g6x1 := g2p.Glyph{6,1,"butterfly","butter31"}
	g6x2 := g2p.Glyph{6,2,"monkey","monkey32"}
	g6x3 := g2p.Glyph{6,3,"snake","snakes33"}
	g6x4 := g2p.Glyph{6,4,"scorpion","scorpi34"}
	g6x5 := g2p.Glyph{6,5,"mantaray","mantar35"}
	g6x6 := g2p.Glyph{6,6,"dolphin","dolphi36"}

	// Structure in the format of the GlyphSet struct
	return g2p.GlyphSet{[...]g2p.Glyph{
								g1x1 , g1x2 , g1x3 , g1x4 , g1x5 , g1x6 ,
								g2x1 , g2x2 , g2x3 , g2x4 , g2x5 , g2x6 ,
								g3x1 , g3x2 , g3x3 , g3x4 , g3x5 , g3x6 ,
								g4x1 , g4x2 , g4x3 , g4x4 , g4x5 , g4x6 ,
								g5x1 , g5x2 , g5x3 , g5x4 , g5x5 , g5x6 ,
								g6x1 , g6x2 , g6x3 , g6x4 , g6x5 , g6x6} , "Basic Set" , "basic"}
}

func MakeAbstractGlyphSet() (g g2p.GlyphSet) {

	// First Row
	g1x1 := g2p.Glyph{1,1,"reply","replyr01"}
	g1x2 := g2p.Glyph{1,2,"key","keykey02"}
	g1x3 := g2p.Glyph{1,3,"tag","tagtag03"}
	g1x4 := g2p.Glyph{1,4,"phone","phonep04"}
	g1x5 := g2p.Glyph{1,5,"check","checkc05"}
	g1x6 := g2p.Glyph{1,6,"radius","radius06"}
	// Second Row
	g2x1 := g2p.Glyph{2,1,"rss","rssrss07"}
	g2x2 := g2p.Glyph{2,2,"cross","crossc08"}
	g2x3 := g2p.Glyph{2,3,"trash","trasht09"}
	g2x4 := g2p.Glyph{2,4,"mouse","mousem10"}
	g2x5 := g2p.Glyph{2,5,"heart","hearth11"}
	g2x6 := g2p.Glyph{2,6,"star","starst12"}
	// Third Row
	g3x1 := g2p.Glyph{3,1,"beaker","beaker13"}
	g3x2 := g2p.Glyph{3,2,"controls","contro14"}
	g3x3 := g2p.Glyph{3,3,"share","shares15"}
	g3x4 := g2p.Glyph{3,4,"wand","wandwa16"}
	g3x5 := g2p.Glyph{3,5,"chat","chatch17"}
	g3x6 := g2p.Glyph{3,6,"eye","eyeeye18"}
	// Fourth Row
	g4x1 := g2p.Glyph{4,1,"fire","firefi19"}
	g4x2 := g2p.Glyph{4,2,"marker","marker20"}
	g4x3 := g2p.Glyph{4,3,"search","search21"}
	g4x4 := g2p.Glyph{4,4,"mail","mailma22"}
	g4x5 := g2p.Glyph{4,5,"menu","menume23"}
	g4x6 := g2p.Glyph{4,6,"pencil","pencil24"}
	// Fifth Row
	g5x1 := g2p.Glyph{5,1,"up","upupup25"}
	g5x2 := g2p.Glyph{5,2,"play","playpl26"}
	g5x3 := g2p.Glyph{5,3,"power","powerp27"}
	g5x4 := g2p.Glyph{5,4,"gear","gearge28"}
	g5x5 := g2p.Glyph{5,5,"user","userus29"}
	g5x6 := g2p.Glyph{5,6,"shop","shopsh30"}
	// Sixth Row
	g6x1 := g2p.Glyph{6,1,"bell","bellbe31"}
	g6x2 := g2p.Glyph{6,2,"sound","sounds32"}
	g6x3 := g2p.Glyph{6,3,"filter","filter33"}
	g6x4 := g2p.Glyph{6,4,"tools","toolst34"}
	g6x5 := g2p.Glyph{6,5,"printer","printe35"}
	g6x6 := g2p.Glyph{6,6,"video","videov36"}

	// Structure in the format of the GlyphSet struct
	return g2p.GlyphSet{[...]g2p.Glyph{
								g1x1 , g1x2 , g1x3 , g1x4 , g1x5 , g1x6 ,
								g2x1 , g2x2 , g2x3 , g2x4 , g2x5 , g2x6 ,
								g3x1 , g3x2 , g3x3 , g3x4 , g3x5 , g3x6 ,
								g4x1 , g4x2 , g4x3 , g4x4 , g4x5 , g4x6 ,
								g5x1 , g5x2 , g5x3 , g5x4 , g5x5 , g5x6 ,
								g6x1 , g6x2 , g6x3 , g6x4 , g6x5 , g6x6} , "Abstract Set" , "abstract"}
}

package g2p

/*
File: g2p.go
Authors: GlyphPass Programming Team (Holden Jones (Leader), Aaron Waner, Will
Jones)
Project: GlyphPass Graphical Password System
Description: glyph-2-password module for GlyphPass
Submission Date: 4/18/2016
*/

import (
  "bytes"
  "unicode/utf8"
  "errors"
)

// TYPE DEFINITIONS
type Glyph struct {
  YVal int
  XVal int
  GlyphName string
  DataWord string
}

type GlyphSet struct {
  // GlyphSet is an array of 36 glyphs tied to a name and folder
  GlyphArray [36]Glyph
  GlyphSetName string
  GlyphSetFolder string
}

type GlyphPassword struct {
  GlyphSlice []Glyph // Slice with max length 10
}

// TYPE METHODS

func (g GlyphPassword) getConverted() (s string) {
  appendString := new(bytes.Buffer)

  for _ , p := range g.GlyphSlice {
    appendString.WriteString(p.DataWord)
  }

  return appendString.String()
}

func (g *GlyphPassword) Init() {
  g.GlyphSlice = make([]Glyph,10)
}

func (g *GlyphPassword) AddGlyph(gl Glyph) (err error) {
  if len(g.GlyphSlice) < 10 {
    g.GlyphSlice = append(g.GlyphSlice , gl)
    return nil
  }
  return errors.New("No more room for Glyph")
}

func (g *GlyphPassword) ConvertedPhrase() (ret string) {
  converted := new(bytes.Buffer)
  for _ , glyf := range g.GlyphSlice {
    converted.WriteString(glyf.DataWord)
  }
  return converted.String()
}

func (g *GlyphPassword) GeneratePass( passPhrase string ) (ret string) {
  converted := new(bytes.Buffer)
  for _ , glyf := range g.GlyphSlice {
    converted.WriteString(glyf.DataWord)
  }

  if utf8.RuneCountInString(passPhrase) > 10 {
    passPhrase = passPhrase[:10]
  }

  convertInts := make([]int,255)
  addInt := 0

  for i , glyf := range converted.String() {
    addInt += (int(glyf) % utf8.RuneCountInString(passPhrase))
    convertInts[i] = (int(glyf) * addInt)
  }

  gHsh := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()-"
  var symHash = []rune(gHsh)

  finalString := new(bytes.Buffer)

  for i , _ := range converted.String() {
    // Debug line
    // fmt.Println("CONVERT: " , i , "||" , converted.String()[i] , "||" , addInt , "||" , convertInts[i])
    finalString.WriteRune(symHash[convertInts[i] % 73])
  }

  return finalString.String()
}

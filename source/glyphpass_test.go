package main

/*
File: glyphpass_test.go
Authors: GlyphPass Programming Team (Holden Jones (Leader), Aaron Waner, Will
Jones)
Project: GlyphPass Graphical Password System
Description: Golang Test unit testing functionality
Submission Date: 4/18/2016
*/

import (
	"github.com/hjones/glyphpass/lib"
  "testing"
)

// To test, run "go test -v" in directory with Go compiler installed

// Test Glyph basic functionality
func Test_Glyph(t *testing.T) {
  e := g2p.Glyph{1,1,"testglyph","testtest02"}
  if e.DataWord != "testtest02" {
    t.Error("Expected testtest02, got ", e.DataWord)
  }
}

// Test GlyphPassword basic functionality
func Test_GlyphPassword_Basic(t *testing.T) {
  var e g2p.GlyphPassword
  e.AddGlyph(g2p.Glyph{1,1,"testglyph","testtest01"})
  if e.GlyphSlice[0].DataWord != "testtest01" {
    t.Error("Expected testtest01, got ", e.GlyphSlice[0].DataWord)
  }
}

// Test GlyphPassword basic functionality
func Test_GlyphPassword_Addition(t *testing.T) {
  var e g2p.GlyphPassword
  e.AddGlyph(g2p.Glyph{1,1,"testglyph","barbar01"})
  e.AddGlyph(g2p.Glyph{1,2,"testglyph2","foofoo02"})

  if e.ConvertedPhrase() != "barbar01foofoo02" {
    t.Error("Expected barbar01foofoo02, got ", e.ConvertedPhrase())
  }
}

// Test GlyphPassword basic generation
func Test_GlyphPassword_Generation(t *testing.T) {
  var e g2p.GlyphPassword
  e.AddGlyph(g2p.Glyph{1,1,"fish","fishfi01"})

  if e.GeneratePass("testtest") != "Cf4sAHng" {
    t.Error("Expected Cf4sAHng, got ", e.GeneratePass("testtest"))
  }
}

// Test GlyphPassword complex generation
func Test_GlyphPassword_GenerateComplex(t *testing.T) {
  var e g2p.GlyphPassword
  e.AddGlyph(g2p.Glyph{1,1,"fish","fishfi01"})
  e.AddGlyph(g2p.Glyph{2,1,"snail","snails07"})
  e.AddGlyph(g2p.Glyph{3,1,"turtle","turtle13"})
  e.AddGlyph(g2p.Glyph{4,1,"giraffe","giraff19"})
  e.AddGlyph(g2p.Glyph{5,1,"mouse","mousem25"})
  e.AddGlyph(g2p.Glyph{6,1,"butterfly","butter31"})

  if e.GeneratePass("testtest") != "Cf4sAHnggYp191XDYzym2ZVSeUK((z5(8qi3EkXPrC9nYeC2" {
    t.Error("Expected Cf4sAHnggYp191XDYzym2ZVSeUK((z5(8qi3EkXPrC9nYeC2, got ", e.GeneratePass("testtest"))
  }
}

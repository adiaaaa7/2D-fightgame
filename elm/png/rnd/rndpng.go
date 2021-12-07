package rnd

import (
	"bytes"
	"game/myds/elm/png/mp"
	"game/myds/elm/png/pc"

	"image"
	"log"
)

func MpRand(sel int) image.Image {
	var img image.Image
	var err error
	switch sel {
	case 1:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp1_png))
	case 2:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp2_png))
	case 3:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp3_png))
	case 4:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp4_png))
	case 5:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp5_png))
	case 6:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp6_png))
	case 7:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp7_png))
	case 8:
		img, _, err = image.Decode(bytes.NewReader(mp.Mp8_png))
	}
	errLog(err)
	return img
}

func PcOpRand(sel int) image.Image {
	var img image.Image
	var err error
	switch sel {
	case 1:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt1_png))
	case 2:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt2_png))
	case 3:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt3_png))
	case 4:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt4_png))
	case 5:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt5_png))
	case 6:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt6_png))
	case 7:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt7_png))
	case 8:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt8_png))
	case 9:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt9_png))
	case 10:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt10_png))
	case 11:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt11_png))
	case 12:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt12_png))
	case 13:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt13_png))
	case 14:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt14_png))
	case 15:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt15_png))
	case 16:
		img, _, err = image.Decode(bytes.NewReader(pc.Pt16_png))
	}
	errLog(err)
	return img
}

func PcStuRand(sel int) image.Image {
	var img image.Image
	var err error
	switch sel {
	case 1:
		img, _, err = image.Decode(bytes.NewReader(pc.Su1_png))
	case 2:
		img, _, err = image.Decode(bytes.NewReader(pc.Su2_png))
	case 3:
		img, _, err = image.Decode(bytes.NewReader(pc.Su3_png))
	case 4:
		img, _, err = image.Decode(bytes.NewReader(pc.Su4_png))
	case 5:
		img, _, err = image.Decode(bytes.NewReader(pc.Su5_png))
	case 6:
		img, _, err = image.Decode(bytes.NewReader(pc.Su6_png))
	case 7:
		img, _, err = image.Decode(bytes.NewReader(pc.Su7_png))
	case 8:
		img, _, err = image.Decode(bytes.NewReader(pc.Su8_png))
	case 9:
		img, _, err = image.Decode(bytes.NewReader(pc.Su9_png))
	case 10:
		img, _, err = image.Decode(bytes.NewReader(pc.Su10_png))
	case 11:
		img, _, err = image.Decode(bytes.NewReader(pc.Su11_png))
	case 12:
		img, _, err = image.Decode(bytes.NewReader(pc.Su12_png))
	case 13:
		img, _, err = image.Decode(bytes.NewReader(pc.Su13_png))
	case 14:
		img, _, err = image.Decode(bytes.NewReader(pc.Su14_png))
	case 15:
		img, _, err = image.Decode(bytes.NewReader(pc.Su15_png))
	case 16:
		img, _, err = image.Decode(bytes.NewReader(pc.Su16_png))
	case 17:
		img, _, err = image.Decode(bytes.NewReader(pc.Su17_png))
	case 18:
		img, _, err = image.Decode(bytes.NewReader(pc.Su18_png))
	case 19:
		img, _, err = image.Decode(bytes.NewReader(pc.Su19_png))
	case 20:
		img, _, err = image.Decode(bytes.NewReader(pc.Su20_png))
	case 21:
		img, _, err = image.Decode(bytes.NewReader(pc.Su21_png))
	case 22:
		img, _, err = image.Decode(bytes.NewReader(pc.Su22_png))
	case 23:
		img, _, err = image.Decode(bytes.NewReader(pc.Su23_png))
	case 24:
		img, _, err = image.Decode(bytes.NewReader(pc.Su24_png))
	case 25:
		img, _, err = image.Decode(bytes.NewReader(pc.Su25_png))
	case 26:
		img, _, err = image.Decode(bytes.NewReader(pc.Su26_png))
	case 27:
		img, _, err = image.Decode(bytes.NewReader(pc.Su27_png))
	case 28:
		img, _, err = image.Decode(bytes.NewReader(pc.Su28_png))
	case 29:
		img, _, err = image.Decode(bytes.NewReader(pc.Su29_png))
	case 30:
		img, _, err = image.Decode(bytes.NewReader(pc.Su30_png))
	case 31:
		img, _, err = image.Decode(bytes.NewReader(pc.Su31_png))
	case 32:
		img, _, err = image.Decode(bytes.NewReader(pc.Su32_png))
	}
	errLog(err)
	return img
}

func errLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package gomaterial

import "image/color"

func Rgb(c uint32) color.NRGBA {
	return argb((0xff << 24) | c)
}

func argb(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

var (
	Red50  = Rgb(0xffebee)
	Red100 = Rgb(0xffcdd2)
	Red200 = Rgb(0xef9a9a)
	Red300 = Rgb(0xe57373)
	Red400 = Rgb(0xef5350)
	Red500 = Rgb(0xf44336)
	Red600 = Rgb(0xe53935)
	Red700 = Rgb(0xd32f2f)
	Red800 = Rgb(0xc62828)
	Red900 = Rgb(0xb71c1c)

	Pink50  = Rgb(0xfce4ec)
	Pink100 = Rgb(0xf8bbd0)
	Pink200 = Rgb(0xf48fb1)
	Pink300 = Rgb(0xf06292)
	Pink400 = Rgb(0xec407a)
	Pink500 = Rgb(0xe91e63)
	Pink600 = Rgb(0xd81b60)
	Pink700 = Rgb(0xc2185b)
	Pink800 = Rgb(0xad1457)
	Pink900 = Rgb(0x880e4f)

	Purple50  = Rgb(0xf3e5f5)
	Purple100 = Rgb(0xe1bee7)
	Purple200 = Rgb(0xce93d8)
	Purple300 = Rgb(0xba68c8)
	Purple400 = Rgb(0xab47bc)
	Purple500 = Rgb(0x9c27b0)
	Purple600 = Rgb(0x8e24aa)
	Purple700 = Rgb(0x7b1fa2)
	Purple800 = Rgb(0x6a1b9a)
	Purple900 = Rgb(0x4a148c)

	DeepPurple50  = Rgb(0xede7f6)
	DeepPurple100 = Rgb(0xd1c4e9)
	DeepPurple200 = Rgb(0xb39ddb)
	DeepPurple300 = Rgb(0x9575cd)
	DeepPurple400 = Rgb(0x7e57c2)
	DeepPurple500 = Rgb(0x673ab7)
	DeepPurple600 = Rgb(0x5e35b1)
	DeepPurple700 = Rgb(0x512da8)
	DeepPurple800 = Rgb(0x4527a0)
	DeepPurple900 = Rgb(0x311b92)

	Indigo50  = Rgb(0xe8eaf6)
	Indigo100 = Rgb(0xc5cae9)
	Indigo200 = Rgb(0x9fa8da)
	Indigo300 = Rgb(0x7986cb)
	Indigo400 = Rgb(0x5c6bc0)
	Indigo500 = Rgb(0x3f51b5)
	Indigo600 = Rgb(0x3949ab)
	Indigo700 = Rgb(0x303f9f)
	Indigo800 = Rgb(0x283593)
	Indigo900 = Rgb(0x1a237e)

	Blue50  = Rgb(0xe3f2fd)
	Blue100 = Rgb(0xbbdefb)
	Blue200 = Rgb(0x90caf9)
	Blue300 = Rgb(0x64b5f6)
	Blue400 = Rgb(0x42a5f5)
	Blue500 = Rgb(0x2196f3)
	Blue600 = Rgb(0x1e88e5)
	Blue700 = Rgb(0x1976d2)
	Blue800 = Rgb(0x1565c0)
	Blue900 = Rgb(0x0d47a1)

	LightBlue50  = Rgb(0xe1f5fe)
	LightBlue100 = Rgb(0xb3e5fc)
	LightBlue200 = Rgb(0x81d4fa)
	LightBlue300 = Rgb(0x4fc3f7)
	LightBlue400 = Rgb(0x29b6f6)
	LightBlue500 = Rgb(0x03a9f4)
	LightBlue600 = Rgb(0x039be5)
	LightBlue700 = Rgb(0x0288d1)
	LightBlue800 = Rgb(0x0277bd)
	LightBlue900 = Rgb(0x01579b)

	Cyan50  = Rgb(0xe0f7fa)
	Cyan100 = Rgb(0xb2ebf2)
	Cyan200 = Rgb(0x80deea)
	Cyan300 = Rgb(0x4dd0e1)
	Cyan400 = Rgb(0x26c6da)
	Cyan500 = Rgb(0x00bcd4)
	Cyan600 = Rgb(0x00acc1)
	Cyan700 = Rgb(0x0097a7)
	Cyan800 = Rgb(0x00838f)
	Cyan900 = Rgb(0x006064)

	Teal50  = Rgb(0xe0f2f1)
	Teal100 = Rgb(0xb2dfdb)
	Teal200 = Rgb(0x80cbc4)
	Teal300 = Rgb(0x4db6ac)
	Teal400 = Rgb(0x26a69a)
	Teal500 = Rgb(0x009688)
	Teal600 = Rgb(0x00897b)
	Teal700 = Rgb(0x00796b)
	Teal800 = Rgb(0x00695c)
	Teal900 = Rgb(0x004d40)

	Green50  = Rgb(0xe8f5e9)
	Green100 = Rgb(0xc8e6c9)
	Green200 = Rgb(0xa5d6a7)
	Green300 = Rgb(0x81c784)
	Green400 = Rgb(0x66bb6a)
	Green500 = Rgb(0x4caf50)
	Green600 = Rgb(0x43a047)
	Green700 = Rgb(0x388e3c)
	Green800 = Rgb(0x2e7d32)
	Green900 = Rgb(0x1b5e2)

	LightGreen50  = Rgb(0xf1f8e9)
	LightGreen100 = Rgb(0xdcedc8)
	LightGreen200 = Rgb(0xc5e1a5)
	LightGreen300 = Rgb(0xaed581)
	LightGreen400 = Rgb(0x9ccc65)
	LightGreen500 = Rgb(0x8bc34a)
	LightGreen600 = Rgb(0x7cb342)
	LightGreen700 = Rgb(0x689f38)
	LightGreen800 = Rgb(0x558b2f)
	LightGreen900 = Rgb(0x33691e)

	Lime50  = Rgb(0xf9fbe7)
	Lime100 = Rgb(0xf0f4c3)
	Lime200 = Rgb(0xe6ee9c)
	Lime300 = Rgb(0xdce775)
	Lime400 = Rgb(0xd4e157)
	Lime500 = Rgb(0xcddc39)
	Lime600 = Rgb(0xc0ca33)
	Lime700 = Rgb(0xafb42b)
	Lime800 = Rgb(0x9e9d24)
	Lime900 = Rgb(0x827717)

	Yellow50  = Rgb(0xfffde7)
	Yellow100 = Rgb(0xfff9c4)
	Yellow200 = Rgb(0xfff59d)
	Yellow300 = Rgb(0xfff176)
	Yellow400 = Rgb(0xffee58)
	Yellow500 = Rgb(0xffeb3b)
	Yellow600 = Rgb(0xfdd835)
	Yellow700 = Rgb(0xfbc02d)
	Yellow800 = Rgb(0xf9a825)
	Yellow900 = Rgb(0xf57f17)

	Amber50  = Rgb(0xfff8e1)
	Amber100 = Rgb(0xffecb3)
	Amber200 = Rgb(0xffe082)
	Amber300 = Rgb(0xffd54f)
	Amber400 = Rgb(0xffca28)
	Amber500 = Rgb(0xffc107)
	Amber600 = Rgb(0xffb300)
	Amber700 = Rgb(0xffa000)
	Amber800 = Rgb(0xff8f00)
	Amber900 = Rgb(0xff6f00)

	Orange50  = Rgb(0xfff3e0)
	Orange100 = Rgb(0xffe0b2)
	Orange200 = Rgb(0xffcc80)
	Orange300 = Rgb(0xffb74d)
	Orange400 = Rgb(0xffa726)
	Orange500 = Rgb(0xff9800)
	Orange600 = Rgb(0xfb8c00)
	Orange700 = Rgb(0xf57c00)
	Orange800 = Rgb(0xef6c00)
	Orange900 = Rgb(0xe65100)

	DeepOrange50  = Rgb(0xfbe9e7)
	DeepOrange100 = Rgb(0xffccbc)
	DeepOrange200 = Rgb(0xffab91)
	DeepOrange300 = Rgb(0xff8a65)
	DeepOrange400 = Rgb(0xff7043)
	DeepOrange500 = Rgb(0xff5722)
	DeepOrange600 = Rgb(0xf4511e)
	DeepOrange700 = Rgb(0xe64a19)
	DeepOrange800 = Rgb(0xd84315)
	DeepOrange900 = Rgb(0xbf360c)

	Brown50  = Rgb(0xefebe9)
	Brown100 = Rgb(0xd7ccc8)
	Brown200 = Rgb(0xbcaaa4)
	Brown300 = Rgb(0xa1887f)
	Brown400 = Rgb(0x8d6e63)
	Brown500 = Rgb(0x795548)
	Brown600 = Rgb(0x6d4c41)
	Brown700 = Rgb(0x5d4037)
	Brown800 = Rgb(0x4e342e)
	Brown900 = Rgb(0x3e2723)

	Grey50  = Rgb(0xfafafa)
	Grey100 = Rgb(0xf5f5f5)
	Grey200 = Rgb(0xeeeeee)
	Grey300 = Rgb(0xe0e0e0)
	Grey400 = Rgb(0xbdbdbd)
	Grey500 = Rgb(0x9e9e9e)
	Grey600 = Rgb(0x757575)
	Grey700 = Rgb(0x616161)
	Grey800 = Rgb(0x424242)
	Grey900 = Rgb(0x212121)

	BlueGrey50  = Rgb(0xeceff1)
	BlueGrey100 = Rgb(0xcfd8dc)
	BlueGrey200 = Rgb(0xb0bec5)
	BlueGrey300 = Rgb(0x90a4ae)
	BlueGrey400 = Rgb(0x78909c)
	BlueGrey500 = Rgb(0x607d8b)
	BlueGrey600 = Rgb(0x546e7a)
	BlueGrey700 = Rgb(0x455a64)
	BlueGrey800 = Rgb(0x37474f)
	BlueGrey900 = Rgb(0x263238)
)

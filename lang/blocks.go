package lang

var (
	blockRshift uint     = 4
	blocks      []string = make([]string, 0x2fa2)
)

func init() {
	copy(blocks[0x0:0x8], getRepeated("Basic Latin", 0x8))
	copy(blocks[0x0:0x8], getRepeated("Basic Latin", 0x8))
	copy(blocks[0x8:0x10], getRepeated("Extended Latin", 0x8))  // Latin-1 Supplement
	copy(blocks[0x10:0x18], getRepeated("Extended Latin", 0x8)) // Latin Extended-A
	copy(blocks[0x18:0x25], getRepeated("Latin Extended-B", 0xd))
	copy(blocks[0x25:0x2b], getRepeated("Extended Latin", 0x6)) // IPA Extensions
	copy(blocks[0x2b:0x30], getRepeated("Spacing Modifier Letters", 0x5))
	copy(blocks[0x37:0x40], getRepeated("Greek and Coptic", 0x9))
	copy(blocks[0x40:0x50], getRepeated("Cyrillic", 0x10))
	copy(blocks[0x50:0x53], getRepeated("Cyrillic Supplement", 0x3))
	copy(blocks[0x53:0x59], getRepeated("Armenian", 0x6))
	copy(blocks[0x59:0x60], getRepeated("Hebrew", 0x7))
	copy(blocks[0x60:0x70], getRepeated("Arabic", 0x10))
	copy(blocks[0x70:0x75], getRepeated("Syriac", 0x5))
	copy(blocks[0x75:0x78], getRepeated("Arabic Supplement", 0x3))
	copy(blocks[0x78:0x7c], getRepeated("Thaana", 0x4))
	copy(blocks[0x7c:0x80], getRepeated("NKo", 0x4))
	copy(blocks[0x80:0x84], getRepeated("Samaritan", 0x4))
	copy(blocks[0x84:0x86], getRepeated("Mandaic", 0x2))
	copy(blocks[0x8a:0x90], getRepeated("Arabic Extended-A", 0x6))
	copy(blocks[0x90:0x98], getRepeated("Devanagari", 0x8))
	copy(blocks[0x98:0xa0], getRepeated("Bengali", 0x8))
	copy(blocks[0xa0:0xa8], getRepeated("Gurmukhi", 0x8))
	copy(blocks[0xa8:0xb0], getRepeated("Gujarati", 0x8))
	copy(blocks[0xb0:0xb8], getRepeated("Oriya", 0x8))
	copy(blocks[0xb8:0xc0], getRepeated("Tamil", 0x8))
	copy(blocks[0xc0:0xc8], getRepeated("Telugu", 0x8))
	copy(blocks[0xc8:0xd0], getRepeated("Kannada", 0x8))
	copy(blocks[0xd0:0xd8], getRepeated("Malayalam", 0x8))
	copy(blocks[0xd8:0xe0], getRepeated("Sinhala", 0x8))
	copy(blocks[0xe0:0xe8], getRepeated("Thai", 0x8))
	copy(blocks[0xe8:0xf0], getRepeated("Lao", 0x8))
	copy(blocks[0xf0:0x100], getRepeated("Tibetan", 0x10))
	copy(blocks[0x100:0x10a], getRepeated("Myanmar", 0xa))
	copy(blocks[0x10a:0x110], getRepeated("Georgian", 0x6))
	copy(blocks[0x110:0x120], getRepeated("Hangul Jamo", 0x10))
	copy(blocks[0x120:0x138], getRepeated("Ethiopic", 0x18))
	copy(blocks[0x138:0x13a], getRepeated("Ethiopic Supplement", 0x2))
	copy(blocks[0x13a:0x140], getRepeated("Cherokee", 0x6))
	copy(blocks[0x140:0x168], getRepeated("Unified Canadian Aboriginal Syllabics", 0x28))
	copy(blocks[0x168:0x16a], getRepeated("Ogham", 0x2))
	copy(blocks[0x16a:0x170], getRepeated("Runic", 0x6))
	copy(blocks[0x170:0x172], getRepeated("Tagalog", 0x2))
	copy(blocks[0x172:0x174], getRepeated("Hanunoo", 0x2))
	copy(blocks[0x174:0x176], getRepeated("Buhid", 0x2))
	copy(blocks[0x176:0x178], getRepeated("Tagbanwa", 0x2))
	copy(blocks[0x178:0x180], getRepeated("Khmer", 0x8))
	copy(blocks[0x180:0x18b], getRepeated("Mongolian", 0xb))
	copy(blocks[0x18b:0x190], getRepeated("Unified Canadian Aboriginal Syllabics Extended", 0x5))
	copy(blocks[0x190:0x195], getRepeated("Limbu", 0x5))
	copy(blocks[0x195:0x198], getRepeated("Tai Le", 0x3))
	copy(blocks[0x198:0x19e], getRepeated("New Tai Lue", 0x6))
	copy(blocks[0x1a0:0x1a2], getRepeated("Buginese", 0x2))
	copy(blocks[0x1a2:0x1ab], getRepeated("Tai Tham", 0x9))
	copy(blocks[0x1b0:0x1b8], getRepeated("Balinese", 0x8))
	copy(blocks[0x1b8:0x1bc], getRepeated("Sundanese", 0x4))
	copy(blocks[0x1bc:0x1c0], getRepeated("Batak", 0x4))
	copy(blocks[0x1c0:0x1c5], getRepeated("Lepcha", 0x5))
	copy(blocks[0x1c5:0x1c8], getRepeated("Ol Chiki", 0x3))
	copy(blocks[0x1cd:0x1d0], getRepeated("Vedic Extensions", 0x3))
	copy(blocks[0x1d0:0x1d8], getRepeated("Phonetic Extensions", 0x8))
	copy(blocks[0x1d8:0x1dc], getRepeated("Phonetic Extensions Supplement", 0x4))
	copy(blocks[0x1e0:0x1f0], getRepeated("Latin Extended Additional", 0x10))
	copy(blocks[0x1f0:0x200], getRepeated("Greek Extended", 0x10))
	copy(blocks[0x207:0x20a], getRepeated("Superscripts and Subscripts", 0x3))
	copy(blocks[0x210:0x215], getRepeated("Letterlike Symbols", 0x5))
	copy(blocks[0x215:0x219], getRepeated("Number Forms", 0x4))
	copy(blocks[0x2c0:0x2c6], getRepeated("Glagolitic", 0x6))
	copy(blocks[0x2c6:0x2c8], getRepeated("Latin Extended-C", 0x2))
	copy(blocks[0x2c8:0x2d0], getRepeated("Coptic", 0x8))
	copy(blocks[0x2d0:0x2d3], getRepeated("Georgian Supplement", 0x3))
	copy(blocks[0x2d3:0x2d8], getRepeated("Tifinagh", 0x5))
	copy(blocks[0x2d8:0x2de], getRepeated("Ethiopic Extended", 0x6))
	copy(blocks[0x2e0:0x2e8], getRepeated("Supplemental Punctuation", 0x8))
	copy(blocks[0x300:0x304], getRepeated("CJK Symbols and Punctuation", 0x4))
	copy(blocks[0x304:0x30a], getRepeated("Kana", 0x6)) // Hiragana
	copy(blocks[0x30a:0x310], getRepeated("Kana", 0x6)) // Katakana
	copy(blocks[0x310:0x313], getRepeated("Bopomofo", 0x3))
	copy(blocks[0x313:0x319], getRepeated("Hangul Compatibility Jamo", 0x6))
	copy(blocks[0x31a:0x31c], getRepeated("Bopomofo Extended", 0x2))
	copy(blocks[0x31f:0x320], getRepeated("Kana", 0x1)) // Katakana Phonetic Extensions
	copy(blocks[0x340:0x4dc], getRepeated("CJK Unified Ideographs Extension A", 0x19c))
	copy(blocks[0x4e0:0xa00], getRepeated("CJK Unified Ideographs", 0x520))
	copy(blocks[0xa00:0xa49], getRepeated("Yi Syllables", 0x49))
	copy(blocks[0xa4d:0xa50], getRepeated("Lisu", 0x3))
	copy(blocks[0xa50:0xa64], getRepeated("Vai", 0x14))
	copy(blocks[0xa64:0xa6a], getRepeated("Cyrillic Extended-B", 0x6))
	copy(blocks[0xa6a:0xa70], getRepeated("Bamum", 0x6))
	copy(blocks[0xa70:0xa72], getRepeated("Modifier Tone Letters", 0x2))
	copy(blocks[0xa72:0xa80], getRepeated("Latin Extended-D", 0xe))
	copy(blocks[0xa80:0xa83], getRepeated("Syloti Nagri", 0x3))
	copy(blocks[0xa84:0xa88], getRepeated("Phags-pa", 0x4))
	copy(blocks[0xa88:0xa8e], getRepeated("Saurashtra", 0x6))
	copy(blocks[0xa8e:0xa90], getRepeated("Devanagari Extended", 0x2))
	copy(blocks[0xa90:0xa93], getRepeated("Kayah Li", 0x3))
	copy(blocks[0xa93:0xa96], getRepeated("Rejang", 0x3))
	copy(blocks[0xa96:0xa98], getRepeated("Hangul Jamo Extended-A", 0x2))
	copy(blocks[0xa98:0xa9e], getRepeated("Javanese", 0x6))
	copy(blocks[0xaa0:0xaa6], getRepeated("Cham", 0x6))
	copy(blocks[0xaa6:0xaa8], getRepeated("Myanmar Extended-A", 0x2))
	copy(blocks[0xaa8:0xaae], getRepeated("Tai Viet", 0x6))
	copy(blocks[0xaae:0xab0], getRepeated("Meetei Mayek Extensions", 0x2))
	copy(blocks[0xab0:0xab3], getRepeated("Ethiopic Extended-A", 0x3))
	copy(blocks[0xabc:0xac0], getRepeated("Meetei Mayek", 0x4))
	copy(blocks[0xac0:0xd7b], getRepeated("Hangul Syllables", 0x2bb))
	copy(blocks[0xd7b:0xd80], getRepeated("Hangul Jamo Extended-B", 0x5))
	copy(blocks[0xf90:0xfb0], getRepeated("CJK Compatibility Ideographs", 0x20))
	copy(blocks[0xfb0:0xfb5], getRepeated("Alphabetic Presentation Forms", 0x5))
	copy(blocks[0xfb5:0xfe0], getRepeated("Arabic Presentation Forms-A", 0x2b))
	copy(blocks[0xfe7:0xff0], getRepeated("Arabic Presentation Forms-B", 0x9))
	copy(blocks[0xff0:0xfff], getRepeated("Halfwidth and Fullwidth Forms", 0xf))
	copy(blocks[0x1000:0x1008], getRepeated("Linear B Syllabary", 0x8))
	copy(blocks[0x1008:0x1010], getRepeated("Linear B Ideograms", 0x8))
	copy(blocks[0x1028:0x102a], getRepeated("Lycian", 0x2))
	copy(blocks[0x102a:0x102e], getRepeated("Carian", 0x4))
	copy(blocks[0x1030:0x1033], getRepeated("Old Italic", 0x3))
	copy(blocks[0x1033:0x1035], getRepeated("Gothic", 0x2))
	copy(blocks[0x1038:0x103a], getRepeated("Ugaritic", 0x2))
	copy(blocks[0x103a:0x103e], getRepeated("Old Persian", 0x4))
	copy(blocks[0x1040:0x1045], getRepeated("Deseret", 0x5))
	copy(blocks[0x1045:0x1048], getRepeated("Shavian", 0x3))
	copy(blocks[0x1048:0x104b], getRepeated("Osmanya", 0x3))
	copy(blocks[0x1080:0x1084], getRepeated("Cypriot Syllabary", 0x4))
	copy(blocks[0x1084:0x1086], getRepeated("Imperial Aramaic", 0x2))
	copy(blocks[0x1090:0x1092], getRepeated("Phoenician", 0x2))
	copy(blocks[0x1092:0x1094], getRepeated("Lydian", 0x2))
	copy(blocks[0x1098:0x109a], getRepeated("Meroitic Hieroglyphs", 0x2))
	copy(blocks[0x109a:0x10a0], getRepeated("Meroitic Cursive", 0x6))
	copy(blocks[0x10a0:0x10a6], getRepeated("Kharoshthi", 0x6))
	copy(blocks[0x10a6:0x10a8], getRepeated("Old South Arabian", 0x2))
	copy(blocks[0x10b0:0x10b4], getRepeated("Avestan", 0x4))
	copy(blocks[0x10b4:0x10b6], getRepeated("Inscriptional Parthian", 0x2))
	copy(blocks[0x10b6:0x10b8], getRepeated("Inscriptional Pahlavi", 0x2))
	copy(blocks[0x10c0:0x10c5], getRepeated("Old Turkic", 0x5))
	copy(blocks[0x1100:0x1108], getRepeated("Brahmi", 0x8))
	copy(blocks[0x1108:0x110d], getRepeated("Kaithi", 0x5))
	copy(blocks[0x110d:0x1110], getRepeated("Sora Sompeng", 0x3))
	copy(blocks[0x1110:0x1115], getRepeated("Chakma", 0x5))
	copy(blocks[0x1118:0x111e], getRepeated("Sharada", 0x6))
	copy(blocks[0x1168:0x116d], getRepeated("Takri", 0x5))
	copy(blocks[0x1200:0x1240], getRepeated("Cuneiform", 0x40))
	copy(blocks[0x1300:0x1343], getRepeated("Egyptian Hieroglyphs", 0x43))
	copy(blocks[0x1680:0x16a4], getRepeated("Bamum Supplement", 0x24))
	copy(blocks[0x16f0:0x16fa], getRepeated("Miao", 0xa))
	copy(blocks[0x1b00:0x1b10], getRepeated("Kana Supplement", 0x10))
	copy(blocks[0x1d40:0x1d80], getRepeated("Mathematical Alphanumeric Symbols", 0x40))
	copy(blocks[0x1ee0:0x1ef0], getRepeated("Arabic Mathematical Alphabetic Symbols", 0x10))
	copy(blocks[0x2000:0x2a6e], getRepeated("CJK Unified Ideographs Extension B", 0xa6e))
	copy(blocks[0x2a70:0x2b74], getRepeated("CJK Unified Ideographs Extension C", 0x104))
	copy(blocks[0x2b74:0x2b82], getRepeated("CJK Unified Ideographs Extension D", 0xe))
	copy(blocks[0x2f80:0x2fa2], getRepeated("CJK Compatibility Ideographs Supplement", 0x22))
}

// Returns list of repeated string
func getRepeated(name string, nbRepeats int) (result []string) {
	for i := 0; i < nbRepeats; i++ {
		result = append(result, name)
	}

	return result
}

// Returns block name related to the char
func GetBlock(char int32) string {
	return blocks[int(char)>>blockRshift]
}
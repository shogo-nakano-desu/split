package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNormalizeArgsBasicCase(t *testing.T) {
	res := NormalizeArgs([]string{"-l", "10", "-a", "3", "test.txt"})
	expected := []string{"-l", "10", "-a", "3", "test.txt"}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestNormalizeArgsWithNoSpaceBetweenFlagAndArg(t *testing.T) {
	res := NormalizeArgs([]string{"-l10", "-a3", "test.txt"})
	expected := []string{"-l", "10", "-a", "3", "test.txt"}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestGenerateStrings(t *testing.T) {
	res, _ := GenerateStrings(2, "", 0)
	expected := []string{
		"aa",
		"ab",
		"ac",
		"ad",
		"ae",
		"af",
		"ag",
		"ah",
		"ai",
		"aj",
		"ak",
		"al",
		"am",
		"an",
		"ao",
		"ap",
		"aq",
		"ar",
		"as",
		"at",
		"au",
		"av",
		"aw",
		"ax",
		"ay",
		"az",
		"ba",
		"bb",
		"bc",
		"bd",
		"be",
		"bf",
		"bg",
		"bh",
		"bi",
		"bj",
		"bk",
		"bl",
		"bm",
		"bn",
		"bo",
		"bp",
		"bq",
		"br",
		"bs",
		"bt",
		"bu",
		"bv",
		"bw",
		"bx",
		"by",
		"bz",
		"ca",
		"cb",
		"cc",
		"cd",
		"ce",
		"cf",
		"cg",
		"ch",
		"ci",
		"cj",
		"ck",
		"cl",
		"cm",
		"cn",
		"co",
		"cp",
		"cq",
		"cr",
		"cs",
		"ct",
		"cu",
		"cv",
		"cw",
		"cx",
		"cy",
		"cz",
		"da",
		"db",
		"dc",
		"dd",
		"de",
		"df",
		"dg",
		"dh",
		"di",
		"dj",
		"dk",
		"dl",
		"dm",
		"dn",
		"do",
		"dp",
		"dq",
		"dr",
		"ds",
		"dt",
		"du",
		"dv",
		"dw",
		"dx",
		"dy",
		"dz",
		"ea",
		"eb",
		"ec",
		"ed",
		"ee",
		"ef",
		"eg",
		"eh",
		"ei",
		"ej",
		"ek",
		"el",
		"em",
		"en",
		"eo",
		"ep",
		"eq",
		"er",
		"es",
		"et",
		"eu",
		"ev",
		"ew",
		"ex",
		"ey",
		"ez",
		"fa",
		"fb",
		"fc",
		"fd",
		"fe",
		"ff",
		"fg",
		"fh",
		"fi",
		"fj",
		"fk",
		"fl",
		"fm",
		"fn",
		"fo",
		"fp",
		"fq",
		"fr",
		"fs",
		"ft",
		"fu",
		"fv",
		"fw",
		"fx",
		"fy",
		"fz",
		"ga",
		"gb",
		"gc",
		"gd",
		"ge",
		"gf",
		"gg",
		"gh",
		"gi",
		"gj",
		"gk",
		"gl",
		"gm",
		"gn",
		"go",
		"gp",
		"gq",
		"gr",
		"gs",
		"gt",
		"gu",
		"gv",
		"gw",
		"gx",
		"gy",
		"gz",
		"ha",
		"hb",
		"hc",
		"hd",
		"he",
		"hf",
		"hg",
		"hh",
		"hi",
		"hj",
		"hk",
		"hl",
		"hm",
		"hn",
		"ho",
		"hp",
		"hq",
		"hr",
		"hs",
		"ht",
		"hu",
		"hv",
		"hw",
		"hx",
		"hy",
		"hz",
		"ia",
		"ib",
		"ic",
		"id",
		"ie",
		"if",
		"ig",
		"ih",
		"ii",
		"ij",
		"ik",
		"il",
		"im",
		"in",
		"io",
		"ip",
		"iq",
		"ir",
		"is",
		"it",
		"iu",
		"iv",
		"iw",
		"ix",
		"iy",
		"iz",
		"ja",
		"jb",
		"jc",
		"jd",
		"je",
		"jf",
		"jg",
		"jh",
		"ji",
		"jj",
		"jk",
		"jl",
		"jm",
		"jn",
		"jo",
		"jp",
		"jq",
		"jr",
		"js",
		"jt",
		"ju",
		"jv",
		"jw",
		"jx",
		"jy",
		"jz",
		"ka",
		"kb",
		"kc",
		"kd",
		"ke",
		"kf",
		"kg",
		"kh",
		"ki",
		"kj",
		"kk",
		"kl",
		"km",
		"kn",
		"ko",
		"kp",
		"kq",
		"kr",
		"ks",
		"kt",
		"ku",
		"kv",
		"kw",
		"kx",
		"ky",
		"kz",
		"la",
		"lb",
		"lc",
		"ld",
		"le",
		"lf",
		"lg",
		"lh",
		"li",
		"lj",
		"lk",
		"ll",
		"lm",
		"ln",
		"lo",
		"lp",
		"lq",
		"lr",
		"ls",
		"lt",
		"lu",
		"lv",
		"lw",
		"lx",
		"ly",
		"lz",
		"ma",
		"mb",
		"mc",
		"md",
		"me",
		"mf",
		"mg",
		"mh",
		"mi",
		"mj",
		"mk",
		"ml",
		"mm",
		"mn",
		"mo",
		"mp",
		"mq",
		"mr",
		"ms",
		"mt",
		"mu",
		"mv",
		"mw",
		"mx",
		"my",
		"mz",
		"na",
		"nb",
		"nc",
		"nd",
		"ne",
		"nf",
		"ng",
		"nh",
		"ni",
		"nj",
		"nk",
		"nl",
		"nm",
		"nn",
		"no",
		"np",
		"nq",
		"nr",
		"ns",
		"nt",
		"nu",
		"nv",
		"nw",
		"nx",
		"ny",
		"nz",
		"oa",
		"ob",
		"oc",
		"od",
		"oe",
		"of",
		"og",
		"oh",
		"oi",
		"oj",
		"ok",
		"ol",
		"om",
		"on",
		"oo",
		"op",
		"oq",
		"or",
		"os",
		"ot",
		"ou",
		"ov",
		"ow",
		"ox",
		"oy",
		"oz",
		"pa",
		"pb",
		"pc",
		"pd",
		"pe",
		"pf",
		"pg",
		"ph",
		"pi",
		"pj",
		"pk",
		"pl",
		"pm",
		"pn",
		"po",
		"pp",
		"pq",
		"pr",
		"ps",
		"pt",
		"pu",
		"pv",
		"pw",
		"px",
		"py",
		"pz",
		"qa",
		"qb",
		"qc",
		"qd",
		"qe",
		"qf",
		"qg",
		"qh",
		"qi",
		"qj",
		"qk",
		"ql",
		"qm",
		"qn",
		"qo",
		"qp",
		"qq",
		"qr",
		"qs",
		"qt",
		"qu",
		"qv",
		"qw",
		"qx",
		"qy",
		"qz",
		"ra",
		"rb",
		"rc",
		"rd",
		"re",
		"rf",
		"rg",
		"rh",
		"ri",
		"rj",
		"rk",
		"rl",
		"rm",
		"rn",
		"ro",
		"rp",
		"rq",
		"rr",
		"rs",
		"rt",
		"ru",
		"rv",
		"rw",
		"rx",
		"ry",
		"rz",
		"sa",
		"sb",
		"sc",
		"sd",
		"se",
		"sf",
		"sg",
		"sh",
		"si",
		"sj",
		"sk",
		"sl",
		"sm",
		"sn",
		"so",
		"sp",
		"sq",
		"sr",
		"ss",
		"st",
		"su",
		"sv",
		"sw",
		"sx",
		"sy",
		"sz",
		"ta",
		"tb",
		"tc",
		"td",
		"te",
		"tf",
		"tg",
		"th",
		"ti",
		"tj",
		"tk",
		"tl",
		"tm",
		"tn",
		"to",
		"tp",
		"tq",
		"tr",
		"ts",
		"tt",
		"tu",
		"tv",
		"tw",
		"tx",
		"ty",
		"tz",
		"ua",
		"ub",
		"uc",
		"ud",
		"ue",
		"uf",
		"ug",
		"uh",
		"ui",
		"uj",
		"uk",
		"ul",
		"um",
		"un",
		"uo",
		"up",
		"uq",
		"ur",
		"us",
		"ut",
		"uu",
		"uv",
		"uw",
		"ux",
		"uy",
		"uz",
		"va",
		"vb",
		"vc",
		"vd",
		"ve",
		"vf",
		"vg",
		"vh",
		"vi",
		"vj",
		"vk",
		"vl",
		"vm",
		"vn",
		"vo",
		"vp",
		"vq",
		"vr",
		"vs",
		"vt",
		"vu",
		"vv",
		"vw",
		"vx",
		"vy",
		"vz",
		"wa",
		"wb",
		"wc",
		"wd",
		"we",
		"wf",
		"wg",
		"wh",
		"wi",
		"wj",
		"wk",
		"wl",
		"wm",
		"wn",
		"wo",
		"wp",
		"wq",
		"wr",
		"ws",
		"wt",
		"wu",
		"wv",
		"ww",
		"wx",
		"wy",
		"wz",
		"xa",
		"xb",
		"xc",
		"xd",
		"xe",
		"xf",
		"xg",
		"xh",
		"xi",
		"xj",
		"xk",
		"xl",
		"xm",
		"xn",
		"xo",
		"xp",
		"xq",
		"xr",
		"xs",
		"xt",
		"xu",
		"xv",
		"xw",
		"xx",
		"xy",
		"xz",
		"ya",
		"yb",
		"yc",
		"yd",
		"ye",
		"yf",
		"yg",
		"yh",
		"yi",
		"yj",
		"yk",
		"yl",
		"ym",
		"yn",
		"yo",
		"yp",
		"yq",
		"yr",
		"ys",
		"yt",
		"yu",
		"yv",
		"yw",
		"yx",
		"yy",
		"yz",
		"za",
		"zb",
		"zc",
		"zd",
		"ze",
		"zf",
		"zg",
		"zh",
		"zi",
		"zj",
		"zk",
		"zl",
		"zm",
		"zn",
		"zo",
		"zp",
		"zq",
		"zr",
		"zs",
		"zt",
		"zu",
		"zv",
		"zw",
		"zx",
		"zy",
		"zz",
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestGenerateStringsLotOfStrs(t *testing.T) {
	res, _ := GenerateStrings(4, "", 0)
	resLen := len(res)
	expectedLen := 456976
	if resLen != expectedLen {
		t.Errorf("expected %v, got %v", expectedLen, resLen)
	}
}

func TestGenerateStringsZeroLength(t *testing.T) {
	_, err := GenerateStrings(0, "", 0)
	expected := fmt.Errorf("Error: suffix length must be greater than 0")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}

}

func TestGenerateStringsTooBigLength(t *testing.T) {
	_, err := GenerateStrings(6, "", 0)
	expected := fmt.Errorf("Error: suffix length must be less than or equal to 5")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestIllegalArgsChecker(t *testing.T) {
	err := IllegalArgsChecker(Args{LineCount: 1, FileCount: 0, ByteSize: 0, Args: []string{"-l", "10", "-a", "3", "test.txt"}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestIllegalArgsCheckerDuplicateArgs(t *testing.T) {
	err := IllegalArgsChecker(Args{LineCount: 3, FileCount: 0, ByteSize: 0, Args: []string{"-l", "10", "-l", "3", "test.txt"}})
	expected := fmt.Errorf(
		`usage: split [-l line_count] [-a suffix_length] [file [prefix]]
			split -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]
			split -n chunk_count [-a suffix_length] [file [prefix]]
			split -p pattern [-a suffix_length] [file [prefix]]`)
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestIllegalArgsCheckerMultipleArgs(t *testing.T) {
	err := IllegalArgsChecker(Args{LineCount: 3, FileCount: 0, ByteSize: 1, Args: []string{"-b", "1", "-l", "3", "test.txt"}})
	expected := fmt.Errorf(
		`usage: split [-l line_count] [-a suffix_length] [file [prefix]]
			split -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]
			split -n chunk_count [-a suffix_length] [file [prefix]]
			split -p pattern [-a suffix_length] [file [prefix]]`,
	)
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestIllegalArgsCheckerUnknownArgs(t *testing.T) {
	err := IllegalArgsChecker(Args{LineCount: 2, FileCount: 0, ByteSize: 0, Args: []string{"-l", "2", "-t", "3", "test.txt"}})
	expected := fmt.Errorf("Error: unknown option -t")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}

}

func TestIllegalArgsCheckerInvalidValue(t *testing.T) {
	err := IllegalArgsChecker(Args{LineCount: 0, FileCount: 0, ByteSize: 0, Args: []string{"-l", "0", "test.txt"}})
	expected := fmt.Errorf("error: 0: illegal line count")
	if err.Error() != expected.Error() {
		t.Errorf("expected %v, got %v", expected, err)
	}
}

func TestParseArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }() // テスト後にos.Argsを元に戻す

	os.Args = []string{"./main", "-l", "10", "-a", "5"}
	fs := flag.NewFlagSet("./main", flag.ContinueOnError)
	res, _ := ParseArgs(fs)

	expected := ParseArgsResult{
		LineCount: 10,
		FileCount: 0,
		ByteSize:  0,
		SuffixLen: 5,
		Args:      []string{"-l", "10", "-a", "5"},
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestParseArgsInvalidSemantics(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }() // テスト後にos.Argsを元に戻す

	os.Args = []string{"./main", "-l", "10", "-n", "5"}
	fs := flag.NewFlagSet("./main", flag.ContinueOnError)
	res, _ := ParseArgs(fs)

	expected := ParseArgsResult{
		LineCount: 10,
		FileCount: 5,
		ByteSize:  0,
		SuffixLen: 2,
		Args:      []string{"-l", "10", "-n", "5"},
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestGetFileName(t *testing.T) {
	tests := []struct {
		nonFlagArgs []string
		input       string
		expected    string
		shouldError bool
	}{
		{
			nonFlagArgs: []string{"filename.txt"},
			expected:    "filename.txt",
		},
		{
			nonFlagArgs: []string{},
			input:       "inputfile.txt\n",
			expected:    "inputfile.txt",
		},
		{
			nonFlagArgs: []string{},
			input:       "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(bytes.NewBufferString(tt.input))
		result, err := GetFileName(tt.nonFlagArgs, reader)

		if (err != nil) != tt.shouldError {
			t.Fatalf("expected error: %v, got: %v", tt.shouldError, err)
		}

		if result != tt.expected {
			t.Fatalf("expected: %s, got: %s", tt.expected, result)
		}
	}
}

func TestGetFileNameWithValidFileName(t *testing.T) {
	test := struct {
		nonFlagArgs []string
		expected    string
	}{
		nonFlagArgs: []string{"filename.txt"},
		expected:    "filename.txt",
	}

	reader := bufio.NewReader(bytes.NewBufferString(""))
	result, err := GetFileName(test.nonFlagArgs, reader)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != test.expected {
		t.Fatalf("expected: %s, got: %s", test.expected, result)
	}
}

func TestGetFileNameWithInputString(t *testing.T) {
	test := struct {
		nonFlagArgs []string
		input       string
		expected    string
	}{
		nonFlagArgs: []string{},
		input:       "inputfile.txt\n",
		expected:    "inputfile.txt",
	}

	reader := bufio.NewReader(bytes.NewBufferString(test.input))
	result, err := GetFileName(test.nonFlagArgs, reader)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != test.expected {
		t.Fatalf("expected: %s, got: %s", test.expected, result)
	}
}

func TestGetFileNameWithError(t *testing.T) {
	test := struct {
		nonFlagArgs []string
		input       string
		shouldError bool
	}{
		nonFlagArgs: []string{},
		input:       "",
		shouldError: true,
	}

	reader := bufio.NewReader(bytes.NewBufferString(test.input))
	_, err := GetFileName(test.nonFlagArgs, reader)

	if (err == nil) == test.shouldError {
		t.Fatalf("expected error: %v, got: %v", test.shouldError, err)
	}
}

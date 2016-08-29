package structure

var SceneList = map[string]*SceneManager{
	"Title": &SceneManager{
		next: NewScene(imgList[0].name, imgList[0].fl),
	},
	"Opener": &SceneManager{
		prev: NewTitleScene(),
		next: NewScene(imgList[1].name, imgList[1].fl),
	},
}

// TODO: fix error when hitting <left> from the "opener" should go back to Title
func init() {
	var x int
	for x=1; x < (len(imgList)-1); x++ {
		SceneList[imgList[x].name] = &SceneManager{
			prev: NewScene(imgList[x - 1].name, imgList[x - 1].fl),
			next: NewScene(imgList[x + 1].name, imgList[x + 1].fl),
		}
	}

	SceneList[imgList[len(imgList)-1].name] = &SceneManager{
		prev: NewScene(imgList[len(imgList)-2].name, imgList[len(imgList)-2].fl),
		next: NewEmblemScene(),
	}
}

type imgFile struct {
	name string
	fl string
}

var imgList = []imgFile{
	{"Opener", "images/IMG_0462.png"},
	{"Two", "images/IMG_0463.png"},
	{"Three", "images/IMG_0464.png"},
	{"Four", "images/IMG_0465.png"},
	{"Five", "images/IMG_0466.png"},
	{"Six", "images/IMG_0467.png"},
	{"Seven", "images/IMG_0468.png"},
	{"Eight", "images/IMG_0521.png"},
	{"Nine", "images/IMG_0522.png"},
	{"Ten", "images/IMG_0523.png"},
	{"Eleven", "images/IMG_0524.png"},
	{"Twelve", "images/IMG_0526.png"},
	{"Thirteen", "images/IMG_0527.png"},
	{"Fourteen", "images/IMG_0540.png"},
	{"Fifteen", "images/IMG_0541.png"},
	{"Sixteen", "images/IMG_0542.png"},
	{"Seventeen", "images/IMG_0543.png"},
}
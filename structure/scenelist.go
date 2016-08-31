package structure

var SceneList = map[string]*SceneManager{
	"Title": &SceneManager{
		next: NewScene(imgList[0].name, imgList[0].fl, imgList[0].txt),
	},
	"Opener": &SceneManager{
		prev: NewTitleScene(),
		next: NewScene(imgList[1].name, imgList[1].fl, imgList[1].txt),
	},
}

// TODO: fix error when hitting <left> from the "opener" should go back to Title
func init() {
	var x int
	for x=1; x < (len(imgList)-1); x++ {
		SceneList[imgList[x].name] = &SceneManager{
			prev: NewScene(imgList[x - 1].name, imgList[x - 1].fl, imgList[x - 1].txt),
			next: NewScene(imgList[x + 1].name, imgList[x + 1].fl, imgList[x + 1].txt),
		}
	}

//	tmp := len(imgList)

	//SceneList[imgList[tmp-1].name] = &SceneManager{
	//	prev: NewScene(imgList[tmp -1].name, imgList[tmp-1].fl),
	//	next: NewDeckScene(),
	//}
	//
	//SceneList[imgList[len(imgList)-1].name] = &SceneManager{
	//	prev: NewScene(imgList[len(imgList)-2].name, imgList[len(imgList)-2].fl),
	//	next: NewEmblemScene(),
	//}

	SceneList["Seventeen"] = &SceneManager{
		prev: NewScene(imgList[15].name, imgList[15].fl, imgList[15].txt),
		next: NewDeckScene(),
	}

	SceneList["EndGame"] = &SceneManager{
		prev: NewDeckScene(),
		next: NewEmblemScene(),
	}
}

type imgFile struct {
	name string
	fl string
	txt string
}

var imgList = []imgFile{
	{"Opener", "images/IMG_0462.png", "This is Thunder Snowflake"},
	{"Two", "images/IMG_0463.png", "He's a super hero,"},
	{"Three", "images/IMG_0464.png", "stopping bad guys wherever he sees them."},
	{"Four", "images/IMG_0465.png", "blah"},
	{"Five", "images/IMG_0466.png", "blah"},
	{"Six", "images/IMG_0467.png", "stuff"},
	{"Seven", "images/IMG_0468.png", "Things"},
	{"Eight", "images/IMG_0521.png", ""},
	{"Nine", "images/IMG_0522.png", "Oops"},
	{"Ten", "images/IMG_0523.png", "missed that one"},
	{"Eleven", "images/IMG_0524.png", "These are the Pack Rats"},
	{"Twelve", "images/IMG_0526.png", "Mean brothers turned into rats"},
	{"Thirteen", "images/IMG_0527.png", "Thunder's sidekick is Captain Coconut"},
	{"Fourteen", "images/IMG_0540.png", "not much more to say"},
	{"Fifteen", "images/IMG_0541.png", "blah blah"},
	{"Sixteen", "images/IMG_0542.png", "I do not say 'blah blah blah"},
	{"Seventeen", "images/IMG_0543.png", "Bye now"},
}
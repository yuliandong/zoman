package imgana

import (
	"fmt"

	"github.com/Comdex/imgo"
)

func Example() {
	//如果读取出错会panic,返回图像矩阵img
	//img[height][width][4],height为图像高度,width为图像宽度
	//img[height][width][4]为第height行第width列上像素点的RGBA数值数组，值范围为0-255
	//如img[150][20][0]是150行20列处像素的红色值,img[150][20][1]是150行20列处像素的绿
	//色值，img[150][20][2]是150行20列处像素的蓝色值,img[150][20][3]是150行20列处像素
	//的alpha数值,一般用作不透明度参数,如果一个像素的alpha通道数值为0%，那它就是完全透明的.
	img := imgo.MustRead("img/example.jpg")
	defer fmt.Println(img[10][10][0])
	//对原图像矩阵进行日落效果处理
	img2 := imgo.SunsetEffect(img)
	
	img3 := imgo.

	//保存为jpeg,100为质量，1-100
	err := imgo.SaveAsJPEG("img/sunset.jpeg", img2, 100)

}

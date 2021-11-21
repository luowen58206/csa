package lv3

/*
	@Title : Lv3
	@Description : 不同用户对不同视频的简单操作
	@auth : 李江渝
	@param : 结构体绑定方法 param01:*Author param02:*Video 用户对视频的各种操作
			 ReleaseVideo() param01:author param02:videoName 用户发布视频名和作者
			 Run() param01:*Author param02:*Video 开启协程，执行操作
	@return : ReleaseVideo() return 包含作者名和视频名的结构体
*/


import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Wait 等待协程完成工作后在结束
var Wait sync.WaitGroup

// Author 定义一个用户的结构体
type Author struct {
	Name string
	VIP bool			//会员
	Icon string 		// 头像
	Signature string	//签名
	Focus int			//关注人数
}

// Video 定义一个视频结构体
type Video struct {
	Author string	//视频作者
	Name string 	//视频名字
	LinksNum int 	//点赞数
	PlayVolume int  //播放量
	ReprintsNum int //转载数
	CollectNum int	//收藏数
	CoinsNum int 	//投币数
}

// Link 点赞方法
func (this *Video) Link(user *Author,videoName *Video)  {
	videoName.LinksNum++
	fmt.Printf("%s link %s\n",user.Name,videoName.Name)
}

// Collect 收藏方法
func (this *Video) Collect(user *Author,videoName *Video)  {
	videoName.CollectNum++
	fmt.Printf("%s collect %s\n",user.Name,videoName.Name)
}

// Play 播放方法
func (this *Video) Play(user *Author,videoName *Video)  {
	videoName.PlayVolume++
	fmt.Printf("%s play %s\n",user.Name,videoName.Name)
}

// Reprints 转载方法
func (this *Video) Reprints(user *Author,videoName *Video)  {
	videoName.ReprintsNum++
	fmt.Printf("%s reprints %s\n",user.Name,videoName.Name)
}

// Coins 投币方法
func (this *Video) Coins(user *Author,videoName *Video)  {
	videoName.CoinsNum++
	fmt.Printf("%s coin to  %s\n",user.Name,videoName.Name)
}

// ThreeConsecutive 一键三连方法
func (this *Video) ThreeConsecutive(user *Author,videoName *Video)  {
	videoName.LinksNum++
	videoName.CoinsNum++
	videoName.ReprintsNum++
	//fmt.Printf("%s Three consecutive %s\n",user.Name,videoName.Name)
}

// ReleaseVideo 发布视频方法
func ReleaseVideo(author string,videoName string) (resVideo Video) {
	resVideo.Author = author
	resVideo.Name = videoName
	fmt.Printf("%s post an video %s\n",author,videoName)
	return resVideo
}

func Lv3()  {
	//定义两个用户User01，user02
	var user01 = Author{
		Name: "jack",
		VIP: true,
		Icon: "Monkey",
		Signature: "i am a bad duck",
		Focus: 100,
	}
	var user02 = Author{
		Name: "Marry",
		VIP: false,
		Icon: "Cat",
		Signature: "i am a bad duck",
		Focus: 10000,
	}

	//定义两个视频video01，video02
	var video01 Video
	video01 = ReleaseVideo("jack","EveryDayChallenge")
	var video02 Video
	video02 = ReleaseVideo("Mary","EveryMonthChallenge")

	//开启协程，并且执行用户操作视频的方法
	Wait.Add(4)
	go Run(&user01,&video01)
	go Run(&user02,&video01)
	go Run(&user01,&video02)
	go Run(&user02,&video02)

	//等待协程结束
	Wait.Wait()

	//打印两个视频的具体信息
	fmt.Println(video01)
	fmt.Println(video02)

}

// Run 执行user的各种方法
func Run(user *Author,video *Video)  {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10) + 1
	for i := 0; i < num; i++ {
		video.Collect(user, video)
		video.Coins(user, video)
		video.Reprints(user, video)
		video.ThreeConsecutive(user, video)
		video.Play(user, video)
		video.Link(user, video)
	}
	Wait.Done()
}


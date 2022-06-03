package controller

import (
	. "douyin_mine/config"
	"douyin_mine/service"
	"douyin_mine/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type PublishController struct {
}

const FILEMAXSIZE = 50 << 20 //最多传输50M大小的文件

// /public/action 投稿接口
func (pc *PublishController) PostAction(ctx iris.Context) mvc.Result {
	token := ctx.FormValue("token")
	userid, err := utils.CheckToken(token)
	Log.Println(err)
	if err == redis.Nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 100,
				Status_Msg:  "请先登录",
			},
		}
	}
	Log.Println(userid)
	title := ctx.FormValue("title")
	f, fh, err := ctx.FormFile("data")
	if err != nil {
		Log.Println(err)
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 500,
				Status_Msg:  "文件读取失败",
			},
		}
	}
	Log.Println("文件：" + fh.Filename + " 大小为" + strconv.Itoa(int(fh.Size)))
	if fh.Size > FILEMAXSIZE {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 400,
				Status_Msg:  "文件超过规定大小",
			},
		}
	}
	content, err := ioutil.ReadAll(f) //读取文件
	if err != nil {
		Log.Println(err)
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 500,
				Status_Msg:  "文件读取失败",
			},
		}
	}

	//TODO 创建本地视频资源以及本地图片资源
	playerurl := CreatePath(fh)
	index := strings.LastIndexAny(playerurl, ".")
	pictureurl := playerurl[:index+1] + "jpg"
	Log.Println("生成url:", playerurl)
	Log.Println("Picture utl:", pictureurl)
	//target, err := os.OpenFile(AppConfig.GetString("resources.video.path")+string(os.PathSeparator)+playerurl, os.O_EXCL|os.O_CREATE, 0666)
	videoname := AppConfig.GetString("resources.video.path") + string(os.PathSeparator) + playerurl
	picturename := AppConfig.GetString("resources.picture.path") + string(os.PathSeparator) + pictureurl

	index = strings.LastIndexAny(picturename, string(os.PathSeparator))
	filepath := picturename[:index]
	if _, err := os.Stat(filepath); err != nil {
		Log.Println("目录不存在，正在创建目录:" + filepath)
		os.MkdirAll(filepath, 0766)
		Log.Println("目录创建完成")
	}

	index = strings.LastIndexAny(videoname, string(os.PathSeparator))
	filepath = videoname[:index]
	if _, err := os.Stat(filepath); err != nil {
		Log.Println("目录不存在，正在创建目录:" + filepath)
		os.MkdirAll(filepath, 0766)
		Log.Println("目录创建完成")
	}
	target, err := os.Create(videoname)

	if err != nil {
		Log.Println(err)
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 502,
				Status_Msg:  "文件存储失败",
			},
		}
	}
	target.Write(content)
	target.Close()

	err = ffmpeg_go.Input(videoname).Output(picturename, ffmpeg_go.KwArgs{"vframes": "1"}).Run()
	if err != nil {
		Log.Println(err)
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 502,
				Status_Msg:  "封面生成失败",
			},
		}
	}
	//Database.AutoMigrate(&service.Video{}) //若数据库表不存在则初始化数据库表
	f.Close()
	var video service.Video
	video.Video_location = playerurl
	video.Video_authorid = userid
	video.Video_picture_location = AppConfig.GetString("resources.picture.defaultrelativepath")
	video.Video_state = 200 //直接设置为通过不加审核
	video.Video_title = title
	video.Video_latest_time = time.Now()
	video.Video_introduction = "简介功能待客户端实现提交"
	err = Database.Create(&video).Error
	if err != nil {
		fmt.Println("%v", err)
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 504,
				Status_Msg:  "数据存储到数据库出错",
			},
		}
	}
	return mvc.Response{
		Object: statusResponse{
			Status_Code: 0,
			Status_Msg:  "发布成功",
		},
	}
}

// /publish/list 发布列表
func (pc *PublishController) GetList(ctx iris.Context) mvc.Result {
	token := ctx.URLParam("token")
	userid, err := utils.CheckToken(token) //验证token是否可用
	Log.Println("用户：" + strconv.Itoa(userid) + "身份验证")
	if err == redis.Nil {
		return mvc.Response{
			Object: listResponse{
				statusResponse: statusResponse{
					Status_Code: 100,
					Status_Msg:  "请先登录",
				},
				Video_list: nil,
			},
		}
	}
	var videos []service.Video
	Database.Where("`video_authorid` = ?", userid).Order("video_latest_time DESC").Find(&videos)
	Log.Println("查询视频列表")
	videolist := make([]service.VideoJSON, 0, 30)
	for _, video := range videos {
		videolist = append(videolist, service.GetVideoJSON(userid, video))
	}
	return mvc.Response{
		Object: listResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
			},
			Video_list: videolist,
		},
	}
}

//生成用户投稿文件路径 index为.索引值
func CreatePath(fh *multipart.FileHeader) string {
	Log.Println("正在生成投稿文件路径")
	now_time := time.Now()
	year := strconv.Itoa(now_time.Year())
	now_time.Month()
	month := strconv.Itoa(int(now_time.Month()))
	index := strings.LastIndexAny(fh.Filename, ".")
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("uuid生成失败")
	}
	filename := strings.ReplaceAll(uid.String(), "-", "")
	filetype := fh.Filename[index+1:]
	Log.Println("投稿文件路径生成成功")
	return year + string(os.PathSeparator) + month + string(os.PathSeparator) + filename + "." + filetype
}

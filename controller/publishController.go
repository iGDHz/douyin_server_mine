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
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type PublicController struct {
}

type VideoJSON struct {
	Id             int              `json:"id"`
	Author         service.UserJSON `json:"author"`
	Play_url       string           `json:"play_url"`
	Cover_url      string           `json:"cover_url"`
	Favorite_count int              `json:"favorite_count"`
	Comment_count  int              `json:"comment_count"`
	Is_favorite    int              `json:"is_favorite"`
	Title          string           `json:"title"`
}

type listResponse struct {
	statusResponse
	video_list []VideoJSON `json:"video_list"`
}

// /public/action 投稿接口
func (pc *PublicController) PostAction(ctx iris.Context) mvc.Result {
	err := Rdb.Get(RdbContext, ctx.FormValue("token")).Err()
	if err == redis.Nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 100,
				Status_Msg:  "请先登录",
			},
		}
	}
	userid := Rdb.Get(RdbContext, ctx.FormValue("token")) //根据token获取用户id
	fmt.Println(userid)
	title := ctx.FormValue("title")
	f, fh, err := ctx.FormFile("data")
	content, err := ioutil.ReadAll(f) //读取文件
	if err != nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 500,
				Status_Msg:  "文件读取失败",
			},
		}
	}
	playerurl := CreatePath(fh)
	target, err := os.OpenFile(AppConfig.GetString("resources.video.path")+string(os.PathSeparator)+playerurl, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		return mvc.Response{
			Object: statusResponse{
				Status_Code: 502,
				Status_Msg:  "文件存储失败",
			},
		}
	}
	Database.AutoMigrate(&service.Video{}) //若数据库表不存在则初始化数据库表
	target.Write(content)
	var video service.Video
	video.Video_location = playerurl
	video.Video_authorid, _ = strconv.Atoi(userid.Val())
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

// /public/list 发布列表
func (pc *PublicController) GetList(ctx iris.Context) mvc.Result {
	token := ctx.URLParam("token")
	_, err := utils.CheckToken(token) //验证token是否可用
	if err == redis.Nil {
		return mvc.Response{
			Object: listResponse{
				statusResponse: statusResponse{
					Status_Code: 100,
					Status_Msg:  "请先登录",
				},
				video_list: nil,
			},
		}
	}
	userid := ctx.URLParam("user_id") //获取用户id
	var videos []service.Video
	Database.Where("`video_authorid` = ?", userid).Order("video_latest_time DESC").Find(&videos)
	videolist := make([]VideoJSON, 0, 30)
	for _, video := range videos {
		uid, _ := strconv.Atoi(userid)
		videolist = append(videolist, service.GetVideoJSON(uid, video))
	}
	return mvc.Response{
		Object: listResponse{
			statusResponse: statusResponse{
				Status_Code: 0,
			},
			video_list: videolist,
		},
	}
}

//生成用户投稿文件路径
func CreatePath(fh *multipart.FileHeader) string {
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
	return year + string(os.PathSeparator) + month + string(os.PathSeparator) + filename + "." + filetype
}

package authc

//import "github.com/gin-gonic/gin"

import (
	"image/color"
	"image/png"
	"net/http"
	"strings"
	"travel-agency/core/dto"
	"travel-agency/entity"
	dbfactory "travel-agency/factory/db"
	"travel-agency/util"

	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	util2 "github.com/lyr-2000/goUtil/util"
	log "github.com/sirupsen/logrus"
)

type LoginFilter struct {
}

//db
var db *gorm.DB = dbfactory.GetGormDB()

//是否有这个用户
func isUserPresent(username string, pwd string) (bool, entity.User) {
	user := entity.User{}
	db.Raw("select * from t_user  where username = ? ", username).First(&user)
	log.Printf("user %v", &user)
	//加盐加密
	pwd2 := util2.Md5([]byte(pwd + user.Salt))

	return user.Password == pwd2, user
}
func getSession(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	return session
}

//free/code
//生成登录验证码
var cap *captcha.Captcha = captcha.New()

//生成验证码
func GenerateCode(c *gin.Context) {
	//cap
	// 设置字体
	cap.SetFont("./static/comic.ttf")
	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	// 创建验证码 4个字符 captcha.NUM 字符模式数字类型
	// 返回验证码图像对象以及验证码字符串 后期可以对字符串进行对比 判断验证
	img, str := cap.Create(4, captcha.NUM)
	s := getSession(c)
	s.Set("code", str)
	s.Save()
	png.Encode(c.Writer, img)
}

//loginout
func Logout(c *gin.Context) {
	s := getSession(c)
	s.Delete("login")
	s.Delete("username")
	err := s.Save()
	// s.Clear()
	log.Println("err ", err)

	c.Redirect(http.StatusMovedPermanently, "/login")

}

// login
func HandleLogin(c *gin.Context) {
	session := getSession(c)
	code1 := session.Get("code")
	code2 := c.PostForm("code")
	if code1 != code2 {
		log.Printf("登录失败，验证码不正确 %s , %s", code1, code2)
		util.RenderHtml(c, "login.html", gin.H{
			"errInfo": "验证码不正确",
		})
		return
	} else {
		username := c.PostForm("username")
		pwd := c.PostForm("password")
		//username and pwd
		if login, user := isUserPresent(username, pwd); login {
			//设置登录成功
			log.Printf("登录成功")
			session.Set("login", true)
			//session.Set("user",user)
			session.Set("username", user.Username)
			err := session.Save()

			log.Printf("登录成功--- %v ,err=%v", &user, err)
			//重定向
			c.Redirect(http.StatusMovedPermanently, "/index")
			//util.RenderHtml(c,"base.html",nil)
			return
		} else {
			log.Printf("登录失败")
			util.RenderHtml(c, "login.html", gin.H{
				"errInfo": "密码不正确",
			})
			return
		}
	}
}

func ChangePassword(c *gin.Context) {
	oldpwd := c.PostForm("oldpwd")
	newpwd := c.PostForm("newpwd")
	againPwd := c.PostForm("confirmpwd")
	log.Printf("old=%s,new=%s,again=%s", oldpwd, newpwd, againPwd)
	if newpwd == againPwd {
		s := getSession(c)
		username, convertOk := s.Get("username").(string)
		//suser,convertOK := s.Get("username").(entity.User)
		if convertOk == false {
			log.Printf("转化失败 %v", username)
		} else {
			ok, _ := isUserPresent(username, oldpwd)
			if ok {
				//update
				db.Model(&entity.User{}).Where("username=?", username).Updates(&entity.User{
					Password: newpwd,
				})
				util.RenderJson(c, dto.Ok("转化成功"))
			} else {
				util.RenderJson(c, dto.Fail("密码不正确", "密码不正确，修改失败"))
			}
		}

	} else {
		//util.RenderHtml(c,"changePwd.html",gin.H{
		//	"errInfo":"输入错误",
		//})
		util.RenderJson(c, dto.Fail("密码不一致", "密码不一致"))
	}
}

func IsStaticPath(path *string) bool {
	if path == nil {
		return true
	}
	return strings.Contains(*path, "/public/") || strings.Contains(*path, "/static/") ||
		strings.Contains(*path, "/login") || strings.Contains(*path, "/free/")
}

func (*LoginFilter) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := getSession(c)
		if session == nil {
			p := c.Request.URL.Path
			log.Println("path := ", p)
			if IsStaticPath(&p) {
				c.Next()
				return
			}
			//验证不通过，不会调用后面的方法
			c.Abort()
			c.HTML(http.StatusOK, "login.html", gin.H{
				"errorInfo": "出错了",
			})
			return

		} else {
			log.Println("handler verify sesssion")
			p := c.Request.URL.Path
			log.Println("path=", p)
			if IsStaticPath(&p) {
				c.Next()
				return
			}

			//fmt.Println("继续00")
			userLogin := session.Get("login")

			if userLogin == nil {
				c.Abort()
				log.Println("需要登录 %v", userLogin)
				c.HTML(http.StatusOK, "login.html", gin.H{
					"errorInfo": "你需要登录",
				})
				return
			}
			log.Printf("已经登录，执行")
			//执行下一步
			c.Next()
		}

	}
}

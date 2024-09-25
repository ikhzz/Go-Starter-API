package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"time"

	"starterapi/common/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type commonUsecase struct {
	contextTimeout time.Duration
	cr             models.CommonRepository
	v              *validator.Validate
	jwtKey         string
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func NewCommonUsecase(timeout time.Duration, cr models.CommonRepository) models.CommonUsecase {
	jwtKey := viper.GetString("key.jwt")
	// SignedString = []byte(getKey)
	if jwtKey == "" {
		panic("jwt key is missing")
	}

	v := validator.New()
	v.RegisterValidation("alpha_num", containsAlphaNum)

	return &commonUsecase{
		contextTimeout: timeout,
		cr:             cr,
		v:              v,
		jwtKey:         jwtKey,
	}
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func identifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}

	return fmt.Sprintf("pc:%x", pc)
}

// PanicCatcher is use for collecting panic that happened in endpoint
func (g *commonUsecase) PanicCatcher(mw io.Writer) gin.HandlerFunc {
	fmt.Println("panic catcher")
	return func(c *gin.Context) {
		defer func() {
			_ = c.Request.ParseForm()
			reqBodyStr := ""
			for key, val := range c.Request.Form {
				reqBodyStr += key + ":"
				for i, v := range val {
					reqBodyStr += v
					if i+1 != len(val) {
						reqBodyStr += ","
					} else {
						reqBodyStr += " "
					}
				}
			}
			rec := recover()
			if rec != nil {
				user := "unknown user"
				device := "unknown device"
				userdata, isExist := c.Get("user")
				if isExist {
					user = userdata.(string)
				}
				devicedata, isExist2 := c.Get("device")
				if isExist2 {
					device = devicedata.(string)
				}
				fmt.Fprintf(mw, `level=error datetime="%s" ip=%s method=%s url="%s" user="%s" device="%s" panic=%v trace=%v`+"\n",
					time.Now().Format(time.RFC1123),
					c.ClientIP(),
					c.Request.Method,
					c.Request.URL.String(),
					user,
					device,
					rec,
					identifyPanic(),
				)
				// fmt.Println(string(debug.Stack()))
				// errPanic := fmt.Sprintf("Endpoint: %s - panic: %v", c.Request.RequestURI, rec)
				// errorcollector.WritePanic(errPanic, debug.Stack())
				var message []string
				message = append(message, viper.GetString("default_unhandled_error"))
				c.JSON(500, gin.H{
					"status":  false,
					"message": message,
					"data":    new(struct{}),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// CustomLogger provide custom log
func (g *commonUsecase) CustomLogger(mw io.Writer) gin.HandlerFunc {
	fmt.Println("custom logger")
	return func(c *gin.Context) {
		t := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		latency := time.Since(t)
		_ = c.Request.ParseForm()
		reqBodyStr := ""
		for key, val := range c.Request.Form {
			reqBodyStr += key + ":"
			for i, v := range val {
				reqBodyStr += "'" + v + "'"
				if i+1 != len(val) {
					reqBodyStr += ","
				} else {
					reqBodyStr += " "
				}
			}
		}
		reqHeaderStr := ""
		for key, val := range c.Request.Header {
			reqHeaderStr += key + ":"
			for i, v := range val {
				reqHeaderStr += v
				if i+1 != len(val) {
					reqHeaderStr += ","
				} else {
					reqHeaderStr += " "
				}
			}
		}
		res := blw.body.String()
		contentType := blw.ResponseWriter.Header().Get("Content-Type")
		contentTypeSplit := strings.Split(contentType, ";")
		if len(contentTypeSplit) > 0 {
			if contentTypeSplit[0] == "text/html" {
				var re = regexp.MustCompile(`/\s+|\n+|\r/`) //make html string to one line
				res = re.ReplaceAllString(res, "")
				res += " \n"
			}
		}
		if res == "" {
			res = " \n"
		}
		user := "unknown user"
		device := "unknown device"
		userdata, isExist := c.Get("user")
		if isExist {
			user = userdata.(string)
		}
		devicedata, isExist2 := c.Get("device")
		if isExist2 {
			device = devicedata.(string)
		}
		fmt.Fprintf(mw, `level=info datetime="%s" ip=%s method=%s url="%s" proto=%s status=%d latency=%s user="%s" device="%s" req_header:"%s" req_body="%s" response=%s`,
			t.Format(time.RFC1123),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.String(),
			c.Request.Proto,
			c.Writer.Status(),
			latency,
			user,
			device,
			reqHeaderStr,
			reqBodyStr,
			res,
		)
	}
}

func (g *commonUsecase) JwtMiddleware(c *gin.Context) {
	jwtToken := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if c.Request.Header.Get("Authorization") == "" || len(jwtToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.GeneralResponse{
			Message:    []string{"Token is required"},
			Status:     false,
			StatusCode: http.StatusUnauthorized,
			Data:       []models.EmptyResponse{},
		})
		return
	}

	token, err := jwt.Parse(jwtToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(g.jwtKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.GeneralResponse{
			Message:    []string{"bad jwt token"},
			Status:     false,
			StatusCode: http.StatusUnauthorized,
			Data:       []models.EmptyResponse{},
		})
		return
	}

	t, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.GeneralResponse{
			Message:    []string{"bad claims"},
			Status:     false,
			StatusCode: http.StatusUnauthorized,
			Data:       []models.EmptyResponse{},
		})
		return
	}

	c.Set("user_id", t["id"])
	c.Set("device", t["device"])
	id := c.GetString("user_id")
	device := c.GetString("device")
	data, errUser := g.cr.FindUserByUid(id)
	if errUser != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.GeneralResponse{
			Message:    []string{"user not found"},
			Status:     false,
			StatusCode: http.StatusUnauthorized,
			Data:       []models.EmptyResponse{},
		})
		return
	}

	if time.Now().After(data.TokenExpire) {
		c.AbortWithStatusJSON(http.StatusRequestTimeout, models.GeneralResponse{
			Message:    []string{"token expired"},
			Status:     false,
			StatusCode: http.StatusUnauthorized,
			Data:       []models.EmptyResponse{},
		})
		return
	}

	if device == "WEB" && data.Token != jwtToken[1] {
		c.AbortWithStatusJSON(http.StatusRequestTimeout, models.GeneralResponse{
			Message:    []string{"token invalid"},
			Status:     false,
			StatusCode: http.StatusUnauthorized,
			Data:       []models.EmptyResponse{},
		})
		return
	}

	toUpdate := make(map[string]interface{}, 0)
	toUpdate["token_expire"] = time.Now().Add(time.Hour * time.Duration(viper.GetInt("timeout.jwt")))
	g.cr.PutUser(data.UIDUser, toUpdate)

	c.Next()
}

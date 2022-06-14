package usecase 

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
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
func (g *GeneralUsecase) PanicCatcher(mw io.Writer) gin.HandlerFunc {
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
func (g *GeneralUsecase) CustomLogger(mw io.Writer) gin.HandlerFunc {
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

// CheckRoute will check is route authorized to access
// can be implemented with db
func (g *GeneralUsecase) CheckRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "" {
			path = "/"
		}
		// get token setting and path available
		checkPath := viper.GetInt("route."+path)
		checkToken := viper.GetInt("route.token"+path)
		// if token is required
		if checkToken == 1 {
			g.TokenDecrypt(c)
		}

		if checkPath == 1{
			c.Next()
		} else {
			c.JSON(403, gin.H{
				"status":  false,
				"message": "route tidak diijinkan untuk diakses",
			})
			c.AbortWithStatus(403)
		}
		// c.Next()
	}
}




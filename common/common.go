package common

import(
    "os"
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "time"
)

const(
    // PORT listen and serve on 0.0.0.0:8080
    PORT = ":8088" 
)

// GeoLocation store fields, Latitud & Longitude
type( 
	GeoLocation struct {
    	Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)

// CORSMiddleware is a middleware function
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost"+PORT)
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            fmt.Println("OPTIONS")
            c.AbortWithStatus(200)
        } else {
            c.Next()
        }
    }
}

// CheckError function is general error validator
func CheckError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
        os.Exit(1)
    }
}

func sendRequest() ([]byte, error) {
	resp, err := http.Get("https://api.ipdata.co?api-key=test")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// JSONToStruct Make a request to: https://api.ipdata.co?api-key=test, and returns your latitude and longitud
func JSONToStruct()(float64, float64, error) {
	var ipData GeoLocation

	jsonResponse, err := sendRequest()
	CheckError(err)
	
    if err := json.Unmarshal(jsonResponse, &ipData); err != nil {
        return 0, 0, err
    }
	return ipData.Latitude, ipData.Longitude, nil
}

func MakeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}
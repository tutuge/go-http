package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func queryEs() {
	url := "http://10.251.80.44:19200/_search"
	// 表单格式数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=王二小&amp;age=18"
	// json格式数据，一般使用结构体发送或解析json格式数据
	contentType := "application/json"
	data := `{
    "query": {
        "bool": {
            "must": [
                {
                    "range": {
                        "@timestamp": {
                            "gte": "now-5m",
                            "lte": "now"
                        }
                    }
                },
                {
                    "query_string": {
                        "default_field": "message",
                        "query": "error or exception"
                    }
                }
            ],
            "must_not": [
                {
                    "query_string": {
                        
                        "query": "kubernetes.namespace_labels.name:(znyz-prod) and kubernetes.namespace_labels.name:(znyz-test) and kubernetes.namespace_labels.name:(znyz-uat) and kubernetes.namespace_labels.name:(monitor) and kubernetes.namespace_labels.name:agilebi and kubernetes.deployment.name:datavisual-deployment and kubernetes.deployment.name:rcapi-deployment and kubernetes.deployment.name:jobserver-deployment and kubernetes.deployment.name:kubernetes-dashboard"
                    }
                }
            ]
        }
    },
    "_source": [
        "kubernetes.namespace_labels.name",
        "kubernetes.deployment.name",
        "message"
    ]
}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	var result Result
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
	}
	fmt.Printf("test\t", result)
	if result.Hits.Total.Value > 0 {
		dingUrl := "https://oapi.dingtalk.com/robot/send?access_token=26b05a49157794ce472fc54a00330e86cb60e65f6944131863fc4e2de7950cb9"
		var ding DingMarkDown
		ding.Msgtype = "markdown"
		ding.Markdown.Title = "bug来了"
		//具体的报错
		var build strings.Builder
		for _, value := range result.Hits.Hits {
			build.WriteString("## 服务：" + value.Source.Kubernetes.Deployment.Name + "\n")
			build.WriteString("### 环境：" + value.Source.Kubernetes.NamespaceLabels.Name + "\n")
			build.WriteString("具体报错：" + value.Source.Message + "\n")
		}
		ding.Markdown.Text = build.String()
		marshal, err := json.Marshal(ding)
		if err != nil {
			fmt.Print("json failed,err:%v\n", err)
		}
		dingResp, dingErr := http.Post(dingUrl, contentType, strings.NewReader(string(marshal)))
		dingRes, dingErr := io.ReadAll(dingResp.Body)
		fmt.Print("ding return", string(dingRes))
		if dingErr != nil {
			fmt.Print("ding get resp failed,err:%v\n", string(dingRes))
			return
		}
	}

}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//go func() {
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))))
	_, err := c.AddFunc("0 */5  * * * ?", func() {
		unix := time.Now().Unix()
		tm := time.Unix(unix, 0)
		fmt.Printf("time= %s \n", tm.Format("2006-01-02 03:04:05 PM"))
		//执行发送
		queryEs()
	})
	if err != nil {
		log.Printf("%s", err)
	}
	c.Start()
	//}()
	//
	//<-ctx.Done()
	//fmt.Println("bye")

	http.Handle("/", &CustomerHandler{})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

type CustomerHandler struct {
}

func (c *CustomerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("implement http server by self")
	_, err := writer.Write([]byte("server echo"))
	if err != nil {
		return
	}
}

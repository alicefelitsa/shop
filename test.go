package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"shop/config"
	"strings"
	"sync"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func main() {
	//movieInfo := checkMovie("http://v.felitsa.top:7000/movie/%E7%A0%B4%C2%B7%E5%9C%B0%E7%8B%B1-2024/cover.png")
	//fmt.Println(movieInfo)

	var wg sync.WaitGroup
	url := "http://v.felitsa.top:7000/getMovieData"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := make(map[string]interface{})
	_ = json.Unmarshal(body, &result)
	resData := result["data"].([]interface{})
	for _, v := range resData {
		wg.Add(1)
		go func(v interface{}) {
			defer wg.Done()
			var format string
			movieData := v.(map[string]interface{})
			fmt.Println("名称：", movieData["title"])
			fmt.Println("海报：", movieData["cover"])
			playUrl, episode := getPlayerLink(movieData["title"].(string))
			fmt.Println("地址：", playUrl)
			if strings.Contains(playUrl, "mp4") {
				format = "mp4"
			} else {
				format = "m3u8"
			}
			fmt.Println("格式：", format)
			fmt.Println("集数：", episode)

			movieInfo := checkMovie(movieData["cover"].(string))
			var name string
			if len(movieInfo) > 0 || movieInfo != nil {
				fmt.Println("AI返回数据：", movieInfo)
				fmt.Printf("片名: %v\n", movieInfo["片名"])
				fmt.Printf("导演: %v\n", movieInfo["导演"])
				fmt.Printf("年份: %v\n", movieInfo["年份"])
				fmt.Printf("国家: %v\n", movieInfo["国家"])
				fmt.Printf("分类: %v\n", movieInfo["分类"])
				fmt.Printf("评分: %v\n", movieInfo["评分"])
				fmt.Printf("简介: %v\n", movieInfo["简介"])
				category := fmt.Sprintf("%v", movieInfo["分类"])
				director := fmt.Sprintf("%v", movieInfo["导演"])
				name = movieInfo["片名"].(string)
				if name == "未知" {
					name = movieData["title"].(string)
				}
				if len(movieInfo) > 0 {
					ctime := time.Now().Format("2006-01-02 15:04:05")
					exec, err := config.Mysql.Exec(`insert into video
			(title, year, director, score, region, episode, plot, play_url, ctime, format, category) values (?,?,?,?,?,?,?,?,?,?,?)`,
						name, movieInfo["年份"], director, movieInfo["评分"], movieInfo["国家"], episode, movieInfo["简介"], "/movie/"+movieData["title"].(string)+"/", ctime, format, category)
					if err != nil {
						log.Println("写入失败：", err)
					}
					id, _ := exec.LastInsertId()
					if id > 0 {
						fmt.Println("写入成功：", id)
					}
				} else {
					log.Println("获取电影信息失败：", movieData["title"])
				}
			} else {
				name = movieData["title"].(string)
				if len(movieInfo) > 0 {
					ctime := time.Now().Format("2006-01-02 15:04:05")
					exec, err := config.Mysql.Exec(`insert into video
			(title, year, director, score, region, episode, plot, play_url, ctime, format, category) values (?,?,?,?,?,?,?,?,?,?,?)`,
						name, "未知", "未知", "未知", "未知", episode, "未知", "/movie/"+movieData["title"].(string)+"/", ctime, format, "未知")
					if err != nil {
						log.Println("写入失败：", err)
					}
					id, _ := exec.LastInsertId()
					if id > 0 {
						fmt.Println("写入成功：", id)
					}
				} else {
					log.Println("获取电影信息失败：", movieData["title"])
				}
			}
		}(v)
	}
	wg.Wait()
	log.Println("所有任务完成")
}

func getPlayerLink(title string) (string, int) {
	url := "http://v.felitsa.top:7000/apiPlayer?name=" + title + "&page=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", 0
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", 0
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", 0
	}
	result := make(map[string]interface{})
	_ = json.Unmarshal(body, &result)
	movieInfo := result["data"].(map[string]interface{})

	return movieInfo["playerLink"].(string), len(movieInfo["playerLinkDataAll"].([]interface{}))
}

// 请求AI获取电影相关信息
func checkMovie(imagUrl string) map[string]interface{} {
	// 创建 HTTP 客户端
	client := &http.Client{}
	// 构建请求体
	requestBody := RequestBody{
		// 模型列表：https://help.aliyun.com/model-studio/getting-started/models
		Model: "qwen3.6-plus",
		Messages: []Message{
			{
				Role:    "user",
				Content: "<image>" + imagUrl + "</image>从这张电影海报查找相关信息：片名、导演、年份、国家、分类、评分、简介。只返回JSON格式的简单数据，不要多余内容。",
			},
		},
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}
	// 创建 POST 请求
	// 以下是北京地域base_url，如果使用新加坡地域的模型，需要将base_url替换为：https://dashscope-intl.aliyuncs.com/compatible-mode/v1/chat/completions
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	// 设置请求头
	// 若没有配置环境变量，请用阿里云百炼API Key将下行替换为：apiKey := "sk-xxx"
	// 新加坡和北京地域的API Key不同。获取API Key：https://help.aliyun.com/model-studio/get-api-key
	apiKey := "sk-df683c3866a34814b5afbacc602faccc"
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("请求失败：", err)
		return nil
	}
	defer resp.Body.Close()
	// 读取响应体
	bodyText, _ := io.ReadAll(resp.Body)

	// 解析响应
	var resData map[string]interface{}
	if err := json.Unmarshal(bodyText, &resData); err != nil {
		log.Println("内容解析失败：", err)
		return nil
	}

	// 提取choices中的内容
	choices, ok := resData["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		log.Println("No choices in response")
		return nil
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		log.Println("Invalid choice format")
		return nil
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		log.Println("Invalid message format")
		return nil
	}

	content, ok := message["content"].(string)
	if !ok {
		log.Println("No content in message")
		return nil
	}

	// 打印原始内容
	//fmt.Println("AI返回的原始内容:")
	//fmt.Println(content)
	//fmt.Println("\n--- 解析后的电影信息 ---")

	// 尝试解析内容为JSON
	var movieInfo map[string]interface{}
	if err := json.Unmarshal([]byte(content), &movieInfo); err != nil {
		// 如果直接解析失败，尝试提取JSON部分
		//fmt.Println("尝试提取JSON内容...")
		// 查找JSON的开始和结束位置
		start := -1
		end := -1
		for i := 0; i < len(content); i++ {
			if content[i] == '{' && start == -1 {
				start = i
			}
			if content[i] == '}' {
				end = i + 1
			}
		}
		if start != -1 && end != -1 {
			jsonStr := content[start:end]
			if err := json.Unmarshal([]byte(jsonStr), &movieInfo); err != nil {
				log.Println("解析JSON失败: ", err)
				return nil
			}
		} else {
			log.Println("未找到JSON内容")
			return nil
		}
	}

	// 输出提取的电影信息
	//fmt.Printf("片名: %v\n", movieInfo["片名"])
	//fmt.Printf("导演: %v\n", movieInfo["导演"])
	//fmt.Printf("年份: %v\n", movieInfo["年份"])
	//fmt.Printf("国家: %v\n", movieInfo["国家"])
	//fmt.Printf("分类: %v\n", movieInfo["分类"])
	//fmt.Printf("评分: %v\n", movieInfo["评分"])
	//fmt.Printf("简介: %v\n", movieInfo["简介"])

	//fmt.Println(movieInfo)
	return movieInfo
}

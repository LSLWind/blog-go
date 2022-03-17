package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/**
文章同步与导出
*/
type ArticleUtil struct {
}

type Article struct {
	Id         uint
	Title      string
	Keyword    string
	Author     string
	Desc       string
	Content    string
	Numbers    int
	Img_url    string
	Type       int
	State      int
	Tags       string
	CategoryId uint
	Category   string
	Views      int
	Comments   int
	Likes      int
	CreateTime time.Time
	UpdateTime time.Time
}

//操作系统指定的路径分隔符
var fileSeparator = string(os.PathSeparator)
var globalDb *gorm.DB

//连接数据库
func init() {
	url := "root:123456@tcp(127.0.0.1:3306)/blog_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("db connection succedssed")
	}
	globalDb = db
}

/**
指定目录，同步所有文章，需要传递顶级目录
*/
func (a ArticleUtil) ArticleDirSync(dirPath string) {
	//以只读的方式打开目录
	dir, err := os.OpenFile(dirPath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err.Error())
	}
	//延迟关闭目录
	defer dir.Close()
	//读取目录信息
	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err.Error())
	}

	//开始同步，递归处理，目录名为标签
	var tags []string

	for _, info := range fileInfo {
		//判断是否是目录
		if info.IsDir() {
			//忽略.git等隐藏文件
			if info.Name()[0] == '.' {
				continue
			}
			tags = append(tags, info.Name())
			a.ArticleDirSyncRecursion(dirPath+fileSeparator+info.Name(), &tags)
			tags = append(tags[:len(tags)-1]) //删除最后一个元素，出栈
		} else {
			fmt.Println("文件：" + info.Name())
			InsertOneArticleByFilePath(dirPath+fileSeparator+info.Name(), tags)
		}
	}
}

/**
递归同步文章，必须为目录
*/
func (a ArticleUtil) ArticleDirSyncRecursion(dirPath string, tags *[]string) {
	//以只读的方式打开目录
	fileCount := 0
	dir, err := os.OpenFile(dirPath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err.Error())
	}
	//延迟关闭目录
	defer dir.Close()
	//读取目录信息
	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, info := range fileInfo {
		//判断是否是目录
		if info.IsDir() {
			*tags = append(*tags, info.Name())
			fmt.Println(tags)
			a.ArticleDirSyncRecursion(dirPath+fileSeparator+info.Name(), tags)
			*tags = append((*tags)[:len((*tags))-1]) //删除最后一个元素，出栈
		} else {
			//插入文件，只接受.md
			name := info.Name()
			nameL := len(name)
			if name[nameL-3:nameL] == ".md" {
				fmt.Println("文件：" + info.Name())
				InsertOneArticleByFilePath(dirPath+fileSeparator+info.Name(), *tags)
				fileCount++
			}
		}
	}
}

func InsertOneArticleByFilePath(filePath string, tags []string) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeDir)
	article := Article{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	content := string(data)
	category := filePath[:len(filePath)-len(file.Name())-2]
	fmt.Println(category)

	article.Title = file.Name()
	article.Author = "梁松林"
	article.Content = content
	article.Numbers = len(content)
	article.Category = category
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()
	globalDb.Table("articles").Create(&article)

}

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
	Id         uint      `gorm:"column:id"`
	Title      string    `gorm:"column:title"`
	Keyword    string    `gorm:"column:keyword"`
	Author     string    `gorm:"column:author"`
	Desc       string    `gorm:"column:desc"`
	Content    string    `gorm:"column:content"`
	Numbers    int       `gorm:"column:numbers"`
	ImgUrl     string    `gorm:"column:img_url"`
	Type       int       `gorm:"column:type"`
	State      int       `gorm:"column:state"`
	Tags       string    `gorm:"column:tags"`
	CategoryId uint      `gorm:"column:category_id"`
	Category   string    `gorm:"column:category"`
	Views      int       `gorm:"column:views"`
	Comments   int       `gorm:"column:comments"`
	Likes      int       `gorm:"column:likes"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
type Category struct {
	Id              uint
	Name            string
	SubCategoryId   int
	SubCategoryName string
	CreateTime      time.Time
	UpdateTime      time.Time
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
同时插入目录信息，最多支持二级目录，结构为：目录A-目录B(文章）/文章
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
			SimpleInsertCategory(info.Name(), (*tags)[len(*tags)-1]) //插入目录
			*tags = append((*tags)[:len((*tags))-1])                 //删除最后一个元素，出栈
		} else {
			//插入文件，只接受.md
			name := info.Name()
			nameL := len(name)
			if name[nameL-3:nameL] == ".md" {
				fmt.Println("文件：" + info.Name())
				InsertOneArticleByFilePath(dirPath+fileSeparator+info.Name(), *tags)
				SimpleInsertCategory(info.Name(), (*tags)[len(*tags)-1]) //插入目录
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
	fileAbsName := file.Name()
	var category string
	var title string
	for i := len(fileAbsName) - 1; i >= 0; i-- {
		if string(fileAbsName[i]) == fileSeparator {
			category = fileAbsName[:i]
			title = fileAbsName[i+1:]
			break
		}
	}

	fmt.Println(category)

	article.Title = title
	article.Author = "梁松林"

	//注意，中文utf8是3字节，英文是1字节
	utfDesc := []rune(content)
	if len(utfDesc) > 200 {
		article.Desc = string(utfDesc[:200])
	} else {
		article.Desc = string(utfDesc)
	}

	article.Content = content
	article.Numbers = len(content) / 2
	article.Category = category
	article.CreateTime = time.Now()
	article.UpdateTime = time.Now()
	globalDb.Table("articles").Create(&article)

}

/**
插入目录，维护二级目录结构
*/
func SimpleInsertCategory(name string, parentName string) {
	if name == parentName {
		return
	}

	category := Category{}
	category.Name = parentName
	category.SubCategoryName = name
	category.CreateTime = time.Now()
	category.UpdateTime = time.Now()
	globalDb.Table("category").Create(&category)
}

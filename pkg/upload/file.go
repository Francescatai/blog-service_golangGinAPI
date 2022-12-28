package upload

import(
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"go_gin_blog/global"
	"go_gin_blog/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

// 取得文件名稱，通過取得文件後綴並篩出原始文件名進行MD5加密，最後返回經過加密處理後的文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// 取得文件後綴(.jpg/.png...)，主要透過path.Ext方法進行循環查找"."符號，通過切片索引返回對應的文件後綴名稱
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 取得文件保存地址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := io.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// 保存所上傳的文件，主要通過調用os.Create方法創建項目地址的文件，再通過file.Open方法打開源地址的文件，結合io.Copy方法實現兩者之間的文件內容copy
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
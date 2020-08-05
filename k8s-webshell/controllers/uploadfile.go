// package controllers
// import (
//     "path"
//     "crypto/md5"
//     "time"
//     "fmt"
//     "github.com/astaxie/beego"
//     "math/rand"
//     "os"
// )
// type UploadController struct{
//     beego.Controller 
// }

// func (this *UploadController) UpForm(){
//     this.TplName = "upload.tpl"
// }

// func (this *UploadController) UpFile(){

//     f, h, _ := this.GetFile("myfile")//获取上传的文件
//     ext := path.Ext(h.Filename)
//     //验证后缀名是否符合要求
//     var AllowExtMap map[string]bool = map[string]bool{
//         ".jpg":true,
//         ".jpeg":true,
//         ".png":true,
//     }
//     if _,ok:=AllowExtMap[ext];!ok{
//         this.Ctx.WriteString( "后缀名不符合上传要求" )
//         return 
//     }
//     //创建目录
//     uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
//     err := os.MkdirAll( uploadDir , 777)
//     if err != nil {
//         this.Ctx.WriteString( fmt.Sprintf("%v",err) )
//         return 
//     }
//     //构造文件名称
//     rand.Seed(time.Now().UnixNano())
//     randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
//     hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )

//     fileName := fmt.Sprintf("%x",hashName) + ext
//     //this.Ctx.WriteString(  fileName )

//     fpath := uploadDir + fileName
//     defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
//     err = this.SaveToFile("myfile", fpath)
//     if err != nil {
//         this.Ctx.WriteString( fmt.Sprintf("%v",err) )
//     }
//     this.Ctx.WriteString( "上传成功~！！！！！！！" )
// }

package controllers
 
import (
    "fmt"
    "strings"
    "path"
    "github.com/astaxie/beego"
)
 
type FileOptUploadController struct {
    beego.Controller
}
 
//上传下载文件的页面
func (c *FileOptUploadController) Get() {
    c.TplName = "fileopt.html"
}
 
//上传文件
func (this *FileOptUploadController) Post() {
    //image，这是一个key值，对应的是html中input type-‘file’的name属性值
    f, h, _ := this.GetFile("file")
    //得到文件的名称
    fileName := h.Filename 
    arr := strings.Split(fileName, ":")
    if len(arr) > 1 {   
      index := len(arr) - 1
      fileName = arr[index]
    }
    fmt.Println("文件名称:")
    fmt.Println(fileName)
    //关闭上传的文件，不然的话会出现临时文件不能清除的情况
    f.Close()  
    //保存文件到指定的位置
    //static/uploadfile,这个是文件的地址，第一个static前面不要有/
    this.SaveToFile("file", path.Join("k8s-config","config"))
    //显示在本页面，不做跳转操作 
    this.TplName = "fileopt.html" 
}
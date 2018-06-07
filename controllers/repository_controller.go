package controllers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "proj/models"
	"crypto/sha1"
    "os"
	"io"
    "io/ioutil"
	"fmt"
    "time"
    "encoding/base64" 
    //"encoding/json"
)

type NoteApiController struct {
	beego.Controller
}

type ImgUploadController struct {
	beego.Controller
}

type ImgUploadRespond struct {
    State bool `json:"state"`
    Sha1 string `json:"sha1"`
    Msg string `json:"msg"`
    Url string `json:"url"`
}

type NoteApiRespond struct {
    State bool `json:"state"`
    Id int64 `json:"id"`
    Msg string `json:"msg"` 
}

// func (c *NoteApiController) Get() {
//     var res NoteApiRespond

//     o := orm.NewOrm()
//     o.Using("default")
//     op := c.GetString("op")
//     switch op {
//     case "list":
//         c.Ctx.WriteString(json.Marshal(map[string]interface{}{
//             "a": this.a,
//             "b": this.b,
//             "c": this.c,
//             "d": this.d,
//         }))
//     case "del":
//     }
// }

func (c *NoteApiController) Post() {
    var res NoteApiRespond
    defer func() {
        c.Data["json"] = res
        c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
        c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST")
        c.ServeJSON()  
    }()

    o := orm.NewOrm()
    o.Using("default")

    op := c.GetString("op")
    switch op {
    case "add":
        var rec models.Records;
        content := c.GetString("content")
        img := c.GetString("img")
        if (content == "" && img == "") {
            res.State = false;
            res.Msg = "invalid arguments";
            return;
        }

        rec.Content = content;
        rec.Images = img;
        rec.Time = time.Now();
        id, err := o.Insert(&rec)

        if (err != nil) {
            res.State = false;
            res.Msg = err.Error();
        } else {
            res.State = true;
            res.Id = id;
            res.Msg = "";
        }


    case "del":
    }
}

func (c *ImgUploadController) Post() {
    var res ImgUploadRespond

    defer func() {
        c.Data["json"] = res
        c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
        c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST")
        c.ServeJSON()  
    }()

    img_path := beego.AppConfig.String("img_save_path")
    url_header := beego.AppConfig.String("img_url_header")

    is_base64, _ := c.GetBool("isbase64")
    if (!is_base64) {
        file, header, err := c.GetFile("img")
        if (err != nil) {
        	c.Ctx.WriteString(err.Error())
        } else {
            defer file.Close()
            h := sha1.New()
            _, erro := io.Copy(h, file)
            if erro != nil {
                res.State = false;
                res.Msg = erro.Error();
            }

            sha1_str := fmt.Sprintf("%x", h.Sum(nil))
            file_name := sha1_str
            switch (header.Header["Content-Type"][0]) {
            case "image/jpeg": 
                file_name += ".jpg";
            case "image/png":
                file_name += ".png";
            case "image/gif":
                file_name += ".gif"
            default:
                res.State = false;
                res.Msg = "invalid Content-Type";
                return;
            }
            file_path := img_path + file_name

            // check if exist
            _, err := os.Stat(file_path)
            if (err != nil) {
                if (os.IsNotExist(err)) {
                    c.SaveToFile("img", file_path)
                } else {
                    res.State = false
                    res.Msg = err.Error()
                    res.Sha1 = ""
                    return;
                }       
            }

            res.State = true;
            res.Sha1 = sha1_str;
            res.Url = url_header + file_name
        }

    } else {
        data := c.GetString("img")
        if (data != "") {
            var ext string
            if (data[0:21] == "data:image/png;base64") {
                ext = ".png"
            } else if (data[0:21] == "data:image/gif;base64") {
                ext = ".gif"
            } else if (data[0:21] == "data:image/jpg;base64") {
                ext = ".jpg"
            } else  {
                res.State = false
                res.Msg = "invalid img"
            }
	    fmt.Println(data[22:102])
            buffer, err := base64.StdEncoding.DecodeString(data[22:]) // remove "data:image/gif;base64,"
            if (err != nil) {
                res.State = false
                res.Msg = err.Error()
            } else {
                h := sha1.New()
                h.Write(buffer)
                sha1_str := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(sha1_str)     
                file_name := sha1_str + ext
                file_path := img_path + file_name
                res.State = true
                res.Sha1 = sha1_str
                res.Url = url_header + file_name
                // check if exist
                _, err := os.Stat(file_path)
                if (err != nil) {
                    if (os.IsNotExist(err)) {
                        ioutil.WriteFile(file_path, buffer, os.ModeAppend)
                    } else {
                        res.State = false
                        res.Msg = err.Error()
                        res.Sha1 = ""
                    }
                    
                }
                
            }
            
        }
    }


}

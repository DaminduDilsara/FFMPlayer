package api

import (
	"AdvancedNetwork/pkg/dbOps"
	"AdvancedNetwork/pkg/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func UploadFileApi(c echo.Context) (string, error) {
	// Read form fields
	title := c.FormValue("title")
	description := c.FormValue("description")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("video")
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("uploads/" + file.Filename)

	filename := file.Filename
	fmt.Println("file name : ", filename)

	if err != nil {
		return "", err
	}
	defer dst.Close()

	newVideo := models.Video{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Description: description,
	}

	idObj, err := dbOps.DB_Create_Video(newVideo)
	if err != nil {
		log.Fatal(err)
	}
	fileId := idObj.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println("created video id : ", fileId)

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	aspectSet := "ffmpeg -y -i ./uploads/" + filename + " -aspect 16:9 -c copy ./uploads/processed/" + filename
	aspectArgs := strings.Split(aspectSet, " ")
	aspectSetCmd := exec.Command(aspectArgs[0], aspectArgs[1:]...)
	b, err := aspectSetCmd.CombinedOutput()
	if err != nil {
		log.Printf("Running aspectSetCmd failed: %v", err)
	}
	fmt.Printf("%s\n", b)

	getThumbnail := "ffmpeg -i ./uploads/" + filename + " -vframes 1 ./segments/thumbnails/" + fileId + ".jpg"

	thumbArgs := strings.Split(getThumbnail, " ")
	getThumbnailCmd := exec.Command(thumbArgs[0], thumbArgs[1:]...)
	fmt.Println("---iuadiufysgw efuygwe------", getThumbnailCmd.String())

	b, err = getThumbnailCmd.CombinedOutput()
	if err != nil {
		log.Printf("Running getThumbnailCmd failed: %v", err)
	}
	fmt.Printf("%s\n", b)

	err = os.Mkdir("segments/"+fileId, os.ModePerm)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("./segmenting", filename, fileId)
	fmt.Println(")))))))))))))))))))))))))))))", cmd.String())
	b, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Running getThumbnailCmd failed: %v", err)
	}
	fmt.Printf("%s\n", b)

	//	segment := `ffmpeg -re -i ./uploads/processed/` + filename + ` -map 0 -map 0 -map 0 -c:a aac -c:v libx264 -b:v:1 20000k -b:v:2 20000k -b:v:2 20000k -s:v:0 1920x1080 -s:v:1 1280x720 -s:v:2 720x480 -profile:v:1 baseline -profile:v:2 baseline -profile:v:0 main
	//-bf 1 -keyint_min 120 -g 120 -sc_threshold 0 -b_strategy 0 -ar:a:1 22050 -use_timeline 1 -use_template 1 -adaptation_sets "id=0,streams=v id=1,streams=a" -f dash ./segments/` + fileId + `/` + fileId + `_out.mpd`
	//	segmentArgs := strings.Split(segment, " ")
	//segmentCmd := exec.Command("ffmpeg", "-re", "-i", "./uploads/processed/test.mp4", "-map", "0", "-map", "0", "-c:a", "aac",
	//	"-c:v", "libx264", "-b:v:0", "20000k", "-b:v:1", "20000k", "-s:v:1", "1280x720",
	//	"-profile:v:1", "baseline", "-profile:v:0", "main", "-bf", "1", "-keyint_min", "120", "-g", "120", "-sc_threshold", "0", "-b_strategy",
	//	"0", "-ar:a:1", "22050", "-use_timeline", "1", "-use_template", "1", "-window_size", "5", "-adaptation_sets", `"id=0,streams=v id=1,streams=a"`, "-f", "dash", "redda_out.mp4")
	//fmt.Println("---------", segmentCmd.String())
	//
	//b, err = segmentCmd.CombinedOutput()
	//if err != nil {
	//	log.Printf("Running segmentCmd failed: %v", err)
	//}
	//fmt.Printf("================================================================== %s\n ==========================================", b)

	return "", c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields title=%s and description=%s.</p>", file.Filename, title, description))
}

package usecase

import (
	"cleanArchCrud/models"
	"context"
	"fmt"
	"strings"
	"time"
	"strconv"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/spf13/viper"
)


func (g *GeneralUsecase) UploadFile(ctx context.Context, file models.FileModel) (models.FileResponse, error) {
	var result models.FileResponse
	if file.Category == "" {
		file.Category = "All"
	}
	result.Filename = file.File.Filename
	var imgExt = strings.Split(result.Filename, ".")[len(strings.Split(result.Filename, ".")) - 1]
	var newFileName string
	projectName := viper.GetString("cloudinary.projectName")
	if projectName == "" {
		projectName = "random"
	}
	username := fmt.Sprintf("%v",ctx.Value("payload_id"))
	files, err := file.File.Open()
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	if imgExt == "jpg" || imgExt == "png" || imgExt == "jpeg" || imgExt == "pdf" || imgExt == "mp4" {
		newFileName = "file_" + strconv.Itoa(int(time.Now().Unix()))
	} else {
		newFileName = "file_" + strconv.Itoa(int(time.Now().Unix())) + "." + imgExt
	}

	username = strings.ReplaceAll(username, " ", "_")
	file.Category = strings.ReplaceAll(file.Category, " ", "_")
	newFileName = strings.ReplaceAll(newFileName, " ", "_")
	cld, _ := cloudinary.NewFromParams(viper.GetString("cloudinary.user"), viper.GetString("cloudinary.key"), viper.GetString("cloudinary.secret"))
	uploadResult, err := cld.Upload.Upload(
		ctx,
		files,
		uploader.UploadParams{PublicID: fmt.Sprintf(`%s/%s/%s/%s`, projectName, username, file.Category, newFileName), ResourceType: "raw"})
	if err != nil {
		fmt.Println(err)
	}

	result.Fileurl = uploadResult.SecureURL
	return result, nil
}
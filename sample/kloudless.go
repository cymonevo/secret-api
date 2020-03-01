package sample

import (
	"context"

	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/provider"
	"github.com/cymon1997/go-backend/sdk/kloudless"
)

func KloudlessSample() {
	client := provider.GetKloudlessClient()
	folder(client)
	file(client)
}

func folder(client kloudless.Client) {
	folder, err := client.CreateFolder(context.Background(), kloudless.CreateFolderRequest{
		DestID: "root",
		Name:   "sandbox",
	})
	if err != nil {
		return
	}
	log.Infof("CreateFolder", "success create folder %+v", folder)
	contents, err := client.GetFolderContents(context.Background(), kloudless.GetFolderContentsRequest{
		FolderID: folder.ID,
	})
	if err != nil {
		return
	}
	log.Infof("CreateFolder", "success get folder contents %+v", contents)
}

func file(client kloudless.Client) {
	file, err := client.DownloadFile(context.Background(), kloudless.DownloadFileRequest{
		FileID: "FviBGig8hJzj7TVEqIzjP459zeLLKpy9Gl5eoM4v3mCOOCpeuHZv--892QAeS1Yh0",
	})
	if err != nil {
		return
	}
	log.Infof("DownloadFile", "success download file")

	upload, err := client.UploadFile(context.Background(), kloudless.UploadFileRequest{
		DestID:  "root",
		Name:    "sample.jpg",
		RawFile: file.RawFile,
	})
	if err != nil {
		return
	}
	log.Infof("UploadFile", "success upload file %+v", upload)
}

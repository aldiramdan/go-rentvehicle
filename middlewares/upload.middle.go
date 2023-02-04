package middlewares

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aldiramdan/go-backend/libs"
)

func AuthUploadFile() Middle {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			file, fileHeader, err := r.FormFile("picture")
			if err != nil {
				if err == http.ErrMissingFile {
					imgName := "default_image.jpg"
					ctx := context.WithValue(r.Context(), "imageName", "public/"+imgName)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				libs.GetResponse(err.Error(), 401, true).Send(w)
				return
			}

			defer file.Close()

			buff := make([]byte, 512)
			_, err = file.Read(buff)
			if err != nil {
				libs.GetResponse(err.Error(), 500, true).Send(w)
				return
			}

			filetype := http.DetectContentType(buff)
			if filetype != "image/jpeg" && filetype != "image/png" {
				libs.GetResponse(err.Error(), 401, true).Send(w)
				return
			}

			_, err = file.Seek(0, io.SeekStart)
			if err != nil {
				libs.GetResponse(err.Error(), 500, true).Send(w)
				return
			}

			err = os.MkdirAll("./public", os.ModePerm)
			if err != nil {
				libs.GetResponse(err.Error(), 401, true).Send(w)
				return
			}

			imgName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
			pathRes := filepath.Join("./public", imgName)
			dst, err := os.Create(pathRes)
			if err != nil {
				_ = os.Remove(pathRes)
				libs.GetResponse(err.Error(), 401, true).Send(w)

				return
			}

			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				_ = os.Remove(pathRes)
				libs.GetResponse("error copy filesystem", 401, true).Send(w)
				return
			}

			ctx := context.WithValue(r.Context(), "imageName", "public/"+imgName)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

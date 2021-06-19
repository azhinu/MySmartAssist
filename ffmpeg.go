package main
import (
  "github.com/ssttevee/go-ffmpeg"
  "net/http"
  "context"
  "log"
)

func ConvertOpus(FileURL string){
  cfg, err := ffmpeg.DefaultConfiguration()
  check(err)
  job := cfg.NewJob(ffmpeg.Flag("-y"))
  resp, err := http.Get(FileURL)
  check(err)
  defer resp.Body.Close()
  job.AddInputReader(resp.Body)
  job.AddOutputFile("/tmp/out.mp3")
  _, statusChan, _ := job.Start(context.Background())
  for status := range statusChan {
      switch v := status.(type) {
      case *ffmpeg.Progress:
          log.Printf("%#v", v)
      case *ffmpeg.Done:
          log.Printf("done")
          return
      case *ffmpeg.Error:
          log.Fatalf(v.Error())
      }
  }
}

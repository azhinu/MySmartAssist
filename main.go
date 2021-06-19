package main
import (
  "context"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "math/rand"
  "sync"
  "time"
  "regexp"
  "github.com/Code-Hex/Neo-cowsay"
  "github.com/ssttevee/go-ffmpeg"
  "github.com/wit-ai/wit-go"
  "github.com/yanzay/tbot"
  // speech "cloud.google.com/go/speech/apiv1"
  // speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"

)


func main() {
  // Load config. getconfig.go
  config := SetEnv()
  // Define ehmmm.. This stuff. mb it's a methods? idk.
  bot := tbot.New(config.Token,
    tbot.WithWebhook(config.URL, config.Listener))
	client := bot.Client()
  wit := witai.NewClient(config.WitToken)


  // Handle /start command
	bot.HandleMessage("^/start", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, say("Hi. I'm your smart assistant. You can ask me anything."))
	})
  //Handle /help command
  bot.HandleMessage("^/help", func(m *tbot.Message) {
    client.SendMessage(m.Chat.ID, say("Just ask me. It's easy, trust me.\nSay <Hey, Suchka!>"))
  })
  // Handle anything other
  bot.HandleMessage(".+", func(m *tbot.Message) {
    client.SendMessage(m.Chat.ID, say("Please, try to use you mouth to ask me."))
  })
  bot.HandleMessage("", func(m *tbot.Message) {
    // Handle voice
		if m.Voice != nil {
			file, err := client.GetFile(m.Voice.FileID)
			check(err)

      // Wit STT
      VoiceReply, err := SpeechToText(wit, GetVoice(client.FileURL(file),"/tmp/"+m.Voice.FileID+".ogg","/tmp/out.mp3"))
      check(err)
      regex, err := regexp.MatchString("hey \\w+", VoiceReply)
      check(err)
      if regex {
        if chance() {
          client.SendMessage(m.Chat.ID, GetArt("love"))
        } else {
          client.SendMessage(m.Chat.ID, say("I glad hear U 2"))
        }
      } else {
        client.SendMessage(m.Chat.ID, say("Sorry but I didn't understand you..."))
        client.SendMessage(m.Chat.ID, GetArt("confused"))
      }
		} else {
      client.SendMessage(m.Chat.ID, say("I not properly sure what you mean."))
    }
	})

  log.Println("Bot started\nListen: ", config.Listener, "\nWebhook URL: ", config.URL, "\nToken: ", config.Token)
	log.Fatal(bot.Start())

}

func say(text string) string {
	result, err := cowsay.Say(
		cowsay.Phrase(text),
		cowsay.Type("hellokitty"),
	)
	check(err)
	return result
}
func GetVoice(FileURL, InputF, OutputF string)  (io.Reader){
  resp, err := http.Get(FileURL)
  check(err)
  defer resp.Body.Close()
  out, err := os.Create(InputF)
  check(err)
  defer out.Close()
  io.Copy(out, resp.Body)
  cfg, err := ffmpeg.DefaultConfiguration()
  check(err)
  job := cfg.NewJob(ffmpeg.Flag("-y"))
  job.AddInputFile(InputF)
  job.AddOutputFile(OutputF)
  _, statusChan, _ := job.Start(context.Background())
  for status := range statusChan {
      switch v := status.(type) {
      case *ffmpeg.Progress:
          log.Printf("%#v", v)
      case *ffmpeg.Done:
          log.Printf("done")
      case *ffmpeg.Error:
          log.Fatalf(v.Error())
      }
  }
  file, err := os.Open(OutputF)
  check(err)
  return file
}
func SpeechToText(wit *witai.Client, file io.Reader) (text string, err error) {

  check(err)
  msg, err := wit.Speech(&witai.MessageRequest{
    Speech: &witai.Speech{
      File: file,
      ContentType: "audio/mpeg3",
    },
  })
  log.Printf("Somebody says :"+msg.Text)

	return msg.Text, err
}
func GetArt(art string) (string){
  content, err := ioutil.ReadFile("art/"+art)
  check(err)
  result := string(content)
  return result
}
var onlyOnce sync.Once
func chance() bool {
  onlyOnce.Do(func() {
    rand.Seed(time.Now().UnixNano()) // only run once
  })
  return rand.Intn(10)/2>2
}
func check(err error) {
    if err != nil {
        log.Println(err)
        panic(err)
    }
}

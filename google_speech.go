package main
//Unfortunately it's requred CC to use.

// import (
        // "log"
        // "os"
        // speech "cloud.google.com/go/speech/apiv1"
        // speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
// func VoiceRecognize(voiceURL string){
//   ctx := context.Background()
//
//   // Creates a client.
//   client, err := speech.NewClient(ctx)
//   if err != nil {
//           log.Fatalf("Failed to create client: %v", err)
//   }
//   defer client.Close()
//   // Detects speech in the audio file.
//         resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
//                 Config: &speechpb.RecognitionConfig{
//                         Encoding:        speechpb.RecognitionConfig_LINEAR16,
//                         SampleRateHertz: 16000,
//                         LanguageCode:    "en-US",
//                 },
//                 Audio: &speechpb.RecognitionAudio{
//                         AudioSource: &speechpb.RecognitionAudio_Uri{Uri: voiceURL},
//                 },
//         })
//         if err != nil {
//                 log.Fatalf("failed to recognize: %v", err)
//         }
//
//         // Prints the results.
//         for _, result := range resp.Results {
//                 for _, alt := range result.Alternatives {
//                         log.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
//                 }
//         }
//
// }

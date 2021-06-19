package main
// For some reason it's not working.




// import (
//         "gopkg.in/hraban/opus.v2"
//         "bytes"
//         "encoding/binary"
// )
// func OpusDecode(voiceData []byte, sampleRate, channels int) (*bytes.Reader, error){
//   dec, err := opus.NewDecoder(sampleRate, channels)
//   frameSizeMs := 60  // if you don't know, go with 60 ms.
//   frameSize := channels * frameSizeMs * sampleRate / 1000
//   pcm := make([]int16, int(frameSize))
//   n, err := dec.Decode(voiceData, pcm)
//   pcm = pcm[:n*channels]
//   buf := new(bytes.Buffer)
//   binary.Write(buf, binary.LittleEndian, pcm)
//   result := bytes.NewReader(buf.Bytes())
//   return result, err
// }

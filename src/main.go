package main

import (
  "fmt"
  "log"
  "os/exec"
)

const (
  AUDIOIN = "audio-in/cogdiss_full_05312024.wav"
  OUTNAME = "audio-out/dawn_sky_06012024.wav"
  TSTART = "00:37:36"
  TEND = "00:41:00"
)

func main() {
  ffmpegCmd := exec.Command(
    "ffmpeg",
    "-i", AUDIOIN,
    "-ss", TSTART,
    "-to", TEND,
    "-c", "copy",
    OUTNAME,
  )
  ffmpegOutput, err := ffmpegCmd.CombinedOutput()
  if err != nil {
    log.Fatalf("\nmain(): An error occured while running ffmpegCmd: %s\n%v", string(ffmpegOutput), err)
  }
  fmt.Printf("\nmain(): Successfully sliced audio file %s and saved as %s with ffmpeg\n\n%s", AUDIOIN, OUTNAME, string(ffmpegOutput))
}

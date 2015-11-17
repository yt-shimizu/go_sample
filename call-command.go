package main

import (
  "log"
  "os/exec"
)

func main() {
  out, error := exec.Command("Date").Output()

  if error != nil {
    log.Printf("エラー: %s", error)
    return
  }

  log.Printf("現在時刻: %s", out)
}

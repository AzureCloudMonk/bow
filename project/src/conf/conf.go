package conf

import (
  "os"
)

var Env map[string] string

func Init()  {
  Env = make (map[string] string)

  var val string
  if val = os.Getenv("BS_LOG_SILENT"); val != "" {
      Env["log_silent"] = val
    } else {
      Env["log_silent"] = "no"
  }
  if val = os.Getenv("BS_SERVE_ADD"); val != "" {
      Env["servadd"] = val
    } else {
      Env["servadd"] = ":19808"
  }
  if val = os.Getenv("BS_DB_PATH"); val != "" {
      Env["dbpath"] = val
    } else {
      Env["dbpath"] = "/var/lib/bow"
  }
  os.Mkdir(Env["dbpath"], 0600)
  if val = os.Getenv("BS_DB_NAME"); val != "" {
      Env["dbname"] = val
    } else {
      Env["dbname"] = "asapdrf"
  }
}

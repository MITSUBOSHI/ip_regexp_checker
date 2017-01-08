package main

import(
   "fmt"
   "strconv"
   "os"
   "regexp"
 )

func main (){
  var target_regexp = os.Args[1]
  var target_ip_parts = os.Args[2]
  until_first_oct := regexp.MustCompile(`^\d{1,3}(\.\d{1,3}){0}$`)
  until_second_oct := regexp.MustCompile(`^\d{1,3}(\.\d{1,3}){1}$`)
  until_third_oct := regexp.MustCompile(`^\d{1,3}(\.\d{1,3}){2}$`)
  match_ip_regexp := regexp.MustCompile(target_regexp)
  if until_first_oct.MatchString(target_ip_parts){
    for second_oct :=0; second_oct < 256; second_oct++{
      for third_oct :=0; third_oct < 256; third_oct++{
        for fourth_oct :=0; fourth_oct < 256; fourth_oct++{
          var second_oct string = strconv.Itoa(second_oct)
          var third_oct string = strconv.Itoa(third_oct)
          var fourth_oct string = strconv.Itoa(fourth_oct)
          var target_ip_addresses string = target_ip_parts+"."+second_oct+"."+third_oct+"."+fourth_oct
          if match_ip_regexp.MatchString(target_ip_addresses){
            fmt.Println(target_ip_addresses)
          }
        }
      }
    }
  }else if until_second_oct.MatchString(target_ip_parts){
    for third_oct :=0; third_oct < 256; third_oct++{
      for fourth_oct :=0; fourth_oct < 256; fourth_oct++{
        var third_oct string = strconv.Itoa(third_oct)
        var fourth_oct string = strconv.Itoa(fourth_oct)
        var target_ip_addresses string = target_ip_parts+"."+third_oct+"."+fourth_oct
        if match_ip_regexp.MatchString(target_ip_addresses){
          fmt.Println(target_ip_addresses)
        }
      }
    }
  }else if until_third_oct.MatchString(target_ip_parts){
    for fourth_oct :=0; fourth_oct < 256; fourth_oct++{
      var fourth_oct string = strconv.Itoa(fourth_oct)
      var target_ip_addresses string = target_ip_parts+"."+fourth_oct
      if match_ip_regexp.MatchString(target_ip_addresses){
        fmt.Println(target_ip_addresses)
      }
    }
  }else{
    var target_ip_addresses string = target_ip_parts
    if match_ip_regexp.MatchString(target_ip_addresses){
      fmt.Println(target_ip_addresses)
    }
  }
}

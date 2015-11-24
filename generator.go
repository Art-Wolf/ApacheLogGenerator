package main

import (
   "fmt"
   "math/rand"
   "strconv"
   "time"
)

func randInt(min int, max int) string {
   return strconv.Itoa(min + rand.Intn(max-min))
}

func ipAddress() string {
   var ipStr string;

   ipStr = randInt(1, 254) + ".";
   ipStr += randInt(1, 254) +  ".";
   ipStr += randInt(1, 254) + ".";
   ipStr += randInt(1, 254);

   return ipStr;
}

func timeStamp() string {
   return time.Now().Format("02/Jan/2006:15:04:05 -0700");
}

func requestType() string {
   s := []string{"GET", "POST", "PUT", "DELETE"};
   return s[rand.Intn(len(s))];
}

func docType() string {
   s:= []string{".html", ".php", ".gif", ".jpeg", ".png"};
   return s[rand.Intn(len(s))];
}

func returnCode() string {
   s:= []string{"200", "301", "403", "404", "500"};
   return s[rand.Intn(len(s))];
}

func RandomString(strlen int) string {
   rand.Seed(time.Now().UTC().UnixNano());
   const chars = "abcdefghijklmnopqrstuvwxyz0123456789";
   result := make([]byte, strlen);
   
   for i := 0; i < strlen; i++ {
      result[i] = chars[rand.Intn(len(chars))];
   }
   
   return string(result);
}

func main() {
   rand.Seed( time.Now().UTC().UnixNano());

   for {
      time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond);
      fmt.Println(ipAddress() + " - - [" + timeStamp() + "] \"" + requestType() + " /" + RandomString(5) + "/" + RandomString(10) + "/" + RandomString(10) + docType() + " HTTP/1.0\" " + returnCode() + " " + randInt(20, 5000));
   }
}

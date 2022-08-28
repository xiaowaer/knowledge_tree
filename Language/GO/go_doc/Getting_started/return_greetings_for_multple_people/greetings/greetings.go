//声明包
package greetings

//导入本地模块
import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
// 和每个给定名字的人打招呼

func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    // 没有给定的名字输入，就返回一条错误消息
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    // 生成一条随机选定的切片中的消息
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
//给一群人打招呼
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    //用map 保存人名和打招呼对应的用语Hello
    messages := make(map[string]string)
    //消息体
    // Loop through the received slice of names, calling
    //根据切片做循环，做取出值，不取索引 
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        //map 的值放入hello返回的打招呼用语
        messages[name] = message
    }
    return messages, nil
}

// Init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}

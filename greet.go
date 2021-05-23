import "fmt"
import "errors"

func Hi(name string,lang string ) (string,error){

    switch lang{
    case "en":
        return fmt.Sprintf("hi %s!",name),nil
    case "ch":
        return fmt.Sprintf("nihao %s!",,name) ,nil
    default:
        return "", errors.New("unknown lang")
    }
}

package swift_test

import (
	"fmt"
	"github.com/ncw/swift"
	"os"
	"testing"
)

var c swift.Connection

func init() {
	UserName := os.Getenv("RCLOUD_API_USER")
	ApiKey := os.Getenv("RCLOUD_API_KEY")
	AuthUrl := os.Getenv("RCLOUD_AUTH_URL")
	if UserName == "" || ApiKey == "" || AuthUrl == "" {
		panic("RCLOUD_API_USER, RCLOUD_API_KEY and RCLOUD_AUTH_URL not all set")
	}
	c = swift.Connection{
		UserName: UserName,
		ApiKey:   ApiKey,
		AuthUrl:  AuthUrl,
	}
	err := c.Authenticate()
	if err != nil {
		panic(err)
	}
	fmt.Println("Authenticated")
}

func TestMain(t *testing.T) {
	fmt.Println(c)
	containers, err := c.ListContainers(nil)
	fmt.Println(containers, err)
	containerinfos, err2 := c.ListContainersInfo(nil)
	fmt.Println(containerinfos, err2)

	objects, err3 := c.ListObjects("SquirrelSave", nil)
	fmt.Println(objects, err3)
	objectsinfo, err4 := c.ListObjectsInfo("SquirrelSave", nil)
	fmt.Println(objectsinfo, err4)
	objects, err3 = c.ListObjects("SquirrelSave", &swift.ListObjectsOpts{Delimiter: '/', Path: ""})
	fmt.Println(objects, err3)
	objects, err3 = c.ListObjects("SquirrelSave", &swift.ListObjectsOpts{Delimiter: '/', Path: "Downloads/"})
	fmt.Println(objects, err3)
	fmt.Println(c.AccountInfo())
	fmt.Println(c.CreateContainer("sausage"))

	fmt.Println("Create", c.CreateObjectString("sausage", "test_object", "12345", ""))
	fmt.Println(c.GetObjectString("sausage", "test_object"))
	fmt.Println(c.GetObjectString("sausage", "test_object"))
	fmt.Println("delete 1", c.DeleteObject("sausage", "test_object"))
	fmt.Println("delete 2", c.DeleteObject("sausage", "test_object"))

	fmt.Println("delete container 1", c.DeleteContainer("sausage"))
	fmt.Println("delete container again", c.DeleteContainer("sausage"))
	fmt.Println(c.ContainerInfo("SquirrelSave"))

}
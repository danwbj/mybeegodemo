package blog

import (
	"fmt"
	_ "mybeegodemo/app/modules/blog/routers"  
	// _ "mybeegodemo/app/modules/blog/models"  
)

func init() {
	fmt.Println("blog init")
}

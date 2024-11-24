package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(ComparePassword("$2a$04$9xia989r7FDgyM3XKbtvuuVjAkILpN/TAF0TyGvbuJFkCUHbulIz6", "1234"))
}

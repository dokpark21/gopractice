package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Decoder에는 Reader를 제공해서 새로운 디코더에 대한 포인터를 반환받는다.
	// 반대로 Encoder에는 Writer를 제공해서 새로운 인코더에 대한 포인터를 반환 받는다.
	// 그 후 얻은 디코더나 인코더로 디코딩 및 인코딩를 실행한다.

	// Decode는 Unmarshal과 비슷하다. JSON으로 인코딩된 값을 value에 저장한다.

	// Writer , string
	// os.Stdout은 표준콘솔출력으로 io.Writer를 지원하는 타입이다.
	fmt.Fprintln(os.Stdout, "Hello golang!")
	// Writer , string
	io.WriteString(os.Stdout, "Hello golang!")
}

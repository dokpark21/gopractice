package main

import "fmt"

func main() {
	// 슬라이스는 배열을 기반으로 만들어 졌다. 슬라이스는 동적으로 크기가 바뀔 수 있다.
	// 대체로 슬라이스를 만들 때에는 합성 리터럴을 사용한다
	// make를 사용해서도 슬라이스를 만들 수 있다
	// 원소가  얼마나 들어갈지 알고 있으면 make를 사용해 하부의 배열을 만들어 줄 수 있다.
	// 슬라이스는 배열의 위에 생성되기 때문에 하부의 배열이 충분히 많은 값들을 원하는 만큼 저장할 수 있도록 만들어주면 컴파일러와 런타임이 슬라이스가 바뀔 때마다 배열을 복사해서 새로운 배열을 만들고
	// 이전의 슬라이스는 삭제하는 형식이다
	// make(slice, length, 하부 배열의 크기(새로운 슬라이스 생성 시 사용))
	x := make([]int, 10, 11)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) //cap은 용령

	x[0] = 10
	x[1] = 11
	x[2] = 21
	fmt.Println(x)

	//append하면 길이는 늘어나지만 용량은 그대로 이다.
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//만약 설정한 하부 배열 크기를 능가하게 된다면 알아서 증가한다.
	x = append(x, 123)
	fmt.Println(cap(x))
}

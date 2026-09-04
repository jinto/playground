# playground

Taskyard 실험용 작은 Go 저장소. 이슈를 만들고 자동 실행이 코드를 고치는 흐름을 시험하는 데 쓴다.

## 요구 사항

- Go 1.26 이상 (`go.mod` 참고)

## 구성

| 파일 | 역할 |
| --- | --- |
| `greet.go` | `Greet(name string) string` — `Hello, <name>!` 형식의 인사말을 돌려준다. |
| `greet_test.go` | `Greet`의 기본 동작을 확인하는 테스트 |

## 사용 예

모듈 경로가 `playground`이므로 아래 코드는 이 모듈 안(예: `cmd/` 하위)에서만 빌드된다.

```go
package main

import (
	"fmt"

	"playground"
)

func main() {
	fmt.Println(playground.Greet("Jay")) // Hello, Jay!
}
```

## 테스트

```sh
go test ./...
```

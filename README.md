# pangyo-lunch
## HUP PoC 프로젝트 저장소

### 로컬 환경에 설치
go 설치
git clone repository
go env -w GO111MODULE=auto (gin-gonic 패키지 설치 오류 날 경우)
go get -u github.com/gin-gonic/gin
go build .
go run .

### 테스트 URL
localhost:8080/restaurants
localhost:8080/restaurnats/정식류, 일식류, 국물 및 국밥류, 면류, 양식류, 회식류 택 1
localhost:8080/restaurnats/정식류, 일식류, 국물 및 국밥류, 면류, 양식류, 회식류 택 1/random

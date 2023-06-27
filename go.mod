module awesomeProject

//하나의 프로젝트 단위. 모듈내에 메인패키지가 여러개있는건 상관없지만 메인패키지안에 메인함수는 반드시 하나
// https://dev-yakuza.posstree.com/ko/golang/module/
// go에서는 'go get' 명령어로 외부 package를 가져올 수 있다.
// Go는 Node.js의 NPM과 같은 중앙 package 저장소가 없다
//  문제는... 'go get'으로는 특정 branch나 tag, version 를 지정하지 못해서  가져올 수 없다는 것이다.
//모든 packages 는 $GOPATH/src 아래에 source file들이 저장되고, Go는 바로 compile 해서 $GOPATH/pkg 아래에 .a 파일을 생성한다.
// 그리고 'go build' 로 compile 하면 나의 main packages 는 $GOPATH/bin 에 실행파일로 생성되고 내부 참조 packages는 $GOPATH/pkg 에 .a 파일이 생성된다.

go 1.20

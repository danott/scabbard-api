default: build

clean:
	rm scabbard-osx scabbard-linux

build:
	go build -o scabbard-osx
	GOOS=linux GOARCH=386 go build -o scabbard-linux

deploy:
	-ssh danott.us "rm /home/danott/go.danott.us/scabbard.fcgi"
	scp scabbard-linux danott.us:go.danott.us/scabbard.fcgi

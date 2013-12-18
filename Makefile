all:
	go build iksoj.go

install:
	install -Dm755 iksoj /usr/bin/iksoj

prepare:
	install -Dm755 iksoj usr/bin/iksoj

uninstall:
	rm /usr/bin/iksoj

clean:
	rm iksoj

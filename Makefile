installto?=/usr/bin

all:
	go build iksoj.go

install:
	install -Dm755 iksoj $(installto)/iksoj

uninstall:
	rm /usr/bin/iksoj

clean:
	rm iksoj

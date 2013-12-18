
# Maintainter: Joop Kiefte <joop@kiefte.eu>
pkgname=iksoj
pkgver=0.1
pkgrel=1
pkgdesc="iksoj is a command to convert the Esperanto x-system to unicode"
arch=('i686' 'x86_64')
url=http://github.com/LaPingvino/iksoj/
license=('CC0')
makedepends=('go')
source=("master.zip::https://github.com/LaPingvino/iksoj/archive/master.zip")
md5sums=('266b3c00ec43683feb5e18fbaeed6b44')

build() {
	cd "${srcdir}/${pkgname}-${pkgver}"
	make
}

package() {
	cd "${srcdir}/${pkgname}-${pkgver}"
	make prepare
}


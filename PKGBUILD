
# Maintainter: Joop Kiefte <joop@kiefte.eu>
pkgname=iksoj
pkgver=0.3
pkgrel=1
pkgdesc="iksoj is a command to convert the Esperanto x-system to unicode"
arch=('i686' 'x86_64')
url=http://github.com/LaPingvino/iksoj/
license=('CC0')
makedepends=('go')
source=("${pkgname}-${pkgver}.tar.gz::http://github.com/LaPingvino/iksoj/archive/${pkgver}.tar.gz")
md5sums=('6d5ae73f3760d7982023ec69218dcdef')

build() {
	cd "${srcdir}/${pkgname}-${pkgver}"
	make
}

package() {
	cd "${srcdir}/${pkgname}-${pkgver}"
	sed -i "s@/usr/bin@${pkgdir}/usr/bin@" Makefile
	make install
}


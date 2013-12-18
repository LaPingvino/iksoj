
# Maintainter: Joop Kiefte <joop@kiefte.eu>
pkgname=iksoj
pkgver=0.2
pkgrel=1
pkgdesc="iksoj is a command to convert the Esperanto x-system to unicode"
arch=('i686' 'x86_64')
url=http://github.com/LaPingvino/iksoj/
license=('CC0')
makedepends=('go')
source=("${pkgname}-${pkgver}.tar.gz::http://github.com/LaPingvino/iksoj/archive/${pkgver}.tar.gz")
md5sums=('1c8f7dc7d781b4615dd584ec694e2c7f')

build() {
	cd "${srcdir}/${pkgname}-${pkgver}"
	make
}

package() {
	cd "${srcdir}/${pkgname}-${pkgver}"
	sed -i "s@/usr/bin@${pkgdir}/usr/bin@" Makefile
	make install
}


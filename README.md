# lang

Aquí vamos a registrar nuestro proyecto de crear un lenguaje de programación. Esto incluye el diseño de lenguaje y la implementación de un compilador.

## Uso

Tenemos un `Makefile` con varios comandos:

- `make build-docker` construye la imagen Docker para el sistema,
- `make run-docker` ejecuta la imagen,
- `make run-tests` ejecuta las pruebas unitarias.

Los siguientes comandos se deberían ejecutar en la terminal del contenedor, tras haber hecho `make run-docker`:

- `make install-dependencies` instala las dependencias de Go,
- `make run-tests-ci` ejecuta las pruebas que se van a hacer en la CI de GitHub,
- `make generate-parsing-files` genera el código de ANTLR para el sistema. Se debería ejecutar si la gramática del lenguaje fue modificada o los archivos no existen. El código generado por ANTLR no es rastreado en este repositorio.

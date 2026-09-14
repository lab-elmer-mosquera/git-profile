Crear los comandos que faciliten crear una identidad en git y la creacion de un ssh.

Crear la funcionalidad git-shh include, para crear el include en el git config de el ssh a utilizar.

---

Necesito una forma rapida de configurar git ahora mismo, aunque posteriorme optare por crear la herrmienta descrita en la vision esta es una version simple de optener las funcionalidades que necesito ahora.

Crear multiples ssh y poder asociarlas a un directorio especifico, aunque exite una llamada **`git-context`** que se evaluo en la vision, prefiero crear una de forma con IA, para optener justo lo que me puede sacar del paso sin gestionar una herramienta externa con la cual tendria que familiarizarme porteriormente.

## Modo de uso

**gprofile-includef** Es el comando base cuando se ejecuta

Ejecutar

```bash
gprofile-includef
```

Tiene el siguiente flujo

```
=== Configuración de Nuevo Perfil Git/SSH ===
Ingresa tu nombre para Git: Juan Perez
Ingresa tu correo para Git: juan.perez@udea.edu.co
Ingresa el nombre del perfil (ej: udea): udea

[1/3] Generando llave SSH Ed25519 en: /home/usuario/.ssh/id_ed25519_udea...
[2/3] Creando archivo de configuración Git: /home/usuario/.gitconfig-udea...
[3/3] ¡Perfil creado con éxito!

Para activar este perfil en una carpeta, corre:
  gprofile-includef --include . --profile udea


```


Para hacer incluidef a una carpeta

```bash

./gprofile-includef --include . --profile udea

```
---
tipo: andamio
id: AND-001
requisito: REQ-001
decisiones:
  - ADR-001
---

# [AND-001] Modo de uso interactivo, CLI y estructura del derivado MVP `gprofile-includef`

Dado que existe [[REQ-001]], con [[ADR-001]] relacionada, se deduce la siguiente interfaz de usuario y estructura de implementación modular simple sin dependencias externas para la versión beta en Windows.

---

## 💻 1. Flujo de Interacción y Modo de Uso Esperado

### A. Creación Interactiva de Perfil Git / SSH
Al ejecutar `./gprofile-includef` sin argumentos (o con el subcomando de inicialización), se despliega el siguiente asistente interactivo:

```text
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

### B. Vinculación de Perfil a Carpeta (`includeIf`)
Para asociar el perfil recién creado a un directorio específico de trabajo/estudio, el usuario ejecuta desde la CLI:

```bash
./gprofile-includef --include . --profile udea
```

---

## 📂 2. Estructura de Implementación del Derivado MVP

Para mantener el código aislado, modular y sin dependencias externas pesadas, la solución residirá dentro de la sección `software/` del repositorio:

* **Ubicación en el repositorio:**  
  `github.com/owner/git-profile/software/beta/win/gprofile-includef`

* **Árbol de archivos del subproducto:**
  ```text
  software/
  └── beta/
      └── win/
          └── gprofile-includef/
              ├── gprofile-includef.sh    # Script principal ejecutable (Git Bash / MSYS)
              ├── lib/
              │   ├── ssh_gen.sh          # Módulo de generación de llaves ed25519
              │   └── git_config.sh       # Módulo de inyección de includeIf
              └── README.md               # Guía rápida de instalación y uso local
  ```

---

## ⚙️ 3. Reglas Internas de Deducción Técnica
1. **Rutas en Windows / Git Bash:** Al ejecutar `--include .`, el script resuelve la ruta absoluta mediante `pwd -P` o `cygpath -m` y asegura la barra diagonal de cierre `/` para que Git la reconozca en `includeIf.gitdir:`.
2. **Generación de SSH:** Utiliza de forma nativa `ssh-keygen -t ed25519 -C "<email>" -f "$HOME/.ssh/id_ed25519_<perfil>"` sin solicitar passphrase para agilizar el flujo de aprendizaje.
3. **Comando SSH en Git:** El archivo `~/.gitconfig-<perfil>` incluye automáticamente la directiva:
   ```ini
   [core]
       sshCommand = ssh -i "~/.ssh/id_ed25519_<perfil>" -F /dev/null
   ```

---
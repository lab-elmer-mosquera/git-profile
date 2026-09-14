---
tipo: inbox
id: inbox-s37-perfil-como-entidad
semana: s37
destino: seccionar
  - VA (actualización, si aplica)
  - Dominio "Perfil"
  - ADR (modelo, SSH, sintaxis, persistencia)
  - REQ (derivados)
---

# Inbox s37 — Perfil como entidad


## A. Lectura operativa de la visión

Git Profile busca que el usuario pueda expresar:

    Perfil + operación Git + argumentos

sin tener que reconstruir manualmente el comando final ni
modificar la configuración global de Git.

El Perfil se almacena de forma independiente y reúne la información
necesaria para representar una identidad de trabajo.

---

## B. Perfil

Hipótesis de atributos que podrían formar parte de un Perfil:

- name
- email
- identidad SSH (referencia, no gestión)
- configuración relacionada con repositorios y remotos (abierto)

Representación conceptual:

    personal
    ├── name
    ├── email
    └── SSH identity (referencia)

SSH puede seguir gestionándose con mecanismos tradicionales
(ssh-keygen y afines). Git Profile no es un gestor de claves.

---

## C. Operaciones

Interfaz tentativa:

    git <perfil> <operación> <args>

Ejemplo:

    git personal clone owner/repo.git

resolución posible:

    git clone git@github-personal:owner/repo.git

Git Profile actúa como capa de resolución y composición
sobre la operación nativa de Git.

---

## D. Clone

Dado:

    git personal clone owner/repo.git

pasos hipotéticos:

1. resolver el Perfil `personal`
2. determinar la identidad SSH correspondiente
3. construir el remoto apropiado
4. ejecutar el `git clone` tradicional
5. configurar el repositorio resultante con la identidad del Perfil
   cuando corresponda

La operación real la ejecuta Git.

---

## E. Configuración del repositorio

Tras clone o creación, Git Profile podría fijar a nivel local:

- user.name
- user.email
- remote.origin.url (abierto)

Objetivo: no tocar la configuración global del usuario.

---

## F. Uso puntual de un Perfil

    git personal commit -m "message"

Permitiría usar la identidad `personal` en una operación concreta
sin modificar la configuración persistente del repositorio.

Distinción:

- Perfil asociado al repositorio
- Perfil usado explícitamente para una operación

---

## G. Principio fundamental

    Perfil
       ↓
    resolver información
       ↓
    construir contexto / comando
       ↓
    ejecutar Git

Git Profile une las piezas. Git sigue haciendo el trabajo.

---

## H. Fuera de alcance

- reemplazar Git
- reemplazar SSH
- administrar claves SSH
- modificar globalmente la configuración del usuario

Operaciones independientes de identidad siguen usándose con Git directo:

    git status
    git diff
    git log

---

## I. Diferenciación

Las herramientas existentes resuelven el problema mediante
cambio / activación / inyección de configuración de Git.

Git Profile explora:

    Perfil
       ↓
    operación
       ↓
    resolución de identidad
       ↓
    Git / SSH

El Perfil no es solo un conjunto de valores activables:
es una entidad del dominio desde la cual se construyen operaciones.

---

## J. Cuestiones abiertas

- ¿Qué atributos forman exactamente parte de un Perfil?
- ¿Cómo se representa una identidad SSH?
- ¿Un Perfil puede tener múltiples identidades SSH según el host?
- ¿Cómo se resuelve owner/repo contra un host concreto?
- ¿Qué operaciones admiten un Perfil?
- ¿Cómo se implementa un override puntual sin configuración persistente?
- ¿Qué configuración queda local tras clone?
- ¿Cómo se comporta un Perfil cuando el repo ya tiene identidad?
- ¿`remote.origin.url` forma parte del Perfil o es resultado de resolución?
- ¿Cómo se diferencian credenciales, autoría e identidad SSH?
- ¿Qué parte vive en SQLite y qué parte se delega a Git/SSH?
- ¿La sintaxis `git <profile> <operation>` será definitiva?

---

## K. Principios de diseño (tentativos)

- Git sigue siendo Git.
- SSH sigue siendo SSH.
- El Perfil reúne y relaciona información existente.
- Git Profile resuelve y conecta esas piezas.
- La config global de Git no es el mecanismo para cambiar de Perfil.
- Un Perfil puede aplicarse persistente o puntualmente.
- Mínima abstracción necesaria.
- Validar contra casos de uso reales antes de fijar API.

---

## Seccionamiento pendiente

Al cerrar s37 (o cuando el material madure), este inbox debe resolverse en:

- [ ] ¿Actualiza la VA vigente? → sí / no / qué sección
- [ ] ¿Nace `01_Product/Dominios/perfil.md`? → sí / no
- [ ] ADR candidatos identificados → listar
- [ ] REQ candidatos derivados → listar (tras ADR)
- [ ] Cuestiones abiertas que siguen abiertas → mover a discusión de ADR
- [ ] Este archivo → archivar cuando todo lo anterior esté resuelto

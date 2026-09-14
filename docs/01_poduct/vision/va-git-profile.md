---
tipo: vision
id: VA-1
status: Activa
naturaleza: Utilidad
alternativas_evaluadas: true
---
## Qué es y cómo se usa
 
 Git ofrece entidades fundamentales como repositorios, ramas, remotos, etiquetas y commits, pero carece de una representación explícita de la identidad desde la cual se realizan determinadas operaciones.
 
 Git Profile introduce la entidad Perfil como un concepto de primer nivel dentro del flujo de trabajo de Git. Un Perfil representa una identidad reutilizable sobre la cual pueden ejecutarse aquellas operaciones cuya semántica depende de quién realiza la acción. Dicha identidad puede comprender todos aquellos atributos que participan en esas operaciones, incluyendo, cuando corresponda, la autoría y los mecanismos de firma utilizados para garantizar su autenticidad.
 
 La visión del proyecto es incorporar esta entidad al modelo de trabajo de Git para que las operaciones que requieren una identidad puedan expresarse y ejecutarse de forma explícita sobre un Perfil, manteniendo la identidad como parte del propio flujo de trabajo y no como una configuración implícita.
## Lo que no es
Un Perfil no reemplaza las entidades ni las operaciones propias de Git. Su alcance se limita exclusivamente a aquellas acciones cuya ejecución depende de una identidad de trabajo. En consecuencia, Git Profile no redefine ni envuelve operaciones cuyo resultado es independiente del Perfil, como `git status`, `git diff`, `git log` o cualquier otra operación que no requiera conocer quién realiza la acción.
Git Profile tampoco pretende sustituir a Git ni replicar su interfaz; su propósito es incorporar la entidad Perfil al dominio de Git y proporcionar operaciones coherentes con dicha entidad.
## Alternativas cercanas (opcional)
- **`git-profile`** (dmfutcher, Rust) — gestión de perfiles como CRUD de configuración (nombre/email/URL de remoto), sin noción de operación dependiente de identidad como concepto de dominio.
- **`git-context`** (Go, YAML) — la más cercana técnicamente: asigna perfiles a directorios para aplicación automática vía `includeIf`. Sigue modelando el perfil como configuración a volcar en `~/.gitconfig`, no como entidad con operaciones propias.
- **`gid`** — inyecta la configuración del perfil por comando sin persistir nada en el repositorio ni en la config global. 
- **`gitup`** (Rust) — asocia llaves de firma GPG/SSH a cada perfil. Con la firma ya confirmada dentro del alcance de Perfil, esta es la alternativa con mayor solapamiento funcional de la lista; la diferenciación frente a ella depende por completo del modelo de entidad-con-operaciones, no de la cobertura de firma.
- **`bgit`** (Go) — switch único que gestiona config de Git y de SSH juntos.
- Mecanismo nativo: `git config` con `includeIf`, sobre el cual la mayoría de estas herramientas están construidas.

Todas resuelven el problema mediante configuración a activar/cambiar (`use`, `switch`) que se vuelca al `.gitconfig`, no como una entidad de dominio con operaciones propias. Git Profile propone tratar el Perfil como entidad explícita del dominio sobre la cual se construyen operaciones específicas (como la resolución automática en el momento del `clone`), sin extender el alcance de Git hacia operaciones independientes de identidad.

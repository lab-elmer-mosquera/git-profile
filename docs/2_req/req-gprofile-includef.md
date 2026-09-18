---
tipo: requisito
id: REQ-001
status: Pendiente
prioridad: Alta
modulo: "Subproductos"
origen: ADR-001
subproducto: "gprofile-includeif"
responsable: "@lab-elmer-mosquera"
sprint: "Sprint-01"
rama: "feature/REQ-001-gprofile-includeif-v0.1-windows"
---

# [REQ-001] Herramienta Beta `gprofile-includeif` para creación de identidades, llaves SSH y vinculación por directorio

## 🎯 1. El "Por Qué"
Durante etapas de exploración técnica y aprendizaje de redes, el desarrollo de la arquitectura final consolidada en Go resulta demasiado lento para las necesidades operativas inmediatas. Se requiere un mecanismo simple que cree llaves SSH, defina identidades de Git y las vincule automáticamente a directorios específicos mediante la directiva `includeIf` de Git, resolviendo el problema de múltiples perfiles sin depender de herramientas externas complejas.

## 👥 2. Actores y Alcance
* **Actor:** Desarrollador / Estudiante en entornos Windows (Git Bash) y Linux.
* **Dentro del alcance:**
  * Creación asistida de llaves SSH (`ed25519`) por perfil.
  * Generación de archivos de configuración de perfil independientes (`~/.gitconfig-<perfil>`).
  * Inyección automática del bloque `[includeIf "gitdir:<directorio>/"]` en el `~/.gitconfig` global.
* **Fuera del alcance:**
  * Reemplazar comandos nativos de Git o gestionar almacenamiento cifrado avanzado.

## 📋 3. Criterios de Aceptación
- [ ] **Creación de Identidad y Llave SSH:** Dado un nombre de perfil, un correo y un nombre de usuario, la herramienta genera la clave SSH `~/.ssh/id_ed25519_<perfil>` y escribe la identidad en `~/.gitconfig-<perfil>`.
- [ ] **Vinculación por Directorio (`includeIf`):** Dado un perfil existente y una ruta de directorio, la herramienta inyecta la directiva `includeIf "gitdir:<directorio>/"` en el `~/.gitconfig` global apuntando al perfil correspondiente.
- [ ] **Multiplataforma (Windows/Linux):** La solución opera mediante Bash/Shell compatible tanto con Linux como con Git Bash en Windows.

## 🔗 4. Trazabilidad
* **Origen:** [[ADR-001]]
* **Decisiones relacionadas:**
* **Casos de Prueba:**
* **Implementación:**